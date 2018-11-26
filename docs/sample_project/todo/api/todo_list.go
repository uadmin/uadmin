package api

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
	"github.com/uadmin/uadmin/docs/sample_project/todo/models"
)

// TodoListHandler !
func TodoListHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called /todo_list
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/todo_list")

	// Initializes res as a map[string]interface{}{} where you can put
	// anything inside it.
	res := map[string]interface{}{}

	// If r.URL.Path has no .json, it will display this error message in
	// JSON format.
	if r.URL.Path == "" || r.URL.Path[0] != '.' {
		res["status"] = "ERROR"
		res["err_msg"] = "No data type was specified"
		uadmin.ReturnJSON(w, r, res)
		return
	}

	// Initializes filterList as an array of string and valueList as an
	// array of interface
	filterList := []string{}
	valueList := []interface{}{}

	// Gets the ID of the todo model, append to the filterList and
	// valueList
	if r.URL.Query().Get("todo_id") != "" {
		filterList = append(filterList, "todo_id = ?")
		valueList = append(valueList, r.URL.Query().Get("todo_id"))
	}

	// Concatenates filterList by AND to store all the data in the filter
	// variable
	filter := strings.Join(filterList, " AND ")

	// Fetch Data from DB
	todo := []models.Todo{}
	uadmin.Filter(&todo, filter, valueList...)

	// Accesses and fetches data from another model
	for t := range todo {
		uadmin.Preload(&todo[t])
	}

	// Prints the todo in JSON format
	res["status"] = "ok"
	res["todo"] = todo
	uadmin.ReturnJSON(w, r, res)
}
