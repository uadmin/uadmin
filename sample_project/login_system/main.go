package main

import (
	"net/http"

	"github.com/uadmin/uadmin"
	"github.com/uadmin/uadmin/sample_project/login_system/handlers"
)

func main() {
	// The listener is mapped to /admin/ path in the URL.
	uadmin.RootURL = "/admin/"

	// Sets the name of the website that shows on title and dashboard
	uadmin.SiteName = "Login System"

	// Login Handler
	http.HandleFunc("/login/", handlers.LoginHandler)

	// Run the server
	uadmin.StartServer()
}
