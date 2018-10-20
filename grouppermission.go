package uadmin

import (
	"fmt"
)

// GroupPermission !
type GroupPermission struct {
	Model
	DashboardMenu   DashboardMenu `gorm:"ForeignKey:DashboardMenuID" required:"true" filter:"true"`
	DashboardMenuID uint          `fk:"true" displayName:"DashboardMenu"`
	UserGroup       UserGroup     `gorm:"ForeignKey:UserGroupID" required:"true" filter:"true"`
	UserGroupID     uint          `fk:"true" displayName:"UserGroup"`
	Read            bool
	Add             bool
	Edit            bool
	Delete          bool
}

func (g GroupPermission) String() string {
	return fmt.Sprint(g.ID)
}

// HideInDashboard to return false and auto hide this from dashboard
func (GroupPermission) HideInDashboard() bool {
	return true
}
