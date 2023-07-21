package uadmin

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net"
	"os"
	"path"

	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// CookieTimeout is the timeout of a login cookie in seconds.
// If the value is -1, then the session cookie will not have
// an expiry date.
var CookieTimeout = -1

// Salt is added to password hashing
var Salt = ""

// JWT secret for signing tokens
var JWT = ""

// jwtIssuer is a URL to identify the application issuing JWT tokens.
// If left empty, a partial hash of JWT will be assigned. This is also
// used to identify the as JWT audience.
var JWTIssuer = ""

var JWTAlgo = "HS256" //"RS256"

// AcceptedJWTIssuers is a list of accepted JWT issuers. By default the
// local JWTIssuer is accepted. To accept other issuers, add them to
// this list
var AcceptedJWTIssuers = []string{}

// bcryptDiff
var bcryptDiff = 12

// cachedSessions is variable for keeping active sessions
var cachedSessions map[string]Session

// invalidAttempts keeps track of invalid password attempts
// per IP address
var invalidAttempts = map[string]int{}

var CustomJWT func(r *http.Request, s *Session, payload map[string]interface{}) map[string]interface{}

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

// GenerateBase32 generates a base32 string of length length
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
	if len(password) > 72 {
		password = password[:72]
	}
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

	if strings.HasPrefix(key, "nouser:") {
		return nil
	}

	s := getSessionByKey(key)
	if isValidSession(r, s) {
		return s
	}
	return nil
}

// SetSessionCookie sets the session cookie value, The the value passed in
// session is nil, then the session assigned will be a no user session
func SetSessionCookie(w http.ResponseWriter, r *http.Request, s *Session) string {
	if s == nil {
		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    "nouser:" + GenerateBase64(124),
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
			Expires:  time.Now().AddDate(0, 0, 1),
		})
	} else {
		sessionCookie := &http.Cookie{
			Name:     "session",
			Value:    s.Key,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
		}
		if s.ExpiresOn != nil {
			sessionCookie.Expires = *s.ExpiresOn
		}
		http.SetCookie(w, sessionCookie)

		jwt := createJWT(r, s)
		jwtCookie := &http.Cookie{
			Name:     "access-jwt",
			Value:    jwt,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
		}
		if s.ExpiresOn != nil {
			jwtCookie.Expires = *s.ExpiresOn
		}
		http.SetCookie(w, jwtCookie)

		return jwt
	}
	return ""
}

func createJWT(r *http.Request, s *Session) string {
	if s == nil {
		return ""
	}
	if !isValidSession(r, s) {
		return ""
	}
	alg := JWTAlgo
	aud := JWTIssuer
	if r.Context().Value(CKey("aud")) != nil {
		aud = r.Context().Value(CKey("aud")).(string)
	}
	header := map[string]interface{}{
		"alg": alg,
		"typ": "JWT",
	}
	payload := map[string]interface{}{
		"sub": s.User.Username,
		"iat": s.LastLogin.Unix(),
		"iss": JWTIssuer,
		"aud": aud,
	}
	if s.ExpiresOn != nil {
		payload["exp"] = s.ExpiresOn.Unix()
	}

	// Check for custom JWT handler
	if CustomJWT != nil {
		payload = CustomJWT(r, s, payload)
	}

	if alg == "HS256" {
		jHeader, _ := json.Marshal(header)
		jPayload, _ := json.Marshal(payload)
		b64Header := base64.RawURLEncoding.EncodeToString(jHeader)
		b64Payload := base64.RawURLEncoding.EncodeToString(jPayload)

		hash := hmac.New(sha256.New, []byte(JWT+s.Key))
		hash.Write([]byte(b64Header + "." + b64Payload))
		signature := hash.Sum(nil)
		b64Signature := base64.RawURLEncoding.EncodeToString(signature)
		return b64Header + "." + b64Payload + "." + b64Signature
	} else if alg == "RS256" {
		buf, err := os.ReadFile(".jwt-rsa-private.pem")
		if err != nil {
			return ""
		}
		key, err := jwt.ParseRSAPrivateKeyFromPEM(buf)
		if err != nil {
			return ""
		}
		header["kid"] = "1"
		token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims(payload))

		for k, v := range header {
			token.Header[k] = v
		}

		tokenRaw, err := token.SignedString(key)

		if err != nil {
			return ""
		}
		return tokenRaw
	} else {
		Trail(ERROR, "Unknown algorithm for JWT (%s)", alg)
		return ""
	}

}

