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

	u1 := &User{
		Username: "u1",
		Password: "u1",
		Active:   true,
	}
	u1.Save()

	s2 := &Session{
		UserID: u1.ID,
		Active: true,
	}
	s2.GenerateKey()
	s2.Save()
	Preload(s2)

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

	var w *httptest.ResponseRecorder

	examples := []struct {
		r    *http.Request
		s    *Session
		code int
		f    func(*ModelSchema, *User) (string, []interface{})
	}{
		{httptest.NewRequest("GET", "/testmodelb/", nil), s1, 200, nil},
		{httptest.NewRequest("GET", "/testmodelb/", nil), s2, 404, nil},
		{httptest.NewRequest("GET", "/invalidmodel/", nil), s1, 404, nil},
		{httptest.NewRequest("GET", "/testmodelb/", nil), s1, 200, func(s *ModelSchema, u *User) (string, []interface{}) {
			return "", []interface{}{}
		}},
	}

	// Store the ListModifier
	f := Schema["testmodelb"].ListModifier

	// Run examples
	for _, e := range examples {
		schema := Schema["testmodelb"]
		schema.ListModifier = e.f
		Schema["testmodelb"] = schema

		w = httptest.NewRecorder()
		listHandler(w, e.r, e.s)
		if w.Code != e.code {
			t.Errorf("listHandler returned wrong code. Expected: %d, got %d", e.code, w.Code)
		}
	}

	// Reset the ListModifier func to nil
	schema := Schema["testmodelb"]
	schema.ListModifier = f
	Schema["testmodelb"] = schema

	// Clean up
	Delete(s1)
	Delete(s2)
	Delete(u1)
	Delete(m1)
}
