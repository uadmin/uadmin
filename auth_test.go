package uadmin

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// TestGenerateBase64 is a unit testing function for GenerateBase64() function
func TestGenerateBase64(t *testing.T) {
	examples := []struct {
		length int
	}{
		{0},
		{1},
		{10},
	}

	for _, e := range examples {
		code := GenerateBase64(e.length)
		if len(code) != e.length {
			t.Errorf("length of GenerateBase64(%d) = %d != %d", e.length, len(code), e.length)
		}
	}
}

// TestGenerateBase32 is a unit testing function for GenerateBase32() function
func TestGenerateBase32(t *testing.T) {
	examples := []struct {
		length int
	}{
		{0},
		{1},
		{10},
	}

	for _, e := range examples {
		code := GenerateBase32(e.length)
		if len(code) != e.length {
			t.Errorf("length of GenerateBase32(%d) = %d != %d", e.length, len(code), e.length)
		}
	}
}

// TestHashPass is a unit testing function for hashPass() function
func TestHashPass(t *testing.T) {
	examples := []struct {
		pass string
	}{
		{"1234"},
		{"abc123"},
		{"password"},
		{"password1"},
		{"Password1"},
		{" Password1 "},
		{"Pass 123"},
		{"Pass 123!"},
		{"Pass 123! "},
		{"كلمة السر 123! "},
		{GenerateBase64(10)},
		{GenerateBase64(20)},
		{GenerateBase64(30)},
		{GenerateBase64(40)},
		{GenerateBase64(50)},
		{GenerateBase64(60)},
		{GenerateBase64(70)},
	}

	bcryptDiff = 5

	for _, e := range examples {
		code := hashPass(e.pass)
		if bcrypt.CompareHashAndPassword([]byte(code), []byte(e.pass+Salt)) != nil {
			t.Errorf("hashPass(\"%s\") invalid denied password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte(code), []byte("1"+e.pass+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte(code), []byte("a"+e.pass+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte(code), []byte(" "+e.pass+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte(code), []byte(e.pass+" "+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte("234"+code), []byte(e.pass+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
	}

	for _, e := range examples {
		code := hashPass(e.pass)
		if bcrypt.CompareHashAndPassword([]byte(code), []byte(e.pass+Salt)) != nil {
			t.Errorf("hashPass(\"%s\") invalid denied password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte(code), []byte(Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte(code), []byte("1"+e.pass+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte(code), []byte("a"+e.pass+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte(code), []byte(" "+e.pass+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte(code), []byte(e.pass+" "+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte(code), []byte(e.pass+" "+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte("234"+code), []byte(e.pass+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
	}

}

// TestIsAuthenticated is a unit testing function for IsAuthenticated() function
func TestIsAuthenticated(t *testing.T) {
	// Setup
	yesterday := time.Now().AddDate(0, 0, -1)
	tomorrow := time.Now().AddDate(0, 0, 1)

	tx := db.Begin()

	// deactivated user
	u1 := User{}
	u1.FirstName = "u1"
	u1.Username = "u1"
	u1.Password = "u1"
	u1.Active = false
	u1.Admin = false
	u1.RemoteAccess = false
	u1.ExpiresOn = nil
	tx.Save(&u1)

	// expired user
	u2 := User{}
	u2.Username = "u2"
	u2.Password = "u2"
	u2.Active = true
	u2.Admin = false
	u2.RemoteAccess = false
	u2.ExpiresOn = &yesterday
	tx.Save(&u2)

	// user with expiry in the future
	u3 := User{}
	u3.Username = "u3"
	u3.Password = "u3"
	u3.Active = true
	u3.Admin = false
	u3.RemoteAccess = false
	u3.ExpiresOn = &tomorrow
	tx.Save(&u3)

	s1 := Session{
		UserID:     1,
		Active:     true,
		PendingOTP: false,
		ExpiresOn:  nil,
	}
	s1.GenerateKey()
	tx.Save(&s1)

	s2 := Session{
		UserID:     1,
		Active:     false,
		PendingOTP: false,
		ExpiresOn:  nil,
	}
	s2.GenerateKey()
	tx.Save(&s2)

	s3 := Session{
		UserID:     1,
		Active:     true,
		PendingOTP: true,
		ExpiresOn:  nil,
	}
	s3.GenerateKey()
	tx.Save(&s3)

	s4 := Session{
		UserID:     1,
		Active:     true,
		PendingOTP: false,
		ExpiresOn:  &yesterday,
	}
	s4.GenerateKey()
	tx.Save(&s4)

	s5 := Session{
		UserID:     1,
		Active:     true,
		PendingOTP: false,
		ExpiresOn:  &tomorrow,
	}
	s5.GenerateKey()
	tx.Save(&s5)

	s6 := Session{
		UserID:     u1.ID,
		Active:     true,
		PendingOTP: false,
		ExpiresOn:  nil,
	}
	s6.GenerateKey()
	tx.Save(&s6)

	s7 := Session{
		UserID:     u2.ID,
		Active:     true,
		PendingOTP: false,
		ExpiresOn:  nil,
	}
	s7.GenerateKey()
	tx.Save(&s7)

	s8 := Session{
		UserID:     u3.ID,
		Active:     true,
		PendingOTP: false,
		ExpiresOn:  nil,
	}
	s8.GenerateKey()
	tx.Save(&s8)

	tx.Commit()
	loadSessions()
	loadPermissions()

	examples := []struct {
		r *http.Request
		s *Session
	}{
		{httptest.NewRequest("GET", "/", nil), &s1},
		{httptest.NewRequest("GET", "/?session="+s1.Key, nil), &s1},
		{httptest.NewRequest("GET", "/", nil), &s1},
		{httptest.NewRequest("GET", "/", nil), nil},
		{httptest.NewRequest("GET", "/?session="+s2.Key, nil), nil},
		{httptest.NewRequest("GET", "/?session="+s3.Key, nil), nil},
		{httptest.NewRequest("GET", "/?session="+s4.Key, nil), nil},
		{httptest.NewRequest("GET", "/?session="+s5.Key, nil), &s5},
		{httptest.NewRequest("GET", "/?session="+s6.Key, nil), nil},
		{httptest.NewRequest("GET", "/?session="+s7.Key, nil), nil},
		{httptest.NewRequest("GET", "/?session="+s8.Key, nil), &s8},
	}

	// Prepare requests with session data
	cookie := http.Cookie{}
	cookie.Name = "session"
	cookie.Value = s1.Key
	examples[0].r.AddCookie(&cookie)
	examples[2].r.Form = url.Values{}
	examples[2].r.Form.Add("session", s1.Key)

	for _, e := range examples {
		tempS := IsAuthenticated(e.r)
		if (tempS == nil && e.s != nil) || (tempS != nil && e.s == nil) {
			t.Errorf("Invalid output from IsAuthenticated: %v, expected %v", tempS, e.s)
		} else if (tempS != nil && e.s != nil) && (tempS.ID != e.s.ID) {
			t.Errorf("Invalid session ID from IsAuthenticated: %v, expected %v", tempS.ID, e.s.ID)
		}
	}

	// Clean up
	Delete(s1)
	Delete(s2)
	Delete(s3)
	Delete(s4)
	Delete(s5)
	Delete(s6)
	Delete(s7)
	Delete(s8)
	Delete(u1)
	Delete(u2)
	Delete(u3)
}

// TestGetUserFromRequest is a unit testing function for GetUserFromRequest() function
func TestGetUserFromRequest(t *testing.T) {
	s1 := Session{
		UserID:     1,
		Active:     true,
		PendingOTP: false,
		ExpiresOn:  nil,
	}
	s1.GenerateKey()
	s1.Save()

	admin := User{}
	Get(&admin, "id=?", 1)

	examples := []struct {
		r *http.Request
		s *User
	}{
		{httptest.NewRequest("GET", "/", nil), &admin},
		{httptest.NewRequest("GET", "/?session="+s1.Key, nil), &admin},
		{httptest.NewRequest("GET", "/", nil), &admin},
		{httptest.NewRequest("GET", "/", nil), nil},
	}

	// Prepare requests with session data
	cookie := http.Cookie{}
	cookie.Name = "session"
	cookie.Value = s1.Key
	examples[0].r.AddCookie(&cookie)
	examples[2].r.Form = url.Values{}
	examples[2].r.Form.Add("session", s1.Key)

	for _, e := range examples {
		tempU := GetUserFromRequest(e.r)
		if (tempU == nil && e.s != nil) || (tempU != nil && e.s == nil) {
			t.Errorf("Invalid output from GetUserFromRequest: %v, expected %v", tempU, e.s)
		} else if (tempU != nil && e.s != nil) && (tempU.ID != e.s.ID) {
			t.Errorf("Invalid user ID from GetUserFromRequest: %v, expected %v", tempU.ID, e.s.ID)
		}
	}

	Delete(s1)
}

// TestLogin is a unit testing function for Login() function
func TestLogin(t *testing.T) {
	// Setup
	yesterday := time.Now().AddDate(0, 0, -1)
	tomorrow := time.Now().AddDate(0, 0, 1)

	// deactivated user
	u1 := User{}
	u1.FirstName = "u1"
	u1.Username = "u1"
	u1.Password = "u1"
	u1.Active = false
	u1.Admin = false
	u1.RemoteAccess = false
	u1.ExpiresOn = nil
	u1.Save()

	// expired user
	u2 := User{}
	u2.FirstName = "u2"
	u2.Username = "u2"
	u2.Password = "u2"
	u2.Active = true
	u2.Admin = false
	u2.RemoteAccess = false
	u2.ExpiresOn = &yesterday
	u2.Save()

	// user with expiry in the future
	u3 := User{}
	u3.FirstName = "u3"
	u3.Username = "u3"
	u3.Password = "u3"
	u3.Active = true
	u3.Admin = false
	u3.RemoteAccess = false
	u3.ExpiresOn = &tomorrow
	u3.Save()

	// user OTP required
	u4 := User{}
	u4.FirstName = "u4"
	u4.Username = "u4"
	u4.Password = "u4"
	u4.Active = true
	u4.Admin = false
	u4.RemoteAccess = false
	u4.ExpiresOn = nil
	u4.OTPRequired = true
	u4.Save()

	admin := User{}
	Get(&admin, "id=?", 1)

	examples := []struct {
		username string
		password string
		u        *User
	}{
		{"", "admin", nil},
		{"admin", "", nil},
		{"admin", "admin", &admin},
		{"admin", GenerateBase64(10), nil},
		{"u1", "u1", nil},
		{"", "u1", nil},
		{"u1", "", nil},
		{"u1", GenerateBase64(10), nil},
		{"u2", "u2", nil},
		{"", "u2", nil},
		{"u2", "", nil},
		{"u2", GenerateBase64(10), nil},
		{"u3", "u3", &u3},
		{"", "u3", nil},
		{"u3", "", nil},
		{"u3", GenerateBase64(10), nil},
		{"u4", "u4", &u4},
		{"", "u4", nil},
		{"u4", "", nil},
		{"u4", GenerateBase64(10), nil},
	}
	r := httptest.NewRequest("GET", "/", nil)

	for _, e := range examples {
		tempU, otpRequired := Login(r, e.username, e.password)
		if (tempU == nil && e.u != nil) || (tempU != nil && e.u == nil) {
			t.Errorf("Invalid output from Login: %v, expected %v", tempU, e.u)
		} else if (tempU != nil && e.u != nil) && (tempU.User.ID != e.u.ID) {
			t.Errorf("Invalid user ID from Login: %v, expected %v", tempU.User.ID, e.u.ID)
		} else if (e.u != nil) && (otpRequired != e.u.OTPRequired) {
			t.Errorf("Invalid OTPRequired output from Login: %v, expected %v", otpRequired, e.u.OTPRequired)
		}
	}
	Delete(u1)
	Delete(u2)
	Delete(u3)
	Delete(u4)
}

// TestLogin2FA is a unit testing function for Login2FA() function
func TestLogin2FA(t *testing.T) {
	// Setup

	// user with otp required
	u1 := User{}
	u1.FirstName = "u1"
	u1.Username = "u1"
	u1.Password = "u1"
	u1.Active = true
	u1.Admin = false
	u1.RemoteAccess = false
	u1.ExpiresOn = nil
	u1.OTPRequired = true
	u1.Save()

	examples := []struct {
		username   string
		password   string
		otp        string
		u          *User
		PendingOTP bool
	}{
		{"u1", "u1", "", &u1, true},
		{"", "u1", "", nil, false},
		{"u1", "", "", nil, false},
		{"u1", GenerateBase64(10), "", nil, false},
		{"u1", "u1", "000000", &u1, true},
		{"", "u1", "000000", nil, false},
		{"u1", "", "000000", nil, false},
		{"u1", GenerateBase64(10), "000000", nil, false},
		{"u1", "u1", u1.GetOTP(), &u1, false},
		{"", "u1", u1.GetOTP(), nil, false},
		{"u1", "", u1.GetOTP(), nil, false},
		{"u1", GenerateBase64(10), u1.GetOTP(), nil, false},
	}
	r := httptest.NewRequest("GET", "/", nil)

	for i, e := range examples {
		tempU := Login2FA(r, e.username, e.password, e.otp)
		if (tempU == nil && e.u != nil) || (tempU != nil && e.u == nil) {
			t.Errorf("Invalid output from Login: %v, expected %v in test %d", tempU, e.u, i)
		} else if (tempU != nil && e.u != nil) && (tempU.User.ID != e.u.ID) {
			t.Errorf("Invalid user ID from Login: %v, expected %v in test %d", tempU.User.ID, e.u.ID, i)
		} else if tempU != nil && tempU.PendingOTP != e.PendingOTP {
			t.Errorf("Invalid pending otp status Got: %v, expected %v in test %d", tempU.PendingOTP, e.PendingOTP, i)
		}
	}
	Delete(u1)
}

// TestLogout is a unit testing function for Logout() function
func TestLogout(t *testing.T) {
	// Setup
	r := httptest.NewRequest("GET", "/", nil)
	admin, _ := Login(r, "admin", "admin")
	s1 := admin.User.GetActiveSession()

	c := http.Cookie{}
	c.Name = "session"
	c.Value = s1.Key

	r.AddCookie(&c)

	Logout(r)
	s2 := admin.User.GetActiveSession()
	if s2 != nil {
		t.Errorf("Logout didn't deactivate the user's active session")
	}

	Delete(s1)
}

func TestValidateIP(t *testing.T) {
	examples := []struct {
		ip     string
		allow  string
		block  string
		result bool
	}{
		{"192.168.1.1:1234", "*", "", true},
		{"192.168.1.1:1234", "*", "192.168.1.1", false},
		{"192.168.1.1:1234", "*", "192.168.1.0/24", false},
		{"192.168.1.1:1234", "192.168.1.1", "192.168.1.0/24", true},
		{"192.168.1.1:1234", "192.168.1.0/22", "192.168.1.0/24", false},
		{"192.168.1.1:1234", "192.168.1.0/24", "*", true},
		{"192.168.1.1:1234", "192.168.1.0/24,2400::/64", "*", true},
		{"192.168.1.56:1234", "*", "", true},
		{"[2400::1]:1234", "*", "", true},
		{"[2400::1]:1234", "2400::/64", "", true},
		{"[2400::1]:1234", "2400::1", "", true},
		{"[2400::1]:1234", "2400::/64,192.168.1.1", "", true},
		{"[2400::1]:1234", "192.168.1.1,2400::/64", "", true},
		{"[2400::1]:1234", "2401::/64", "", false},
		{"[2400::1]:1234", "*", "2400::/64", false},
		{"[2400::1]:1234", "*", "2400::1", false},
		{"[2400::1]:1234", "2400::/64", "2400::/80", false},
		{"[2400::1]:1234", "2400::1", "2400::/64", true},
	}
	var r http.Request
	for _, e := range examples {
		r = http.Request{RemoteAddr: e.ip}
		if ValidateIP(&r, e.allow, e.block) != e.result {
			t.Errorf("Invalid output from ValidateIP: %v, expected %v for %s in allow: (%s), block:(%s)", !e.result, e.result, e.ip, e.allow, e.block)
		}
	}
}

// TestGetSessionByKey is a unit testing function for getSessionByKey() function
func TestGetSessionByKey(t *testing.T) {
	s1 := Session{
		UserID:     1,
		Active:     true,
		PendingOTP: false,
		ExpiresOn:  nil,
	}
	s1.GenerateKey()
	s1.Save()

	s2 := getSessionByKey(s1.Key)
	if s2 != nil && (s1.ID != s2.ID) {
		t.Errorf("getSessionByKey didn't return the correct session")
	}
	s3 := getSessionByKey("")
	if s3 != nil {
		t.Errorf("getSessionByKey returned an invalid session")
	}
	Delete(s1)
}

// TestGetSession is a unit testing function for getSession() function
func TestGetSession(t *testing.T) {
	s1 := Session{
		UserID:     1,
		Active:     true,
		PendingOTP: false,
		ExpiresOn:  nil,
	}
	s1.GenerateKey()
	s1.Save()

	examples := []struct {
		r   *http.Request
		key string
	}{
		{httptest.NewRequest("GET", "/", nil), s1.Key},
		{httptest.NewRequest("GET", "/?session="+s1.Key, nil), s1.Key},
		{httptest.NewRequest("POST", "/", nil), s1.Key},
		{httptest.NewRequest("GET", "/", nil), ""},
		{httptest.NewRequest("GET", "/?session=", nil), ""},
		{httptest.NewRequest("POST", "/", nil), ""},
	}

	// Prepare requests with session data
	cookie := http.Cookie{}
	cookie.Name = "session"
	cookie.Value = s1.Key
	examples[0].r.AddCookie(&cookie)
	examples[2].r.Form = url.Values{}
	examples[2].r.Form.Add("session", s1.Key)

	for i, e := range examples {
		if getSession(e.r) != e.key {
			t.Errorf("getSession didn't return the correct session key=%s expected %s at %d", getSession(e.r), e.key, i)
		}
	}

	Delete(s1)
}
