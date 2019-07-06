package uadmin

import (
	"testing"
)

// TestGroupPermission is for testing GroupPermission struct
func TestGroupPermission(t *testing.T) {
	gp := GroupPermission{}
	gp.ID = 100
	if gp.String() != "100" {
		t.Errorf("GroupPermission.String didn't return a valid value. Expected (%s) got (%s).", "100", gp.String())
	}
}
