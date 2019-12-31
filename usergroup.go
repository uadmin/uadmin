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
	Trail(WARNING, "UserGroup.HasAccess will be deprecated in version 0.6.0. Use User.GetAccess instead.")
	return u.hasAccess(modelName)
}

// hasAccess !
func (u *UserGroup) hasAccess(modelName string) GroupPermission {
	up := GroupPermission{}
	dm := DashboardMenu{}
	if CachePermissions {
		modelID := uint(0)
		for _, m := range cachedModels {
			if m.URL == modelName {
				modelID = m.ID
				break
			}
		}
		for _, g := range cacheGroupPerms {
			if g.UserGroupID == u.ID && g.DashboardMenuID == modelID {
				up = g
				break
			}
		}
	} else {
		Get(&dm, "url = ?", modelName)
		Get(&up, "user_group_id = ? AND dashboard_menu_id = ?", u.ID, dm.ID)
	}
	return up
}
