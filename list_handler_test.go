package uadmin

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestListHandler is a unit testing function for listHandler() function
func TestListHandler(t *testing.T) {
	// Setup
	s1 := &Session{
		Active: true,
		UserID: 1,
	}
	s1.GenerateKey()
	s1.Save()
	Preload(s1)

	m1 := TestModelB{
		Name:         "Testing",
		ItemCount:    13,
		Phone:        "+18005551234",
		Active:       true,
		OtherModelID: 0,
		ModelAList:   []TestModelA{},
		ParentID:     0,
		Email:        "uadmin@example.com",
		Greeting:     "Hello uAdmin",
		Image:        "",
		File:         "",
		Secret:       "1234",
		Description:  "<p>This is the description of this fields</p>",
		URL:          "/",
		Code:         "Code good code in here",
		P1:           50,
		P2:           60.0,
		P3:           0.5,
		P4:           0.2,
		P5:           0.8,
		P6:           0.4,
		Price:        100.0,
		List:         testList(1),
	}
	Save(&m1)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/testmodelb/", nil)

	listHandler(w, r, s1)
	if w.Code != http.StatusOK {
		t.Errorf("listHandler returned wrong code. Expected: %d, got %d", http.StatusOK, w.Code)
	}

	// Clean up
	Delete(s1)
}
