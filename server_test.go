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
	os.Remove("uadmin.db")
	os.Remove(".key")
	os.Remove(".salt")
	os.Remove(".uproj")
	os.Remove(".bindip")
}

func TestMain(t *testing.M) {
	setupFunction()
	retCode := t.Run()
	teardownFunction()
	os.Exit(retCode)
}
