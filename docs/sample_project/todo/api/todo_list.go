package api

import (
	"net/http"
	"strings"

	"github.com/rn1hd/todo/models"
	"github.com/uadmin/uadmin"
)

// TodoListHandler !
func TodoListHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called /todo_list
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/todo_list")

	// Fetches all object in the database
	todo := []models.Todo{}
	uadmin.All(&todo)

	// Accesses and fetches data from another model
	for t := range todo {
		uadmin.Preload(&todo[t])
	}

	// Prints the todo in JSON format
	uadmin.ReturnJSON(w, r, todo)
}
