package uadmin

// UserGroup !
type UserGroup struct {
	Model
	GroupName string `uadmin:"filter"`
}

func (u UserGroup) String() string {
	return u.GroupName
}

// HasAccess !
func (u *UserGroup) HasAccess(modelName string) GroupPermission {
	up := GroupPermission{}
	dm := DashboardMenu{}
	Get(&dm, "url = ?", modelName)
	Get(&up, "user_group_id = ? AND dashboard_menu_id = ?", u.ID, dm.ID)
	return up
}
