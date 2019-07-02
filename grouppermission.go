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
	Read            bool `uadmin:"filter"`
	Add             bool `uadmin:"filter"`
	Edit            bool `uadmin:"filter"`
	Delete          bool `uadmin:"filter"`
	Approval        bool `uadmin:"filter"`
}

func (g GroupPermission) String() string {
	return fmt.Sprint(g.ID)
}

// HideInDashboard to return false and auto hide this from dashboard
func (GroupPermission) HideInDashboard() bool {
	return true
}
