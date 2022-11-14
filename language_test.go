package uadmin

// TestLanguage is for testing Language struct
func (t *UAdminTests) TestLanguage() {
	lang := Language{
		Code: "ts",
	}
	if lang.String() != "ts" {
		t.Errorf("Language.String didn't return a valid value. Expected (%s) got (%s).", "ts", lang.String())
	}
}
