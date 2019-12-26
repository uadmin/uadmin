package uadmin

import (
	"math/big"

	"crypto/rand"
	"math"
	network "net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// CookieTimeout is the timeout of a login cookie in seconds
var CookieTimeout = -1

// Salt is extra salt added to password hashing
var Salt = ""

// bcryptDiff
var bcryptDiff = 12

var cachedSessions map[string]Session

// GenerateBase64 generates a base64 string of length length
func GenerateBase64(length int) string {
	base := new(big.Int)
	base.SetString("64", 10)

	base64 := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"
	tempKey := ""
	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, base)
		tempKey += string(base64[int(index.Int64())])
	}
	return tempKey
}

// GenerateBase32 generates a base64 string of length length
func GenerateBase32(length int) string {
	base := new(big.Int)
	base.SetString("32", 10)

	base32 := "234567abcdefghijklmnopqrstuvwxyz"
	tempKey := ""
	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, base)
		tempKey += string(base32[int(index.Int64())])
	}
	return tempKey
}

// hashPass Generates a hash from a password and salt
func hashPass(pass string) string {
	password := []byte(pass + Salt)
	hash, err := bcrypt.GenerateFromPassword(password, bcryptDiff)
	if err != nil {
		Trail(ERROR, "uadmin.auth.hashPass.GenerateFromPassword: %s", err)
		return ""
	}
	return string(hash)
}

// IsAuthenticated returns if the http.Request is authenticated or not
func IsAuthenticated(r *http.Request) *Session {
	key := getSession(r)
	s := Session{}
	if CacheSessions {
		s = cachedSessions[key]
	} else {
		Get(&s, "`key` = ?", key)
	}
	if isValidSession(r, &s) {
		return &s
	}
	return nil
}

func isValidSession(r *http.Request, s *Session) bool {
	if s != nil && s.ID != 0 {
		if s.Active && !s.PendingOTP && (s.ExpiresOn == nil || s.ExpiresOn.After(time.Now())) {
			if s.User.ID != s.UserID {
				Get(&s.User, "id = ?", s.UserID)
			}
			if s.User.Active && (s.User.ExpiresOn == nil || s.User.ExpiresOn.After(time.Now())) {
				// Check for IP restricted session
				if RestrictSessionIP {
					ip, _, _ := network.SplitHostPort(r.RemoteAddr)
					return ip == s.IP
				}
				return true
			}
		}
	}
	return false
}

// GetUserFromRequest returns a user from a request
func GetUserFromRequest(r *http.Request) *User {
	s := getSessionFromRequest(r)
	if s != nil {
		u := User{}
		Get(&u, "id = ?", s.UserID)
		if u.ID != 0 {
			return &u
		}
	}
	return nil
}

// getUserFromRequest returns a session from a request
func getSessionFromRequest(r *http.Request) *Session {
	key := getSession(r)
	s := Session{}
	Get(&s, "`key` = ?", key)
	if s.ID != 0 {
		return &s
	}
	return nil
}

// Login return *User and a bool for Is OTP Required
func Login(r *http.Request, username string, password string) (*Session, bool) {
	// Get the user from DB
	user := User{}
	Get(&user, "username = ?", username)
	if user.ID == 0 {
		IncrementMetric("uadmin/security/invalidlogin")
		go func() {
			log := &Log{}
			if r.Form == nil {
				r.ParseForm()
			}
			r.Form.Set("login-status", "invalid username")
			log.SignIn(username, log.Action.LoginDenied(), r)
			log.Save()
		}()
		return nil, false
	}
	s := user.Login(password, "")
	if s != nil && s.ID != 0 {
		s.IP, _, _ = network.SplitHostPort(r.RemoteAddr)
		s.Save()
		if s.Active && (s.ExpiresOn == nil || s.ExpiresOn.After(time.Now())) {
			s.User = user
			if s.User.Active && (s.User.ExpiresOn == nil || s.User.ExpiresOn.After(time.Now())) {
				IncrementMetric("uadmin/security/validlogin")
				// Store login successful to the user log
				go func() {
					log := &Log{}
					if r.Form == nil {
						r.ParseForm()
					}
					log.SignIn(user.Username, log.Action.LoginSuccessful(), r)
					log.Save()
				}()
				return s, s.User.OTPRequired
			}
		}
	} else {
		go func() {
			log := &Log{}
			if r.Form == nil {
				r.ParseForm()
			}
			r.Form.Set("login-status", "invalid password or inactive user")
			log.SignIn(username, log.Action.LoginDenied(), r)
			log.Save()
		}()
	}

	IncrementMetric("uadmin/security/invalidlogin")
	return nil, false
}

