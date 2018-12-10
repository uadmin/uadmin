package uadmin

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestPage404Handler is a unit testing function for page404Handler() function
func TestPage404Handler(t *testing.T) {
	r := httptest.NewRequest("GET", "http://0.0.0.0:5000/none-handled-url", nil)
	w := httptest.NewRecorder()

	page404Handler(w, r, nil)

	if w.Code != http.StatusNotFound {
		t.Errorf("Invalid code on requesting /. %d", w.Code)
	}

	title, _ := getHTMLTitle(w.Result().Body)
	title = strings.TrimSpace(title)
	if title != "uAdmin - 404" {
		t.Errorf("Invalid page returned. Expected 404, got (%s)", title)
	}
}
