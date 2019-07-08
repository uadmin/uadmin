package uadmin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// TestAPIHandler is a unit testing function for apiHandler() function
func TestAPIHandler(t *testing.T) {
	// Test with no auth
	r := httptest.NewRequest("GET", "/api", nil)
	w := httptest.NewRecorder()

	apiHandler(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Invalid code on requesting /. %d", w.Code)
	}

	title, _, _ := getHTMLTag(w.Result().Body, "title")
	if title != "uAdmin - Login" {
		t.Errorf("Invalid page returned. Expected Login, got (%s)", title)
	}

	s1 := &Session{
		Active: true,
		UserID: 1,
	}
	s1.GenerateKey()
	s1.Save()

	tx := db.Begin()
	for i := 0; i < 10; i++ {
		rec := TestStruct1{
			Name:  fmt.Sprintf("Record%d", i+1),
			Value: i % 2,
		}
		tx.Create(&rec)
	}
	tx.Commit()

	// Test Search
	searchExamples := []struct {
		r     *http.Request
		count int
	}{
		{httptest.NewRequest("GET", "/api/search/?m=teststruct1&q=Record", nil), 10},
		{httptest.NewRequest("GET", "/api/search/?m=teststruct1&q=Record1", nil), 2},
		{httptest.NewRequest("GET", "/api/search/?m=teststruct1&q=Record2", nil), 1},
		{httptest.NewRequest("GET", "/api/search/?m=teststruct1&q=Records", nil), 0},
		{httptest.NewRequest("GET", "/api/search/?m=teststruct1&q=record", nil), 10},
		{httptest.NewRequest("GET", "/api/search/?m=teststruct1&q=record1", nil), 2},
		{httptest.NewRequest("GET", "/api/search/?m=teststruct1&q=record2", nil), 1},
		{httptest.NewRequest("GET", "/api/search/?m=teststruct1&q=records", nil), 0},
		{httptest.NewRequest("GET", "/api/search/?m=teststruct1&q=record&o=id", nil), 10},
		{httptest.NewRequest("GET", "/api/search/?m=teststruct1&q=record&value=0", nil), 5},
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

	// Test with wrong model name
	r = httptest.NewRequest("GET", "/api/search/?m=badname&q=records", nil)
	w = httptest.NewRecorder()

	r.AddCookie(&c)

	apiHandler(w, r)

	if w.Code != http.StatusNotFound {
		t.Errorf("Invalid code on requesting with wrong model name. %d", w.Code)
	}

	// Test List Modifier
	schema := Schema["teststruct1"]
	schema.ListModifier = func(s *ModelSchema, u *User) (string, []interface{}) {
		return "value = ?", []interface{}{1}
	}
	Schema["teststruct1"] = schema

	searchExamples = []struct {
		r     *http.Request
		count int
	}{
		{httptest.NewRequest("GET", "/api/search/?m=teststruct1&q=Record", nil), 5},
	}

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

	schema.ListModifier = nil
	Schema["teststruct1"] = schema

	// Test upload image
	r, err := newfileUploadRequest("/api/upload_image/", map[string]string{}, "file", "./static/uadmin/logo.png")
	if err != nil {
		t.Errorf("newfileUploadRequest unable to create multipart request")
		return
	}
	r.AddCookie(&c)
	w = httptest.NewRecorder()

	apiHandler(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Invalid code on api/image_upload. %d", w.Code)
	}
	buf, _ := ioutil.ReadAll(w.Body)
	res := map[string]string{}

	err = json.Unmarshal(buf, &res)
	if err != nil {
		t.Errorf("apiHandler returned invalid JSON format during image_upload. %s", string(buf))
		return
	}
	if _, ok := res["location"]; !ok {
		t.Errorf("apiHandler didn't return 'location' for image_upload")
		return
	}
	if _, err = os.Stat("." + res["location"]); os.IsNotExist(err) {
		t.Errorf("apiHandler didn't create image file for image_upload %s", "."+res["location"])
		return
	}
}

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	if err != nil {
		return nil, err
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}
