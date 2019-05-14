package handlers

import (
	"net/http"
	"strings"
)

// HTTPHandler !
func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called /http_handler
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/http_handler")

	if strings.HasPrefix(r.URL.Path, "/todo") {
		TodoHandler(w, r)
		return
	}
}
