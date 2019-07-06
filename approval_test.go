package uadmin

import (
	"testing"
)

// TestApprovalStruct is to test Approval struct
func TestApprovalStruct(t *testing.T) {
	app := Approval{
		ModelName:  "test",
		ColumnName: "column",
		ModelPK:    100,
	}

	if app.String() != "test.column 100" {
		t.Errorf("Approval.String didn't return valid value. Got(%s) Expected (%s)", app.String(), "test.column 100")
	}
}
