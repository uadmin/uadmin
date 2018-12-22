package uadmin

import (
	"os"
	"testing"
)

type TestModelA struct {
	Model
	Name string
}

type TestModelB struct {
	Model
	Name         string     `uadmin:"help:This is a test help message;search;list_exclude"`
	ItemCount    int        `uadmin:"max:5;min:1;format:%03d;required;read_only:true,edit"`
	Phone        string     `uadmin:"default_value:09;pattern:[0-9+]{7,15};pattern_msg:invalid phone number;encrypt"`
	Active       bool       `uadmin:"hidden;read_only"`
	OtherModel   TestModelA `uadmin:"categorical_filter;filter;read_only:new"`
	OtherModelID uint
	ModelAList   []TestModelA
	Parent       *TestModelB
	ParentID     uint
	Email        string  `uadmin:"email"`
	Greeting     string  `uadmin:"multilingual"`
	Image        string  `uadmin:"image;upload_to:/home/me/images/"`
	File         string  `uadmin:"file;upload_to:home/me/files"`
	Secret       string  `uadmin:"password"`
	Description  string  `uadmin:"html"`
	URL          string  `uadmin:"link"`
	Code         string  `uadmin:"code"`
	P1           int     `uadmin:"progress_bar"`
	P2           float64 `uadmin:"progress_bar"`
	P3           float64 `uadmin:"progress_bar:1.0"`
	P4           float64 `uadmin:"progress_bar:1.0:red"`
	P5           float64 `uadmin:"progress_bar:1.0:#f00"`
	P6           float64 `uadmin:"progress_bar:0.3:red,0.7:yellow,1.0:lime"`
	Price        float64 `uadmin:"money"`
	List         testList
}

func (TestModelB) Method__List() string {
	return "Value"
}

type testList int

func (testList) A() testList {
	return 1
}

func setupFunction() {
	Register(
		TestStruct1{},
		TestModelA{},
		TestModelB{},
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