func isValidSession(r *http.Request, s *Session) bool {
	valid, otpPending := isValidSessionOTP(r, s)
	return valid && !otpPending
}

func isValidSessionOTP(r *http.Request, s *Session) (bool, bool) {
	if s != nil && s.ID != 0 {
		if s.Active && (s.ExpiresOn == nil || s.ExpiresOn.After(time.Now())) {
			if s.User.ID != s.UserID {
				Get(&s.User, "id = ?", s.UserID)
			}
			if s.User.Active && (s.User.ExpiresOn == nil || s.User.ExpiresOn.After(time.Now())) {
				// Check for IP restricted session
				if RestrictSessionIP {
					ip := GetRemoteIP(r)
					return ip == s.IP, s.PendingOTP
				}
				return true, s.PendingOTP
			}
		}
	}
	return false, false
}

// GetUserFromRequest returns a user from a request
func GetUserFromRequest(r *http.Request) *User {
	s := getSessionFromRequest(r)
	if s != nil {
		if s.User.ID == 0 {
			Get(&s.User, "id = ?", s.UserID)
		}
		if s.User.ID != 0 {
			return &s.User
		}
	}
	return nil
}

// getSessionFromRequest returns a session from a request
func getSessionFromRequest(r *http.Request) *Session {
	key := getSession(r)
	s := getSessionByKey(key)

	if s != nil && s.ID != 0 {
		return s
	}
	return nil
}

// Login return *User and a bool for Is OTP Required
func Login(r *http.Request, username string, password string) (*Session, bool) {
	if PreLoginHandler != nil {
		PreLoginHandler(r, username, password)
	}
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
			ctx := context.WithValue(r.Context(), CKey("login-status"), "invalid username")
			r = r.WithContext(ctx)
			log.SignIn(username, log.Action.LoginDenied(), r)
			log.Save()
		}()
		incrementInvalidLogins(r)
		return nil, false
	}
	s := user.Login(password, "")
	if s != nil && s.ID != 0 {
		s.IP = GetRemoteIP(r)
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
			ctx := context.WithValue(r.Context(), CKey("login-status"), "invalid password or inactive user")
			r = r.WithContext(ctx)
			log.SignIn(username, log.Action.LoginDenied(), r)
			log.Save()
		}()
	}

	incrementInvalidLogins(r)

	// Record metrics
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
		} else if otpRequired && !s.User.VerifyOTP(otpPass) && otpPass != "" {
			incrementInvalidLogins(r)
		}
		return s
	}
	return nil
}

func incrementInvalidLogins(r *http.Request) {
	// Increment password attempts and check if it reached
	// the maximum invalid password attempts
	ip := GetRemoteIP(r)
	invalidAttempts[ip]++

	if invalidAttempts[ip] >= PasswordAttempts {
		rateLimitLock.Lock()
		rateLimitMap[ip] = time.Now().Add(time.Duration(PasswordTimeout)*time.Minute).Unix() * RateLimit
		rateLimitLock.Unlock()
	}
}

