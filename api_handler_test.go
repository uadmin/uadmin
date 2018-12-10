package uadmin

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestAPIHandler is a unit testing function for apiHandler() function
func TestAPIHandler(t *testing.T) {
	r := httptest.NewRequest("GET", "http://0.0.0.0:5000/api", nil)
	w := httptest.NewRecorder()

	apiHandler(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Invalid code on requesting /. %d", w.Code)
	}

	title, _ := getHTMLTitle(w.Result().Body)
	title = strings.TrimSpace(title)
	if title != "uAdmin - Login" {
		t.Errorf("Invalid page returned. Expected Login, got (%s)", title)
	}
}
