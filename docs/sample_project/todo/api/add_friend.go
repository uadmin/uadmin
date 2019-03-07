package api

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin/docs/sample_project/todo/models"
	"github.com/uadmin/uadmin"
)

// AddFriendHandler !
func AddFriendHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/add_friend")
	res := map[string]interface{}{}

	// Fetch data from Friend DB
	friend := models.Friend{}

	// Set the parameters of Name, Email, and Password
	friendName := r.FormValue("name")
	friendEmail := r.FormValue("email")
	friendPassword := r.FormValue("password")

	// Validate if the friendName variable is empty.
	if friendName == "" {
		res["status"] = "ERROR"
		res["err_msg"] = "Name is required."
		uadmin.ReturnJSON(w, r, res)
		return
	}

	// Store input into the Name, Email, and Password fields
	friend.Name = friendName
	friend.Email = friendEmail
	friend.Password = friendPassword

	// Store input in the Friend model
	uadmin.Save(&friend)

	res["status"] = "ok"
	uadmin.ReturnJSON(w, r, res)
}
