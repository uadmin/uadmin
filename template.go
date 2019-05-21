package uadmin

import (
	"html/template"
	"net/http"
)

// HTMLContext creates a new template and applies a parsed template to the specified
// data object.
func HTMLContext(w http.ResponseWriter, data interface{}, path ...string) {
	tmpl := template.Must(template.ParseFiles(path...))
	tmpl.Execute(w, data)
}
