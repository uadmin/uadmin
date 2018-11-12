package uadmin

import (
	"fmt"
)

// GroupPermission !
type GroupPermission struct {
	Model
	DashboardMenu   DashboardMenu `uadmin:"required;filter"`
	DashboardMenuID uint
	UserGroup       UserGroup `uadmin:"required;filter"`
	UserGroupID     uint
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
