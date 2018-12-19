package uadmin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFormHandler(t *testing.T) {
	s1 := &Session{
		Active: true,
		UserID: 1,
	}
	s1.GenerateKey()
	s1.Save()
	Preload(s1)

	m := TestStruct{
		Name: "test",
	}
	Save(&m)

	// Test get form
	r := httptest.NewRequest("GET", fmt.Sprintf("/teststruct/%d", m.ID), nil)
	w := httptest.NewRecorder()

	formHandler(w, r, s1)

	if w.Code != http.StatusOK {
		t.Errorf("formHandler returned wrong code. Expected: %d, got %d", http.StatusOK, w.Code)
	}

	// Test get invalid ID
	r = httptest.NewRequest("GET", fmt.Sprintf("/teststruct/%d", m.ID+10), nil)
	w = httptest.NewRecorder()

	formHandler(w, r, s1)

	if w.Code != http.StatusNotFound {
		t.Errorf("formHandler returned wrong code. Expected: %d, got %d", http.StatusNotFound, w.Code)
	}

	// Test get invalid model name
	r = httptest.NewRequest("GET", fmt.Sprintf("/teststructs/%d", m.ID), nil)
	w = httptest.NewRecorder()

	formHandler(w, r, s1)

	if w.Code != http.StatusNotFound {
		t.Errorf("formHandler returned wrong code. Expected: %d, got %d", http.StatusNotFound, w.Code)
	}

	// Test Save Form
	r = httptest.NewRequest("POST", fmt.Sprintf("/teststruct/%d", m.ID), nil)
	w = httptest.NewRecorder()

	formHandler(w, r, s1)

	if w.Code != http.StatusSeeOther {
		t.Errorf("formHandler returned wrong code. Expected: %d, got %d", http.StatusSeeOther, w.Code)
	}
}
