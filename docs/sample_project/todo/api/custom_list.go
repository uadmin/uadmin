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
	res := map[string]interface{}{}
	if r.URL.Path == "" || r.URL.Path[0] != '.' {
		res["status"] = "ERROR"
		res["err_msg"] = "No data type was specified"
		uadmin.ReturnJSON(w, r, res)
		return
	}
	filterList := []string{}
	valueList := []interface{}{}
	if r.URL.Query().Get("todo_id") != "" {
		filterList = append(filterList, "todo_id = ?")
		valueList = append(valueList, r.URL.Query().Get("todo_id"))
	}
	filter := strings.Join(filterList, " AND ")

	// Fetch Data from DB
	todo := []models.Todo{}

	// Assigns a map as a string of interface to store any types of values
	results := []map[string]interface{}{}

	// Fetches the ID of todo in the first parameter, second parameter as
	// false to sort in descending order, offset to 0 as a starting index
	// point in the third parameter, set the limit value to 5 to return
	// five data in the fourth parameter, calls the model in the fifth
	// parameter, query interface is filter in the sixth parameter, and
	// valueList is the argument called that can be used in the execution
	// process as the last parameter.
	uadmin.AdminPage("id", false, 0, 5, &todo, filter, valueList)

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
	res["status"] = "ok"
	res["todo"] = results
	uadmin.ReturnJSON(w, r, res)
}
