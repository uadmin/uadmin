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
			"ID":   todo[i].ID,
			"Name": todo[i].Name,
			// In fact that description has an html type tag in uAdmin,
			// we have to convert this field from text to HTML so that
			// the HTML tags from models will be applied to HTML file.
			"Description": template.HTML(todo[i].Description),
			"Category":    todo[i].Category.Name,
			"Friend":      todo[i].Friend.Name,
			"Item":        todo[i].Item.Name,
			"TargetDate":  todo[i].TargetDate,
			"Progress":    todo[i].Progress,
		})
	}

	// Assigns mapTodo to the TodoList inside the Context struct
	c.TodoList = mapTodo

	// Reads the HTML file
	tmpl := template.Must(template.ParseFiles("views/todo.html"))

	// Pass back-end TodoList data to the HTML file that we read
	tmpl.Execute(w, c)
}
