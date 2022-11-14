package uadmin

// DashboardMenu is a unit testing function for DashboardMenu.String() function
func (t *UAdminTests) TestDashboardMenu() {
	d := DashboardMenu{
		MenuName: "Hello",
	}
	if d.String() != "Hello" {
		t.Errorf("DashboardMenu.String returned wrong string. Expected %s, got %s", "Hello", d.String())
	}
}
