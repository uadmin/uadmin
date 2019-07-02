package uadmin

import (
	"time"
)

// Model is the standard struct to be embedded
// in any other struct to make it a model for uadmin
type Model struct {
	ID        uint       `gorm:"primary_key"`
	DeletedAt *time.Time `sql:"index"`
}