// Login2FA login using username, password and otp for users with OTPRequired = true
func Login2FAKey(r *http.Request, key string, otpPass string) *Session {
	s := getSessionByKey(key)
	valid, otpPending := isValidSessionOTP(r, s)
	if valid {
		if otpPending && s.User.VerifyOTP(otpPass) {
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

	// Delete the cookie from memory if we sessions are cached
	if CacheSessions {
		delete(cachedSessions, s.Key)
	}

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
	ipStr := GetRemoteIP(r)
	// Check if the IP is V4
	if strings.Contains(ipStr, ".") {
		var ip uint32
		var subnet uint32
		var oct uint64
		var mask uint32

		// check if the net is IPv4
		if !strings.Contains(net, ".") && net != "*" && net != "" {
			return false, 0
		}

		// Convert the IP to uint32
		ipParts := strings.Split(strings.Split(ipStr, ":")[0], ".")
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
	ipS := GetRemoteIP(r)            // [::1]:10000
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
	if strings.Contains(GetRemoteIP(r), ".") {
		// Get the Netmask
		oct, _ = strconv.ParseUint(strings.Split(net, "/")[1], 10, 8)
		maskLength = int(oct)
	}
	return maskLength
}

func getSessionByKey(key string) *Session {
	s := Session{}
	if CacheSessions {
		s = cachedSessions[key]
	} else {
		Get(&s, "`key` = ?", key)
	}
	if s.ID == 0 {
		return nil
	}
	return &s
}

func getJWT(r *http.Request) string {
	// JWT
	if r.Header.Get("Authorization") == "" {
		return ""
	}
	if !strings.HasPrefix(r.Header.Get("Authorization"), "Bearer") {
		return ""
	}

	jwtToken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	jwtParts := strings.Split(jwtToken, ".")

	if len(jwtParts) != 3 {
		return ""
	}

	jHeader, err := base64.RawURLEncoding.WithPadding(base64.NoPadding).DecodeString(jwtParts[0])
	if err != nil {
		return ""
	}
	jPayload, err := base64.RawURLEncoding.WithPadding(base64.NoPadding).DecodeString(jwtParts[1])
	if err != nil {
		return ""
	}

	header := map[string]interface{}{}
	err = json.Unmarshal(jHeader, &header)
	if err != nil {
		return ""
	}

	// Get data from payload
	payload := map[string]interface{}{}
	err = json.Unmarshal(jPayload, &payload)
	if err != nil {
		return ""
	}

	// Verify issuer
	SSOLogin := false
	if iss, ok := payload["iss"].(string); ok {
		if iss != JWTIssuer {
			accepted := false
			for _, fiss := range AcceptedJWTIssuers {
				if fiss == iss {
					accepted = true
					break
				}
			}
			if !accepted {
				return ""
			}
			SSOLogin = true
		}
	} else {
		return ""
	}

	// verify audience
	if aud, ok := payload["aud"].(string); ok {
		if aud != JWTIssuer {
			return ""
		}
	} else if aud, ok := payload["aud"].([]string); ok {
		accepted := false
		for _, audItem := range aud {
			if audItem == JWTIssuer {
				accepted = true
				break
			}
		}
		if !accepted {
			return ""
		}
	} else {
		return ""
	}

	// if there is no subject, return empty session
	if _, ok := payload["sub"].(string); !ok {
		return ""
	}

	sub := payload["sub"].(string)
	user := User{}
	Get(&user, "username = ?", sub)

	if user.ID == 0 && SSOLogin {
		user := User{
			Username:     sub,
			FirstName:    sub,
			Active:       true,
			Admin:        true,
			RemoteAccess: true,
			Password:     GenerateBase64(64),
		}
		user.Save()
	} else if user.ID == 0 {
		return ""
	}

	session := user.GetActiveSession()
	if session == nil && SSOLogin {
		session = &Session{
			UserID:    user.ID,
			Active:    true,
			LoginTime: time.Now(),
			IP:        GetRemoteIP(r),
		}
		session.GenerateKey()
		session.Save()
	} else if session == nil {
		return ""
	}

	// TODO: verify exp

	// Verify the signature
	alg := "HS256"
	if v, ok := header["alg"].(string); ok {
		alg = v
	}
	if _, ok := header["typ"]; ok {
		if v, ok := header["typ"].(string); !ok || v != "JWT" {
			return ""
		}
	}
	// verify signature
	switch alg {
	case "HS256":
		// TODO: allow third party JWT signature authentication
		hash := hmac.New(sha256.New, []byte(JWT+session.Key))
		hash.Write([]byte(jwtParts[0] + "." + jwtParts[1]))
		token := hash.Sum(nil)
		b64Token := base64.RawURLEncoding.EncodeToString(token)
		if b64Token != jwtParts[2] {
			return ""
		}
	case "RS256":
		if !verifyRSA(jwtToken, SSOLogin) {
			return ""
		}
	default:
		// For now, only support HMAC-SHA256
		return ""
	}

	return session.Key

}

var jwtIssuerCerts = map[[2]string][]byte{}

func getJWTRSAPublicKeySSO(jwtToken *jwt.Token) *rsa.PublicKey {
	iss, err := jwtToken.Claims.GetIssuer()
	if err != nil {
		return nil
	}

	kid, _ := jwtToken.Header["kid"].(string)
	if kid == "" {
		return nil
	}

	if val, ok := jwtIssuerCerts[[2]string{iss, kid}]; ok {
		cert, _ := jwt.ParseRSAPublicKeyFromPEM(val)
		return cert
	}

	res, err := http.Get(iss + "/.well-known/openid-configuration")
	if err != nil {
		return nil
	}

	if res.StatusCode != 200 {
		return nil
	}

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return nil
	}

	obj := map[string]interface{}{}
	err = json.Unmarshal(buf, &obj)
	if err != nil {
		return nil
	}

	crtURL := ""
	if val, ok := obj["jwks_uri"].(string); !ok || val == "" {
		return nil
	} else {
		crtURL = val
	}

	res, err = http.Get(crtURL)
	if err != nil {
		return nil
	}

	if res.StatusCode != 200 {
		return nil
	}

	buf, err = io.ReadAll(res.Body)
	if err != nil {
		return nil
	}

	certObj := map[string][]map[string]string{}
	err = json.Unmarshal(buf, &certObj)
	if err != nil {
		return nil
	}

	if val, ok := certObj["keys"]; !ok || len(val) == 0 {
		return nil
	}

	var cert map[string]string
	for i := range certObj["keys"] {
		if certObj["keys"][i]["kid"] == kid {
			cert = certObj["keys"][i]
			break
		}
	}

	if cert == nil {
		return nil
	}

	N := new(big.Int)
	buf, _ = base64.RawURLEncoding.DecodeString(cert["n"])
	N = N.SetBytes(buf)

	E := new(big.Int)
	buf, _ = base64.RawURLEncoding.DecodeString(cert["e"])
	E = E.SetBytes(buf)
	publicCert := rsa.PublicKey{
		N: N,
		E: int(E.Int64()),
	}

	return &publicCert
}

func getJWTRSAPublicKeyLocal(jwtToken *jwt.Token) *rsa.PublicKey {
	pubKeyPEM, err := os.ReadFile(".jwt-rsa-public.pem")
	if err != nil {
		return nil
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyPEM)
	if err != nil {
		return nil
	}

	return pubKey
}

func verifyRSA(token string, SSOLogin bool) bool {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		var pubKey *rsa.PublicKey

		if SSOLogin {
			pubKey = getJWTRSAPublicKeySSO(jwtToken)
		} else {
			pubKey = getJWTRSAPublicKeyLocal(jwtToken)
		}

		if pubKey == nil {
			return nil, fmt.Errorf("Unable to load local public key")
		}

		return pubKey, nil
	})
	if err != nil {
		return false
	}

	_, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return false
	}

	return true
}

