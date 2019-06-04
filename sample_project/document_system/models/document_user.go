package models

import (
	"github.com/uadmin/uadmin"
)

// DocumentUser !
type DocumentUser struct {
	uadmin.Model
	User       uadmin.User
	UserID     uint
	Document   Document
	DocumentID uint
	Read       bool
	Add        bool
	Edit       bool
	Delete     bool
}

// DocumentUser function that returns string value
func (d *DocumentUser) String() string {

	// Gives access to the fields in another model
	uadmin.Preload(d)

	// Returns the full name from the User model
	return d.User.String()
}

// HideInDashboard !
func (DocumentUser) HideInDashboard() bool {
	return true
}
