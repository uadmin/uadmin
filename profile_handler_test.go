package uadmin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// TestProfileHandler is a unit testing function for profileHandler() function
func TestProfileHandler(t *testing.T) {
	// Setup
	var w *httptest.ResponseRecorder

	s1 := &Session{
		UserID: 1,
		Active: true,
	}
	s1.GenerateKey()
	s1.Save()

	Preload(s1)

	type attrExample struct {
		tag           string
		selectorKey   string
		selectorValue string
		checkKey      string
		checkValue    string
		parentIndex   int
		path          string
		expected      bool
	}

	examples := []struct {
		r         *http.Request
		code      int
		s         *Session
		nextURL   string
		postParam map[string][]string
		attr      []attrExample
	}{
		{
			httptest.NewRequest("GET", fmt.Sprintf("/"), nil),
			http.StatusOK,
			s1,
			"/",
			map[string][]string{},
			[]attrExample{},
		},
		{
			httptest.NewRequest("GET", fmt.Sprintf("/?otp_required=1"), nil),
			http.StatusOK,
			s1,
			"/",
			map[string][]string{},
			[]attrExample{},
		},
		{
			httptest.NewRequest("GET", fmt.Sprintf("/?otp_required=0"), nil),
			http.StatusOK,
			s1,
			"/",
			map[string][]string{},
			[]attrExample{},
		},
		{
			httptest.NewRequest("POST", fmt.Sprintf("/"), nil),
			http.StatusOK,
			s1,
			"/",
			map[string][]string{
				"save":      {""},
				"Username":  {"admin"},
				"FirstName": {"Updated System"},
				"LastName":  {"updated Admin"},
				"Email":     {"admin@example.com"},
				"Photo":     {""},
			},
			[]attrExample{},
		},
		{
			httptest.NewRequest("POST", fmt.Sprintf("/"), nil),
			http.StatusOK,
			s1,
			"/",
			map[string][]string{
				"save":            {"password"},
				"oldPassword":     {"wrong pass"},
				"newPassword":     {"new pass"},
				"confirmPassword": {"new pass"},
			},
			[]attrExample{},
		},
		{
			httptest.NewRequest("POST", fmt.Sprintf("/"), nil),
			http.StatusOK,
			s1,
			"/",
			map[string][]string{
				"save":            {"password"},
				"oldPassword":     {"admin"},
				"newPassword":     {"new pass"},
				"confirmPassword": {"pass"},
			},
			[]attrExample{},
		},
		{
			httptest.NewRequest("POST", fmt.Sprintf("/"), nil),
			http.StatusOK,
			s1,
			"/",
			map[string][]string{
				"save":            {"password"},
				"oldPassword":     {"admin"},
				"newPassword":     {"new pass"},
				"confirmPassword": {"new pass"},
			},
			[]attrExample{},
		},
	}

	c := &http.Cookie{}
	c.Name = "session"
	c.Value = s1.Key

	for i, e := range examples {
		w = httptest.NewRecorder()

		if e.r.Form == nil {
			e.r.Form = url.Values{}
		}
		if e.r.PostForm == nil {
			e.r.PostForm = url.Values{}
		}
		for k, v := range e.postParam {
			e.r.Form[k] = v
			e.r.PostForm[k] = v
		}
		e.r.AddCookie(c)
		profileHandler(w, e.r, e.s)

		if w.Code != e.code {
			t.Errorf("profileHandler returned wrong code. Expected: %d, got %d at (%d)", e.code, w.Code, i)
			continue
		}

		doc, err := parseHTML(w.Result().Body, t)
		if err != nil {
			t.Errorf("loginHandler returned invalid HTML content. %s at (%d)", err, i)
			continue
		}

		if w.Code == http.StatusSeeOther {
			if e.nextURL != w.Header().Get("Location") {
				t.Errorf("profileHandler returned invlid next url. Expected %s got %s at (%d)", e.nextURL, w.Header().Get("Location"), i)
			}
		}

		tagList := []string{}
		tagMap := map[string]bool{}
		for _, attr := range e.attr {
			if _, ok := tagMap[attr.tag]; !ok {
				tagMap[attr.tag] = true
				tagList = append(tagList, attr.tag)
			}
		}

		for _, tag := range tagList {
			// Parse HTML response
			path, content, attr := tagSearch(doc, tag, "", 0)
			_ = content

			// Verify input attribues
			for counter, tempAttr := range e.attr {
				if tempAttr.tag != tag {
					continue
				}
				parentPath := ""
				if tempAttr.parentIndex != -1 {
					parentPath = e.attr[tempAttr.parentIndex].path
				}
				index, tempValue := checkTagAttr(tempAttr.selectorKey, tempAttr.selectorValue, tempAttr.checkKey, tempAttr.checkValue, attr, path, parentPath)
				if !xOR(index == -1, tempAttr.expected) {
					t.Errorf("profileHandler returned attrribue %s=%s for attr %s. Expected(%v) %#v, got (%#v) for %s(%d-%d)", tempAttr.selectorKey, tempAttr.selectorValue, tempAttr.checkKey, tempAttr.expected, tempAttr.checkValue, tempValue, tag, i, counter)
				} else {
					if index != -1 {
						e.attr[counter].path = path[index]
					}
				}
			}
		}
	}
}
