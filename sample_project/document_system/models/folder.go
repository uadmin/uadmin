package models

import (
	"github.com/uadmin/uadmin"
)

// Folder !
type Folder struct {
	uadmin.Model
	Name     string
	Parent   *Folder
	ParentID uint
}

// Returns the Name field
func (f Folder) String() string {
	return f.Name
}