func getSession(r *http.Request) string {
	// First, try JWT
	if val := getJWT(r); val != "" {
		return val
	}

	if r.URL.Query().Get("access-token") != "" {
		r.Header.Add("Authorization", "Bearer "+r.URL.Query().Get("access-token"))
		if val := getJWT(r); val != "" {
			return val
		}
	}

	// Then try session
	key, err := r.Cookie("session")
	if err == nil && key != nil {
		return key.Value
	}
	if r.Method == "GET" && r.FormValue("session") != "" {
		return r.FormValue("session")
	}
	if r.Method != "GET" {
		err := r.ParseMultipartForm(2 << 10)
		if err != nil {
			r.ParseForm()
		}
		if r.FormValue("session") != "" {
			return r.FormValue("session")
		}
	}

	return ""
}

// GetRemoteIP is a function that returns the IP for a remote
// user from a request
func GetRemoteIP(r *http.Request) string {
	ips := r.Header.Get("X-Forwarded-For")

	splitIps := strings.Split(ips, ",")

	if ips != "" {
		// trim IP list
		for i := range splitIps {
			splitIps[i] = strings.TrimSpace(splitIps[i])
		}

		// get last IP in list since ELB prepends other user defined IPs, meaning the last one is the actual client IP.
		netIP := net.ParseIP(splitIps[0])
		if netIP != nil {
			return netIP.String()
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}

	netIP := net.ParseIP(ip)
	if netIP != nil {
		ip := netIP.String()
		if ip == "::1" {
			return "127.0.0.1"
		}
		return ip
	}

	return r.RemoteAddr
}

// GetHostName is a function that returns the host name from a request
func GetHostName(r *http.Request) string {
	host := r.Header.Get("X-Forwarded-Host")
	if host != "" {
		return host
	}
	return r.Host
}

// GetSchema is a function that returns the schema for a request (http, https)
func GetSchema(r *http.Request) string {
	schema := r.Header.Get("X-Forwarded-Proto")
	if schema != "" {
		return schema
	}

	if r.URL.Scheme != "" {
		return r.URL.Scheme
	}

	if r.TLS != nil {
		return "https"
	}
	return "http"
}

func verifyPassword(hash string, plain string) error {
	password := []byte(plain + Salt)
	hashedPassword := []byte(hash)
	if len(hashedPassword) > 72 {
		hashedPassword = hashedPassword[:72]
	}
	if len(password) > 72 {
		password = password[:72]
	}
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

// sanitizeFileName is a function to sanitize file names to pretect
// from path traversal attacks using ../
func sanitizeFileName(v string) string {
	return path.Clean(strings.ReplaceAll(v, "../", ""))
}