// Login2FA login using username, password and otp for users with OTPRequired = true
func Login2FA(r *http.Request, username string, password string, otpPass string) *Session {
	s, otpRequired := Login(r, username, password)
	if s != nil {
		if otpRequired && s.User.VerifyOTP(otpPass) {
			s.PendingOTP = false
			s.Save()
		}
		return s
	}
	return nil
}

// Logout logs out a user
func Logout(r *http.Request) {
	s := getSessionFromRequest(r)
	if s.ID == 0 {
		return
	}

	// Store Logout to the user log
	func() {
		log := &Log{}
		log.SignIn(s.User.Username, log.Action.Logout(), r)
		log.Save()
	}()

	s.Logout()
	IncrementMetric("uadmin/security/logout")
}

// ValidateIP is a function to check if the IP in the request is allowed in the allowed based on allowed
// and block strings
func ValidateIP(r *http.Request, allow string, block string) bool {
	allowed := false
	allowSize := uint32(0)

	allowList := strings.Split(allow, ",")
	for _, net := range allowList {
		if v, size := requestInNet(r, net); v {
			allowed = true
			if size > allowSize {
				allowSize = size
			}
		}
	}

	blockList := strings.Split(block, ",")
	for _, net := range blockList {
		if v, size := requestInNet(r, net); v {
			if size > allowSize {
				allowed = false
				break
			}
		}
	}
	if !allowed {
		IncrementMetric("uadmin/security/blockedip")
	}
	return allowed
}

