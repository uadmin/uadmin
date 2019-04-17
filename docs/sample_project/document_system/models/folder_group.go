package models

import (
	"github.com/uadmin/uadmin"
)

// FolderGroup !
type FolderGroup struct {
	uadmin.Model
	Group    uadmin.UserGroup
	GroupID  uint
	Folder   Folder
	FolderID uint
	Read     bool
	Add      bool
	Edit     bool
	Delete   bool
}

// FolderGroup function that returns string value
func (f *FolderGroup) String() string {

	// Gives access to the fields in another model
	uadmin.Preload(f)

	// Returns the GroupName from the Group model
	return f.Group.GroupName
}

// HideInDashboard !
func (FolderGroup) HideInDashboard() bool {
	return true
}
