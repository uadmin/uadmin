package models

import (
	"github.com/uadmin/uadmin"
)

// DocumentGroup !
type DocumentGroup struct {
	uadmin.Model
	Group      uadmin.UserGroup
	GroupID    uint
	Document   Document
	DocumentID uint
	Read       bool
	Add        bool
	Edit       bool
	Delete     bool
}

// DocumentGroup function that returns string value
func (d *DocumentGroup) String() string {

	// Gives access to the fields in another model
	uadmin.Preload(d)

	// Returns the GroupName from the Group model
	return d.Group.GroupName
}

// HideInDashboard !
func (DocumentGroup) HideInDashboard() bool {
	return true
}
