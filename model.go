package uadmin

import (
	"gorm.io/gorm"
)

// Model is the standard struct to be embedded
// in any other struct to make it a model for uadmin
type Model struct {
	ID        uint           `gorm:"primary_key"`
	DeletedAt gorm.DeletedAt `sql:"index" json:",omitempty"`
}
