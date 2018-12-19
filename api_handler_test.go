package uadmin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestAPIHandler is a unit testing function for apiHandler() function
func TestAPIHandler(t *testing.T) {
	r := httptest.NewRequest("GET", "/api", nil)
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

	s1 := &Session{
		Active: true,
		UserID: 1,
	}
	s1.GenerateKey()
	s1.Save()

	for i := 0; i < 10; i++ {
		rec := TestStruct1{
			Name: fmt.Sprintf("Record%d", i+1),
		}
		Save(&rec)
	}

	// Test Search
	searchExamples := []struct {
		r     *http.Request
		count int
	}{
		{httptest.NewRequest("GET", "/api?method=searchTable&model=teststruct1&q=Record", nil), 10},
		{httptest.NewRequest("GET", "/api?method=searchTable&model=teststruct1&q=Record1", nil), 2},
		{httptest.NewRequest("GET", "/api?method=searchTable&model=teststruct1&q=Record2", nil), 1},
		{httptest.NewRequest("GET", "/api?method=searchTable&model=teststruct1&q=Records", nil), 0},
		{httptest.NewRequest("GET", "/api?method=searchTable&model=teststruct1&q=record", nil), 10},
		{httptest.NewRequest("GET", "/api?method=searchTable&model=teststruct1&q=record1", nil), 2},
		{httptest.NewRequest("GET", "/api?method=searchTable&model=teststruct1&q=record2", nil), 1},
		{httptest.NewRequest("GET", "/api?method=searchTable&model=teststruct1&q=records", nil), 0},
	}

	c := http.Cookie{}
	c.Name = "session"
	c.Value = s1.Key

	for i, e := range searchExamples {
		w = httptest.NewRecorder()
		e.r.AddCookie(&c)

		apiHandler(w, e.r)

		if w.Code != http.StatusOK {
			t.Errorf("Invalid code on requesting /api/. %d for request %d", w.Code, i)
			continue
		}
		buf, _ := ioutil.ReadAll(w.Body)
		res := map[string]interface{}{}

		err := json.Unmarshal(buf, &res)
		if err != nil {
			t.Errorf("apiHandler returned invalid JSON format during search. %s for request %d", string(buf), i)
			continue
		}
		if _, ok := res["list"]; !ok {
			t.Errorf("apiHandler didn't return 'list' for request %d", i)
			continue
		}
		list, ok := res["list"].([]interface{})
		if !ok {
			t.Errorf("apiHandler 'list' is not a list of values for request %d", i)
			continue
		}
		if len(list) != e.count {
			t.Errorf("apiHandler returned the wrong number of values. Got %d, expected %d for request %d", len(list), e.count, i)
			continue
		}
	}

}