func requestInNet(r *http.Request, net string) (bool, uint32) {
	// Check if the IP is V4
	if strings.Contains(r.RemoteAddr, ".") {
		var ip uint32
		var subnet uint32
		var oct uint64
		var mask uint32

		// check if the net is IPv4
		if !strings.Contains(net, ".") && net != "*" && net != "" {
			return false, 0
		}

		// Convert the IP to uint32
		ipParts := strings.Split(strings.Split(r.RemoteAddr, ":")[0], ".")
		for i, o := range ipParts {
			oct, _ = strconv.ParseUint(o, 10, 8)
			ip += uint32(oct << ((3 - uint(i)) * 8))
		}

		// convert the net to uint32
		// but first convert standard nets to IPv4 format
		if net == "*" {
			net = "0.0.0.0/0"
		} else if net == "" {
			net = "255.255.255.255/32"
		} else if !strings.Contains(net, "/") {
			net += "/32"
		}
		ipParts = strings.Split(strings.Split(net, "/")[0], ".")
		for i, o := range ipParts {
			oct, _ = strconv.ParseUint(o, 10, 8)
			subnet += uint32(oct << ((3 - uint(i)) * 8))
		}

		maskLength := getNetSize(r, net)
		mask -= uint32(math.Pow(2, float64(32-maskLength)))
		return ((ip & mask) ^ subnet) == 0, uint32(maskLength)
	}
	// Process IPV6
	var ip1 uint64
	var ip2 uint64
	var subnet1 uint64
	var subnet2 uint64
	var oct uint64
	var mask1 uint64
	var mask2 uint64

	// check if the net is IPv6
	if strings.Contains(net, ".") && net != "*" && net != "" {
		return false, 0
	}

	// Normalize IP
	ipS := r.RemoteAddr              // [::1]:10000
	ipS = strings.Trim(ipS, "[")     // ::1]:10000
	ipS = strings.Split(ipS, "]")[0] // ::1
	if strings.HasPrefix(ipS, "::") {
		ipS = "0" + ipS
	} else if strings.HasSuffix(ipS, "::") {
		ipS = ipS + "0"
	}
	// find and replace ::
	ipParts := strings.Split(ipS, ":")
	ipFinalParts := []uint16{}
	processedDC := false
	for i := range ipParts {
		if ipParts[i] == "" && !processedDC {
			processedDC = true
			for counter := 0; counter < 8-i-(len(ipParts)-(i+1)); counter++ {
				//oct, _ = strconv.ParseUint(ipParts[i], 16, 16)
				ipFinalParts = append(ipFinalParts, uint16(0))
			}
		} else {
			oct, _ = strconv.ParseUint(ipParts[i], 16, 16)
			ipFinalParts = append(ipFinalParts, uint16(oct))
		}
	}

	// Parse the IP into two uint64 variables
	for i := 0; i < 4; i++ {
		oct = uint64(ipFinalParts[i])
		ip1 += uint64((oct << ((3 - uint(i)) * 16)))
	}
	for i := 0; i < 4; i++ {
		oct = uint64(ipFinalParts[i+4])
		ip2 += uint64((oct << ((3 - uint(i)) * 16)))
	}

	subnetv6 := net
	if subnetv6 == "*" {
		subnetv6 = "0::0/0"
	} else if subnetv6 == "" {
		subnetv6 = "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128"
	} else if !strings.Contains(subnetv6, "/") {
		subnetv6 = subnetv6 + "/128"
	}
	maskS := strings.Split(subnetv6, "/")[1]
	subnetv6 = strings.Split(subnetv6, "/")[0]
	if strings.HasPrefix(subnetv6, "::") {
		subnetv6 = "0" + subnetv6
	} else if strings.HasSuffix(subnetv6, "::") {
		subnetv6 = subnetv6 + "0"
	}
	// find and replace ::
	ipParts = strings.Split(subnetv6, ":")
	ipFinalParts = []uint16{}
	processedDC = false
	for i := range ipParts {
		if ipParts[i] == "" && !processedDC {
			processedDC = true
			for counter := 0; counter < 8-i-(len(ipParts)-(i+1)); counter++ {
				//oct, _ = strconv.ParseUint(ipParts[i], 16, 16)
				ipFinalParts = append(ipFinalParts, uint16(0))
			}
		} else {
			oct, _ = strconv.ParseUint(ipParts[i], 16, 16)
			ipFinalParts = append(ipFinalParts, uint16(oct))
		}
	}

	for i := 0; i < 4; i++ {
		oct = uint64(ipFinalParts[i])
		subnet1 += uint64((oct << ((3 - uint(i)) * 16)))
	}
	for i := 0; i < 4; i++ {
		oct = uint64(ipFinalParts[i+4])
		subnet2 += uint64((oct << ((3 - uint(i)) * 16)))
	}

	oct, _ = strconv.ParseUint(maskS, 10, 8)
	maskLength := int(oct)

	maskLength2 := math.Max(float64(maskLength-64), 0)
	maskLength1 := float64(maskLength) - maskLength2

	mask1 -= uint64(math.Pow(2, 64-maskLength1))
	mask2 -= uint64(math.Pow(2, 64-maskLength2))
	if maskLength1 == 0 {
		mask1 = 0
	}
	if maskLength2 == 0 {
		mask2 = 0
	}

	xored1 := (ip1 & mask1) ^ subnet1
	xored2 := (ip2 & mask2) ^ subnet2

	return xored1 == 0 && xored2 == 0, uint32(maskLength)
}

func getNetSize(r *http.Request, net string) int {
	var maskLength int
	var oct uint64

	// Check if the IP is V4
	if strings.Contains(r.RemoteAddr, ".") {
		// Get the Netmask
		oct, _ = strconv.ParseUint(strings.Split(net, "/")[1], 10, 8)
		maskLength = int(oct)
	}
	return maskLength
}

func getSessionByKey(key string) *Session {
	s := Session{}
	Get(&s, "`key` = ?", key)
	if s.ID == 0 {
		return nil
	}
	return &s
}

func getSession(r *http.Request) string {
	key, err := r.Cookie("session")
	if err == nil && key != nil {
		return key.Value
	}
	if r.Method == "GET" && r.FormValue("session") != "" {
		return r.FormValue("session")
	}
	if r.Method == "POST" {
		r.ParseForm()
		if r.FormValue("session") != "" {
			return r.FormValue("session")
		}
	}
	return ""
}
