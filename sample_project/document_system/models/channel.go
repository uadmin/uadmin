package models

import (
	"github.com/uadmin/uadmin"
)

// Channel !
type Channel struct {
	uadmin.Model
	Name string `uadmin:"required"`
}
