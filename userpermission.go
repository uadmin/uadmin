package uadmin

import (
	"fmt"
)

var cacheUserPerms []UserPermission
var cacheGroupPerms []GroupPermission
var cachedModels []DashboardMenu

// UserPermission !
type UserPermission struct {
	Model
	DashboardMenu   DashboardMenu `uadmin:"filter"`
	DashboardMenuID uint          ``
	User            User          `uadmin:"filter"`
	UserID          uint          ``
	Read            bool          `uadmin:"filter"`
	Add             bool          `uadmin:"filter"`
	Edit            bool          `uadmin:"filter"`
	Delete          bool          `uadmin:"filter"`
	Approval        bool          `uadmin:"filter"`
}

func (u UserPermission) String() string {
	return fmt.Sprint(u.ID)
}

// HideInDashboard to return false and auto hide this from dashboard
func (UserPermission) HideInDashboard() bool {
	return true
}

func loadPermissions() {
	cacheUserPerms = []UserPermission{}
	cacheGroupPerms = []GroupPermission{}
	cachedModels = []DashboardMenu{}
	All(&cacheUserPerms)
	All(&cacheGroupPerms)
	All(&cachedModels)
}
