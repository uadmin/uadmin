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
	DashboardMenu   DashboardMenu `gorm:"ForeignKey:DashboardMenuID" required:"true" filter:"true" uadmin:"filter"`
	DashboardMenuID uint          `fk:"true" displayName:"DashboardMenu"`
	User            User          `gorm:"ForeignKey:UserID" required:"true" filter:"true" uadmin:"filter"`
	UserID          uint          `fk:"true" displayName:"User"`
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
