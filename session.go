package uadmin

import (
	"time"
)

// Session !
type Session struct {
	Model
	Key        string
	User       User `gorm:"ForeignKey:UserID" uadmin:"filter"`
	UserID     uint `fk:"true" displayName:"User"`
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
}

// GenerateKey !
func (s *Session) GenerateKey() {
	session := Session{}
	for {
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
	Save(s)
}

// HideInDashboard to return false and auto hide this from dashboard
func (Session) HideInDashboard() bool {
	return true
}
