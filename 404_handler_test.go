package uadmin

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestPage404Handler is a unit testing function for page404Handler() function
func TestPage404Handler(t *testing.T) {
	r := httptest.NewRequest("GET", "http://0.0.0.0:5000/none-handled-url", nil)
	w := httptest.NewRecorder()

	pageErrorHandler(w, r, nil)

	if w.Code != http.StatusNotFound {
		t.Errorf("Invalid code on requesting /. %d", w.Code)
	}

	title, _, _ := getHTMLTag(w.Result().Body, "title")
	if title != "uAdmin - 404" {
		t.Errorf("Invalid page returned. Expected 404, got (%s)", title)
	}
}
