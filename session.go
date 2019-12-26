package uadmin

import (
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
	Save(s)
	if CacheSessions {
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
	for {
		// TODO: Increase the session length to 124 and add 4 bytes for User.ID
		s.Key = GenerateBase64(24)
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
	sList := []Session{}
	Filter(&sList, "active = ? AND (expires_on IS NULL OR expires_on > ?)", true, time.Now())
	cachedSessions = map[string]Session{}
	for _, s := range sList {
		Preload(&s)
		Preload(&s.User)
		cachedSessions[s.Key] = s
	}
}
