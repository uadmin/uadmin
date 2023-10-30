package uadmin

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"time"
)

// Session !
type Session struct {
	Model
	Key        string
	User       User `uadmin:"filter"`
	UserID     uint
	LoginTime  time.Time
	LastLogin  time.Time
	Active     bool   `uadmin:"filter"`
	IP         string `uadmin:"filter"`
	PendingOTP bool   `uadmin:"filter"`
	ExpiresOn  *time.Time
}

// String return string
func (s Session) String() string {
	return s.Key
}

// Save !
func (s *Session) Save() {
	u := s.User
	s.User = User{}

	// Verify required date fields
	if s.LoginTime.IsZero() {
		s.LoginTime = time.Now()
	}
	if s.LastLogin.IsZero() {
		s.LastLogin = time.Now()
	}
	Save(s)
	s.User = u
	if CacheSessions {
		cachedSessionsMutex.Lock()         // Lock the mutex in order to protect from concurrent writes
		defer cachedSessionsMutex.Unlock() // Ensure the mutex is unlocked when the function exits
		if s.Active {
			Preload(s)
			cachedSessions[s.Key] = *s
		} else {
			delete(cachedSessions, s.Key)
		}
	}
}

// GenerateKey !
func (s *Session) GenerateKey() {
	session := Session{}
	hash := md5.New()
	hash.Write([]byte(fmt.Sprint(s.UserID)))
	userID := hash.Sum(nil)
	for {
		s.Key = GenerateBase64(102) + base64.URLEncoding.EncodeToString(userID)
		s.Key = s.Key[:124]
		Get(&session, "`key` = ?", s.Key)
		if session.ID == 0 {
			break
		}
	}
}

// Logout deactivates a session
func (s *Session) Logout() {
	s.Active = false
	s.Save()
}

// HideInDashboard to return false and auto hide this from dashboard
func (Session) HideInDashboard() bool {
	return true
}

func loadSessions() {
	if !CacheSessions {
		return
	}
	cachedSessionsMutex.Lock()         // Lock the mutex in order to protect from concurrent writes
	defer cachedSessionsMutex.Unlock() // Ensure the mutex is unlocked when the function exits

	sList := []Session{}
	Filter(&sList, "`active` = ? AND (expires_on IS NULL OR expires_on > ?)", true, time.Now())
	cachedSessions = map[string]Session{}
	for _, s := range sList {
		Preload(&s)
		Preload(&s.User)
		cachedSessions[s.Key] = s
	}
}
