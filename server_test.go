package uadmin

import (
	"os"
	"testing"
)

func setupFunction() {
	Register(
		TestStruct1{},
	)

	Port = 5000
	EmailFrom = "uadmin@example.com"
	EmailPassword = "password"
	EmailUsername = "uadmin@example.com"
	EmailSMTPServer = "localhost"
	EmailSMTPServerPort = 2525

	go StartServer()
	go startEmailServer()
}

func teardownFunction() {
	// Remove Generated Files
	os.Remove("uadmin.db")
	os.Remove(".key")
	os.Remove(".salt")
	os.Remove(".uproj")
	os.Remove(".bindip")

	// Delete temp media file
	os.RemoveAll("./media")
	os.RemoveAll("./static/i18n")
}

func TestMain(t *testing.M) {
	teardownFunction()
	setupFunction()
	retCode := t.Run()
	teardownFunction()
	os.Exit(retCode)
}
