package uadmin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// TestLoginHandler is a unit testing function for loginHandler() function
func TestLoginHandler(t *testing.T) {
	// Setup

	u1 := &User{
		Username: "user1",
		Password: "password1",
		Active:   false,
	}
	u1.Save()

	u2 := &User{
		Username:    "user2",
		Password:    "password2",
		Email:       "test@example.com",
		Active:      true,
		OTPRequired: true,
	}
	u2.Save()

	var w *httptest.ResponseRecorder

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
		nextURL   string
		postParam map[string][]string
		attr      []attrExample
	}{
		{
			httptest.NewRequest("POST", fmt.Sprintf("/"), nil),
			http.StatusSeeOther,
			"/",
			map[string][]string{
				"username": {"admin"},
				"password": {"admin"},
			},
			[]attrExample{},
		},
		{
			httptest.NewRequest("POST", fmt.Sprintf("/?next=/testmodelb/"), nil),
			http.StatusSeeOther,
			"/testmodelb/",
			map[string][]string{
				"username": {"admin"},
				"password": {"admin"},
			},
			[]attrExample{},
		},
		{
			httptest.NewRequest("POST", fmt.Sprintf("/"), nil),
			http.StatusOK,
			"/",
			map[string][]string{
				"username": {"admin"},
				"password": {"admin1"},
			},
			[]attrExample{},
		},
		{
			httptest.NewRequest("POST", fmt.Sprintf("/"), nil),
			http.StatusOK,
			"/",
			map[string][]string{
				"username": {"admin1"},
				"password": {"admin"},
			},
			[]attrExample{},
		},
		{
			httptest.NewRequest("POST", fmt.Sprintf("/"), nil),
			http.StatusOK,
			"/",
			map[string][]string{
				"username": {"user1"},
				"password": {"password1"},
			},
			[]attrExample{},
		},
		{
			httptest.NewRequest("POST", fmt.Sprintf("/"), nil),
			http.StatusOK,
			"/",
			map[string][]string{
				"username": {"user2"},
				"password": {"password2"},
			},
			[]attrExample{},
		},
		{
			httptest.NewRequest("POST", fmt.Sprintf("/"), nil),
			http.StatusSeeOther,
			"/",
			map[string][]string{
				"username": {"user2"},
				"password": {"password2"},
				"otp":      {u2.GetOTP()},
			},
			[]attrExample{},
		},
		{
			httptest.NewRequest("POST", fmt.Sprintf("/"), nil),
			http.StatusOK,
			"/",
			map[string][]string{
				"save":  {"Send Request"},
				"email": {"test@example.com"},
			},
			[]attrExample{},
		},
		{
			httptest.NewRequest("POST", fmt.Sprintf("/"), nil),
			http.StatusOK,
			"/",
			map[string][]string{
				"save":  {"Send Request"},
				"email": {"test1@example.com"},
			},
			[]attrExample{},
		},
	}

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

		loginHandler(w, e.r)

		if w.Code != e.code {
			t.Errorf("loginHandler returned wrong code. Expected: %d, got %d at (%d)", e.code, w.Code, i)
			continue
		}

		doc, err := parseHTML(w.Result().Body, t)
		if err != nil {
			t.Errorf("loginHandler returned invalid HTML content. %s at (%d)", err, i)
			continue
		}

		if w.Code == http.StatusSeeOther {
			if e.nextURL != w.Header().Get("Location") {
				t.Errorf("loginHandler returned invlid next url. Expected %s got %s at (%d)", e.nextURL, w.Header().Get("Location"), i)
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
					t.Errorf("loginHandler returned attrribue %s=%s for attr %s. Expected(%v) %#v, got (%#v) for %s(%d-%d)", tempAttr.selectorKey, tempAttr.selectorValue, tempAttr.checkKey, tempAttr.expected, tempAttr.checkValue, tempValue, tag, i, counter)
				} else {
					if index != -1 {
						e.attr[counter].path = path[index]
					}
				}
			}
		}
	}

	// Clean up
	Delete(u1)
	Delete(u2)
}
