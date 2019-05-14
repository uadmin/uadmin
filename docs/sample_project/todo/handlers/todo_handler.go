package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
	"github.com/uadmin/uadmin/docs/sample_project/todo/models"
)

// TodoHandler !
func TodoHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called /todo_html
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/todo")

	// TodoList field inside the Context that will be used in Golang
	// HTML template
	type Context struct {
		TodoList []map[string]interface{}
	}

	// Assigns Context struct to the c variable
	c := Context{}

	// Initializes mapTodo as a map[string]interface{}{} where you can
	// create a dictionary that has a key and value from the database
	mapTodo := []map[string]interface{}{}

	// Fetch Data from DB
	todo := []models.Todo{}
	uadmin.All(&todo)

	for i := range todo {
		// Accesses and fetches the record of the linking models in Todo
		uadmin.Preload(&todo[i])

		// Assigns the string of interface in each Todo fields
		mapTodo = append(mapTodo, map[string]interface{}{
			"ID":          todo[i].ID,
			"Name":        todo[i].Name,
			"Description": todo[i].Description,
			"Category":    todo[i].Category.Name,
			"Friend":      todo[i].Friend.Name,
			"Item":        todo[i].Item.Name,
			"TargetDate":  todo[i].TargetDate,
			"Progress":    todo[i].Progress,
		})
	}

	// Assigns mapTodo to the TodoList inside the Context struct
	c.TodoList = mapTodo

	// Creates a new template and parses the template definitions
	// from todo.html
	t, _ := template.New("").ParseFiles("./views/todo.html")

	// Applies the template associated with t that has todo.html to
	// the specified object and writes the output to w variable that
	// is the http.ResponseWriter
	_ = t.ExecuteTemplate(w, "todo.html", c)
}
