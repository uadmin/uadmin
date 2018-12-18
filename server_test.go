package uadmin

import (
	"os"
	"testing"
)

func setupFunction() {
	Register()
	Port = 5000
	go StartServer()
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
}

func TestMain(t *testing.M) {
	teardownFunction()
	setupFunction()
	retCode := t.Run()
	teardownFunction()
	os.Exit(retCode)
}
