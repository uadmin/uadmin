package uadmin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// TestProcessDelete is a unit testing function for processDelete() function
func TestProcessDelete(t *testing.T) {
	// Setup
	w := httptest.NewRecorder()

	idList := []string{}

	for i := 0; i < 10; i++ {
		record := TestStruct{
			Name: fmt.Sprintf("test%d", i),
		}
		Save(&record)
		idList = append(idList, fmt.Sprint(record.ID))
	}

	s := &Session{
		UserID: 1,
		Active: true,
	}
	s.GenerateKey()
	s.Save()

	// Tests
	examples := []struct {
		r     *http.Request
		count int
		data  map[string]string
	}{
		{
			httptest.NewRequest("POST", "/", nil), 0,
			map[string]string{},
		},
		{
			httptest.NewRequest("POST", "/", nil), 0,
			map[string]string{"listID": strings.Join(idList, ",")},
		},
		{
			httptest.NewRequest("POST", "/", nil), 0,
			map[string]string{"x-csrf-token": s.Key},
		},
		{
			httptest.NewRequest("POST", "/", nil), len(idList),
			map[string]string{"listID": strings.Join(idList, ","), "x-csrf-token": s.Key},
		},
	}

	user := User{}
	Get(&user, "id = ?", 1)

	for i, e := range examples {
		// Setup post values
		e.r.PostForm = url.Values{}
		for k, v := range e.data {
			e.r.PostForm.Set(k, v)
			e.r.AddCookie(&http.Cookie{Name: "session", Value: s.Key})
		}
		countBefore := Count(&TestStruct{}, "")
		processDelete("teststruct", w, e.r, nil, &user)
		countAfter := Count(TestStruct{}, "")
		if (countBefore - countAfter) != e.count {
			t.Errorf("Invalid number of deleted records by processDelete in example(%d). Expected %d, Got %d", i, e.count, (countBefore - countAfter))
		}
	}

	// Clean up
	DeleteList(TestStruct{}, "")
}
