package api

import (
	"fmt"
	"net/http"
	"strings"
)

// API_HELP This part of code is the API HELP to be printed out in the body of the
// web page.
const APIHelp = `TODO API HELP
For more assistance please contact Integritynet:
support@integritynet.biz

- todo:
============
    # method     : todo_list
    # Parameters:
	# Return    : json object that returns the list of your todo activities
============
    # method     : custom_list
    # Parameters:
	# Return    : json object that returns the list your last 5 todo activities sorted in descending order
============
	# method     : add_friend
	# Parameters:  name (string), email (string), password (string)
	# Return    : inserts the information in the Friend model
`

// APIHandler !
func APIHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called /api
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")

	// If there is no subsequent method, it will call the API_HELP
	// variable to display the message.
	if r.URL.Path == "/" {
		fmt.Fprintf(w, API_HELP)
	}
	if strings.HasPrefix(r.URL.Path, "/todo_list") {
		TodoListHandler(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/custom_list") {
		CustomListHandler(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/add_friend") {
		AddFriendHandler(w, r)
		return
	}
}
