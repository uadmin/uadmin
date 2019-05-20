package uadmin

import (
	"html/template"
	"io"
)

// HTMLContext creates a new template and applies a parsed template to the specified
// data object.
func HTMLContext(wr io.Writer, data interface{}, filenames ...string) {
	tmpl := template.Must(template.ParseFiles(filenames...))
	tmpl.Execute(wr, data)
}
