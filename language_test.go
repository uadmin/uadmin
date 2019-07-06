package uadmin

import (
	"testing"
)

// TestLanguage is for testing Language struct
func TestLanguage(t *testing.T) {
	lang := Language{
		Code: "ts",
	}
	if lang.String() != "ts" {
		t.Errorf("Language.String didn't return a valid value. Expected (%s) got (%s).", "ts", lang.String())
	}
}
