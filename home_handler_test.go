package uadmin

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHomeHandler is a unit testing function for homeHandler() function
func TestHomeHandler(t *testing.T) {
	// Setup
	s1 := &Session{
		Active: true,
		UserID: 1,
	}
	s1.GenerateKey()
	s1.Save()
	Preload(s1)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	homeHandler(w, r, s1)
	if w.Code != http.StatusOK {
		t.Errorf("homeHandler returned wrong code. Expected: %d, got %d", http.StatusOK, w.Code)
	}

	// Clean up
	Delete(s1)
}
