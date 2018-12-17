package uadmin

import (
	"math/big"

	"crypto/rand"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// CookieTimeout is the timeout of a login cookie in seconds
var CookieTimeout = -1

// Salt is extra salt added to password hashing
var Salt = ""

// bcryptDiff
var bcryptDiff = 12

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
	Get(&s, "`key` = ?", key)
	if isValidSession(&s) {
		return &s
	}
	return nil
}

func isValidSession(s *Session) bool {
	if s != nil && s.ID != 0 {
		if s.Active && !s.PendingOTP && (s.ExpiresOn == nil || s.ExpiresOn.After(time.Now())) {
			Get(&s.User, "id = ?", s.UserID)
			if s.User.Active && (s.User.ExpiresOn == nil || s.User.ExpiresOn.After(time.Now())) {
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
func Login(r *http.Request, username string, password string) (*User, bool) {
	// Get the user from DB
	user := User{}
	Get(&user, "username = ?", username)
	if user.ID == 0 {
		return nil, false
	}
	s := user.Login(password, "")
	if s != nil && s.ID != 0 {
		if s.Active && (s.ExpiresOn == nil || s.ExpiresOn.After(time.Now())) {
			Get(&s.User, "id = ?", s.UserID)
			if s.User.Active && (s.User.ExpiresOn == nil || s.User.ExpiresOn.After(time.Now())) {
				return &s.User, s.User.OTPRequired
			}
		}
	}
	return nil, false
}

// Login2FA login using username, password and otp for users with OTPRequired = true
func Login2FA(r *http.Request, username string, password string, otpPass string) *User {
	u, otpRequired := Login(r, username, password)
	if u != nil {
		if !otpRequired || u.VerifyOTP(otpPass) {
			return u
		}
	}
	return nil
}

// Logout logs out a user
func Logout(r *http.Request) {
	s := getSessionFromRequest(r)
	s.Active = false
	s.Save()
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
