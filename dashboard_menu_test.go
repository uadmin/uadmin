package uadmin

import (
	"testing"
)

// DashboardMenu is a unit testing function for DashboardMenu.String() function
func TestDashboardMenu(t *testing.T) {
	d := DashboardMenu{
		MenuName: "Hello",
	}
	if d.String() != "Hello" {
		t.Errorf("DashboardMenu.String returned wrong string. Expected %s, got %s", "Hello", d.String())
	}
}
