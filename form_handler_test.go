package uadmin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestFormHandler(t *testing.T) {
	var w *httptest.ResponseRecorder
	now := time.Now()
	s1 := &Session{
		Active: true,
		UserID: 1,
	}
	s1.GenerateKey()
	s1.Save()
	Preload(s1)

	m1 := TestStruct{
		Name: "test",
	}
	Save(&m1)

	m2 := TestStruct{
		Name:     "test",
		ParentID: m1.ID,
	}
	Save(&m2)

	m3 := TestStruct1{
		Name:  "testing the name",
		Value: 5,
	}
	Save(&m3)

	m4 := TestStruct2{
		Name:           "",
		Count:          1,
		Value:          2.54,
		Start:          time.Now(),
		End:            &now,
		Type:           TestType(1),
		OtherModelID:   m3.ID,
		AnotherModelID: m3.ID,
		Active:         true,
		Hidden:         "dd",
	}
	Save(&m4)

	m5 := TestModelA{
		Name: "Test",
	}
	Save(&m5)

	m6 := TestModelB{
		Name:         "Testing",
		ItemCount:    13,
		Phone:        "+18005551234",
		Active:       true,
		OtherModelID: m5.ID,
		ModelAList:   []TestModelA{m5},
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
	Save(&m6)

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

	// Test get form with existing record
	examples := []struct {
		r    *http.Request
		code int
		attr []attrExample
	}{
		{
			httptest.NewRequest("GET", fmt.Sprintf("/teststruct/%d", m1.ID), nil),
			http.StatusOK,
			[]attrExample{
				attrExample{"input", "name", "ID", "value", fmt.Sprint(m1.ID), -1, "", true},
				attrExample{"input", "name", "Name", "value", m1.Name, -1, "", true},
				attrExample{"select", "name", "ParentID", "name", "ParentID", -1, "", true},
				attrExample{"option", "value", "0", "selected", "", 2, "", true},
				attrExample{"option", "value", fmt.Sprint(m1.ID), "selected", "", 2, "", false},
				attrExample{"option", "value", fmt.Sprint(m2.ID), "selected", "", 2, "", false},
			},
		},
		{
			httptest.NewRequest("GET", fmt.Sprintf("/teststruct/%d", m2.ID), nil),
			http.StatusOK,
			[]attrExample{
				attrExample{"input", "name", "ID", "value", fmt.Sprint(m2.ID), -1, "", true},
				attrExample{"input", "name", "Name", "value", m2.Name, -1, "", true},
				attrExample{"select", "name", "ParentID", "name", "ParentID", -1, "", true},
				attrExample{"option", "value", "0", "selected", "", 2, "", false},
				attrExample{"option", "value", fmt.Sprint(m1.ID), "selected", "", 2, "", true},
				attrExample{"option", "value", fmt.Sprint(m2.ID), "selected", "", 2, "", false},
			},
		},
		{
			httptest.NewRequest("GET", fmt.Sprintf("/teststruct1/%d", m3.ID), nil),
			http.StatusOK,
			[]attrExample{
				attrExample{"input", "name", "ID", "value", fmt.Sprint(m3.ID), -1, "", true},
				attrExample{"input", "name", "Name", "value", m3.Name, -1, "", true},
				attrExample{"input", "name", "Value", "value", fmt.Sprint(m3.Value), -1, "", true},
			},
		},
		{
			httptest.NewRequest("GET", fmt.Sprintf("/teststruct2/%d", m4.ID), nil),
			http.StatusOK,
			[]attrExample{
				attrExample{"input", "name", "ID", "value", fmt.Sprint(m4.ID), -1, "", true},
				attrExample{"input", "name", "Name", "value", m4.Name, -1, "", true},
				attrExample{"input", "name", "Count", "value", fmt.Sprint(m4.Count), -1, "", true},
				attrExample{"input", "name", "Value", "value", fmt.Sprint(m4.Value), -1, "", true},
				attrExample{"input", "name", "Start", "value", m4.Start.Format("2006-01-02 15:04:05"), -1, "", true},
				attrExample{"input", "name", "End", "value", m4.End.Format("2006-01-02 15:04:05"), -1, "", true},
				attrExample{"select", "name", "Type", "name", "Type", -1, "", true},
				attrExample{"option", "value", "0", "selected", "", 6, "", false},
				attrExample{"option", "value", "1", "selected", "", 6, "", true},
				attrExample{"option", "value", "2", "selected", "", 6, "", false},
				attrExample{"select", "name", "OtherModelID", "name", "OtherModelID", -1, "", true},
				attrExample{"option", "value", "0", "selected", "", 10, "", false},
				attrExample{"option", "value", fmt.Sprint(m3.ID), "selected", "", 10, "", true},
				attrExample{"select", "name", "AnotherModelID", "name", "AnotherModelID", -1, "", true},
				attrExample{"option", "value", "0", "selected", "", 12, "", false},
				attrExample{"option", "value", fmt.Sprint(m3.ID), "selected", "", 12, "", true},
				attrExample{"input", "name", "Active", "checked", "", -1, "", true},
				attrExample{"input", "name", "Hidden", "value", m4.Hidden, -1, "", true},
			},
		},
		{
			httptest.NewRequest("GET", fmt.Sprintf("/testmodela/%d", m5.ID), nil),
			http.StatusOK,
			[]attrExample{
				attrExample{"input", "name", "ID", "value", fmt.Sprint(m5.ID), -1, "", true},
				attrExample{"input", "name", "Name", "value", m5.Name, -1, "", true},
			},
		},
		{
			httptest.NewRequest("GET", fmt.Sprintf("/testmodelb/%d", m6.ID), nil),
			http.StatusOK,
			[]attrExample{
				attrExample{"input", "name", "ID", "value", fmt.Sprint(m6.ID), -1, "", true},
				attrExample{"input", "name", "Name", "value", m6.Name, -1, "", true},
			},
		},
		{
			httptest.NewRequest("GET", fmt.Sprintf("/teststruct/%d", m1.ID+100), nil),
			http.StatusNotFound,
			[]attrExample{},
		},
		{
			httptest.NewRequest("GET", fmt.Sprintf("/badname/%d", m1.ID), nil),
			http.StatusNotFound,
			[]attrExample{},
		},
	}

	for i, e := range examples {
		w = httptest.NewRecorder()
		formHandler(w, e.r, s1)

		if w.Code != e.code {
			t.Errorf("formHandler returned wrong code. Expected: %d, got %d at (%d)", e.code, w.Code, i)
			continue
		}

		doc, err := parseHTML(w.Result().Body, t)
		if err != nil {
			t.Errorf("formHandler returned invalid HTML content. %s at (%d)", err, i)
			continue
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
					t.Errorf("formHandler returned attrribue %s=%s for attr %s. Expected(%v) %s, got (%s) for %s(%d-%d)", tempAttr.selectorKey, tempAttr.selectorValue, tempAttr.checkKey, tempAttr.expected, tempAttr.checkValue, tempValue, tag, i, counter)
				} else {
					if index != -1 {
						e.attr[counter].path = path[index]
					}
				}
			}
		}
	}

	// Clean up
	Delete(s1)
	Delete(m1)
	Delete(m2)
	Delete(m3)
	Delete(m4)
	Delete(m5)
	Delete(m6)
}

func checkTagAttr(selectorKey string, selectorValue string, checkKey string, checkValue string, attr []map[string]string, path []string, pathPrefix string) (int, string) {
	for i := range attr {
		if tempName, ok := attr[i][selectorKey]; ok && tempName == selectorValue {
			if tempValue, ok := attr[i][checkKey]; ok {
				if tempValue != checkValue {
					if selectorKey == checkKey {
						continue
					}
					return -1, tempValue
				}
				if pathPrefix != "" {
					if strings.HasPrefix(path[i], pathPrefix) {
						return i, tempValue
					}
				} else {
					return i, tempValue
				}
			}
		}
	}
	return -1, ""
}

func xOR(a, b bool) bool {
	return (a || b) && !(a && b)
}
