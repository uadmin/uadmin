package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/rn1hd/todo/models"
	"github.com/uadmin/uadmin"
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

	// Fetch Data from DB
	todo := []models.Todo{}
	uadmin.All(&todo)

	for i := range todo {
		// Accesses and fetches the record of the linking models in Todo
		uadmin.Preload(&todo[i])

		// Assigns the string of interface in each Todo fields
		c.TodoList = append(c.TodoList, map[string]interface{}{
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

	// Pass TodoList data object to the specified HTML path
	uadmin.HTMLContext(w, c, "views/todo.html")
}
