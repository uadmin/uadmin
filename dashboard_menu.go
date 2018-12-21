package uadmin

// DashboardMenu !
type DashboardMenu struct {
	Model
	MenuName string `uadmin:"required;list_exclude;multilingual;filter"`
	URL      string `uadmin:"required"`
	ToolTip  string
	Icon     string `uadmin:"image"`
	Cat      string `uadmin:"filter"`
	Hidden   bool   `uadmin:"filter"`
}

func (m DashboardMenu) String() string {
	return Translate(m.MenuName, "", true)
}
