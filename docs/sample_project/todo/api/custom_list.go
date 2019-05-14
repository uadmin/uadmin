package api

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
	"github.com/uadmin/uadmin/docs/sample_project/todo/models"
)

// CustomListHandler !
func CustomListHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/custom_list")

	// Fetch Data from DB
	todo := []models.Todo{}

	// Assigns a map as a string of interface to store any types of values
	results := []map[string]interface{}{}

	// "id" - order the todo model by id
	// false - to sort in descending order
	// 0 - start at index 0
	// 5 - get five records
	// &todo - todo model to execute
	// "" - fetch the id of the model itself
	uadmin.AdminPage("id", false, 0, 5, &todo, "")

	// Loop to fetch the record of todo
	for i := range todo {
		// Accesses and fetches the record of the linking models in Todo
		uadmin.Preload(&todo[i])

		// Assigns the string of interface in each Todo fields
		results = append(results, map[string]interface{}{
			"ID":          todo[i].ID,
			"Name":        todo[i].Name,
			"Description": todo[i].Description,
			// This returns only the name of the Category model, not the
			// other fields
			"Category": todo[i].Category.Name,
			// This returns only the name of the Friend model, not the
			// other fields
			"Friend": todo[i].Friend.Name,
			// This returns only the name of the Item model, not the other
			// fields
			"Item":       todo[i].Item.Name,
			"TargetDate": todo[i].TargetDate,
			"Progress":   todo[i].Progress,
		})
	}

	// Prints the results in JSON format
	uadmin.ReturnJSON(w, r, results)
}
