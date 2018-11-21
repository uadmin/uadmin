package models

import "github.com/uadmin/uadmin"

// Category model ...
type Category struct {
	uadmin.Model
	Name string `uadmin:"required"`
	Icon string `uadmin:"image"`
}
