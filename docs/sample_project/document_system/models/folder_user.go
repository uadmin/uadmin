package models

import (
	"github.com/uadmin/uadmin"
)

// FolderUser !
type FolderUser struct {
	uadmin.Model
	User     uadmin.User
	UserID   uint
	Folder   Folder
	FolderID uint
	Read     bool
	Add      bool
	Edit     bool
	Delete   bool
}

// FolderUser function that returns string value
func (f *FolderUser) String() string {

	// Gives access to the fields in another model
	uadmin.Preload(f)

	// Returns the full name from the User model
	return f.User.String()
}

// HideInDashboard !
func (FolderUser) HideInDashboard() bool {
	return true
}
