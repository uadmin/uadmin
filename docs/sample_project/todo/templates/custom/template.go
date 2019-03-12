package templates

import (
	"net/http"
	"strings"
)

// TemplateHandler !
func TemplateHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called /template
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/template")

	if strings.HasPrefix(r.URL.Path, "/todo_html") {
		TodoTemplateHandler(w, r)
		return
	}
}
