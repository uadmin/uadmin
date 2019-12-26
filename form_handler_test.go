package uadmin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

// TestFormHandler is a unit testing function for formHandler() function
func TestFormHandler(t *testing.T) {
	// Setup
	var w *httptest.ResponseRecorder
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, now.Location())
	tomorrow := now.AddDate(0, 0, 1)
	s1 := &Session{
		Active: true,
		UserID: 1,
	}
	s1.GenerateKey()
	s1.Save()
	Preload(s1)

	g1 := UserGroup{
		GroupName: "g1",
	}
	Save(&g1)

	testModelAdm := DashboardMenu{}
	Get(&testModelAdm, "url = ?", "testmodela")

	testModelBdm := DashboardMenu{}
	Get(&testModelBdm, "url = ?", "testmodelb")

	testApprovalModeldm := DashboardMenu{}
	Get(&testApprovalModeldm, "url = ?", "testapproval")

	gp1 := GroupPermission{
		DashboardMenuID: testModelAdm.ID,
		UserGroupID:     g1.ID,
		Read:            true,
		Edit:            true,
		Add:             true,
	}
	Save(&gp1)

	gp2 := GroupPermission{
		DashboardMenuID: testModelBdm.ID,
		UserGroupID:     g1.ID,
		Read:            true,
		Edit:            false,
		Add:             false,
	}
	Save(&gp2)

	gp3 := GroupPermission{
		DashboardMenuID: testApprovalModeldm.ID,
		UserGroupID:     g1.ID,
		Read:            true,
		Edit:            true,
		Add:             true,
	}
	Save(&gp3)

	u1 := &User{
		Username:    "u1",
		FirstName:   "User 1",
		Active:      true,
		Password:    "password",
		UserGroupID: g1.ID,
	}
	u1.Save()

	s2 := &Session{
		Active: true,
		UserID: u1.ID,
	}
	s2.GenerateKey()
	s2.Save()
	Preload(s2)

	u2 := &User{
		Username:    "u2",
		FirstName:   "User 2",
		Active:      true,
		Password:    "password",
		UserGroupID: g1.ID,
	}
	u2.Save()

	up1 := UserPermission{
		DashboardMenuID: testModelAdm.ID,
		UserID:          u2.ID,
		Read:            true,
		Edit:            false,
		Add:             false,
	}
	Save(&up1)

	s3 := &Session{
		Active: true,
		UserID: u2.ID,
	}
	s3.GenerateKey()
	s3.Save()
	Preload(s3)

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

	m7 := TestApproval{
		Name:        "Test",
		Start:       now,
		End:         &now,
		Count:       1,
		Price:       1.0,
		List:        testList(0),
		TestModelID: m5.ID,
		Active:      true,
	}
	Save(&m7)

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

	fileR1, _ := newfileUploadRequest(fmt.Sprintf("/testmodelb/%d", m6.ID), map[string]string{}, "File", "./static/uadmin/favicon.ico")
	fileR2, _ := newfileUploadRequest(fmt.Sprintf("/testmodelb/%d", m6.ID), map[string]string{}, "Image", "./media/user/image_raw.png")
	fileR3, _ := newfileUploadRequest(fmt.Sprintf("/testmodelb/%d", m6.ID), map[string]string{}, "Image", "./media/user/image_raw.jpg")
	fileR4, _ := newfileUploadRequest(fmt.Sprintf("/testmodelb/%d", m6.ID), map[string]string{}, "Image", "./media/user/image_raw.gif")

	// Test get form with existing record
	examples := []struct {
		r         *http.Request
		code      int
		s         *Session
		postParam map[string][]string
		attr      []attrExample
		approvals []Approval
	}{
		//0
		{
			httptest.NewRequest("GET", fmt.Sprintf("/teststruct/%d", m1.ID), nil),
			http.StatusOK,
			s1,
			map[string][]string{},
			[]attrExample{
				{"input", "name", "ID", "value", fmt.Sprint(m1.ID), -1, "", true},
				{"input", "name", "Name", "value", m1.Name, -1, "", true},
				{"select", "name", "ParentID", "name", "ParentID", -1, "", true},
				{"option", "value", "", "selected", "", 2, "", true},
				{"option", "value", fmt.Sprint(m1.ID), "selected", "", 2, "", false},
				{"option", "value", fmt.Sprint(m2.ID), "selected", "", 2, "", false},
			},
			[]Approval{},
		},
		//1
		{
			httptest.NewRequest("GET", fmt.Sprintf("/teststruct/%d", m2.ID), nil),
			http.StatusOK,
			s1,
			map[string][]string{},
			[]attrExample{
				{"input", "name", "ID", "value", fmt.Sprint(m2.ID), -1, "", true},
				{"input", "name", "Name", "value", m2.Name, -1, "", true},
				{"select", "name", "ParentID", "name", "ParentID", -1, "", true},
				{"option", "value", "", "selected", "", 2, "", false},
				{"option", "value", fmt.Sprint(m1.ID), "selected", "", 2, "", true},
				{"option", "value", fmt.Sprint(m2.ID), "selected", "", 2, "", false},
			},
			[]Approval{},
		},
		//2
		{
			httptest.NewRequest("GET", fmt.Sprintf("/teststruct1/%d", m3.ID), nil),
			http.StatusOK,
			s1,
			map[string][]string{},
			[]attrExample{
				{"input", "name", "ID", "value", fmt.Sprint(m3.ID), -1, "", true},
				{"input", "name", "Name", "value", m3.Name, -1, "", true},
				{"input", "name", "Value", "value", fmt.Sprint(m3.Value), -1, "", true},
			},
			[]Approval{},
		},
		//3
		{
			httptest.NewRequest("GET", fmt.Sprintf("/teststruct2/%d", m4.ID), nil),
			http.StatusOK,
			s1,
			map[string][]string{},
			[]attrExample{
				{"input", "name", "ID", "value", fmt.Sprint(m4.ID), -1, "", true},
				{"input", "name", "Name", "value", m4.Name, -1, "", true},
				{"input", "name", "Count", "value", fmt.Sprint(m4.Count), -1, "", true},
				{"input", "name", "Value", "value", fmt.Sprint(m4.Value), -1, "", true},
				{"input", "name", "Start", "value", m4.Start.Format("2006-01-02 15:04:05"), -1, "", true},
				{"input", "name", "End", "value", m4.End.Format("2006-01-02 15:04:05"), -1, "", true},
				{"select", "name", "Type", "name", "Type", -1, "", true},
				{"option", "value", "", "selected", "", 6, "", false},
				{"option", "value", "1", "selected", "", 6, "", true},
				{"option", "value", "2", "selected", "", 6, "", false},
				{"select", "name", "OtherModelID", "name", "OtherModelID", -1, "", true},
				{"option", "value", "0", "selected", "", 10, "", false},
				{"option", "value", fmt.Sprint(m3.ID), "selected", "", 10, "", true},
				{"select", "name", "AnotherModelID", "name", "AnotherModelID", -1, "", true},
				{"option", "value", "0", "selected", "", 12, "", false},
				{"option", "value", fmt.Sprint(m3.ID), "selected", "", 12, "", true},
				{"input", "name", "Active", "checked", "", -1, "", true},
				{"input", "name", "Hidden", "value", m4.Hidden, -1, "", true},
			},
			[]Approval{},
		},
		//4
		{
			httptest.NewRequest("GET", fmt.Sprintf("/testmodela/%d", m5.ID), nil),
			http.StatusOK,
			s1,
			map[string][]string{},
			[]attrExample{
				{"input", "name", "ID", "value", fmt.Sprint(m5.ID), -1, "", true},
				{"input", "name", "Name", "value", m5.Name, -1, "", true},
				{"button", "name", "save", "value", "", -1, "", true},
				{"button", "name", "save", "value", "continue", -1, "", true},
				{"button", "name", "save", "value", "another", -1, "", true},
			},
			[]Approval{},
		},
		//5
		{
			httptest.NewRequest("GET", fmt.Sprintf("/testmodelb/%d", m6.ID), nil),
			http.StatusOK,
			s1,
			map[string][]string{},
			[]attrExample{
				{"input", "name", "ID", "value", fmt.Sprint(m6.ID), -1, "", true},
				{"input", "name", "Name", "value", m6.Name, -1, "", true},
				{"input", "name", "ItemCount", "value", fmt.Sprintf("%03d", m6.ItemCount), -1, "", true},
				{"input", "name", "ItemCount", "required", "", -1, "", true},
				{"input", "name", "ItemCount", "readonly", "", -1, "", true},
				{"input", "name", "Phone", "value", m6.Phone, -1, "", true},
				{"input", "name", "Active", "name", "Active", -1, "", false},
				{"select", "name", "OtherModelID", "name", "OtherModelID", -1, "", true},
				{"select", "name", "OtherModelID", "readonly", "", -1, "", false},
				{"option", "value", "", "selected", "", 7, "", false},
				{"option", "value", fmt.Sprint(m5.ID), "selected", "", 7, "", true},
				{"select", "name", "ModelAList", "name", "ModelAList", -1, "", true},
				{"select", "name", "ModelAList", "multiple", "", -1, "", true},
				{"option", "value", fmt.Sprint(m5.ID), "selected", "", 10, "", true},
				{"select", "name", "ParentID", "name", "ParentID", -1, "", true},
				{"option", "value", "", "selected", "", 14, "", true},
				{"option", "value", fmt.Sprint(m6.ID), "selected", "", 14, "", false},
				{"input", "name", "Email", "value", m6.Email, -1, "", true},
				{"input", "name", "Email", "type", "email", -1, "", true},
				{"input", "name", "en-Greeting", "value", m6.Greeting, -1, "", true},
				{"input", "name", "File", "name", "File", -1, "", true},
				{"input", "name", "Image", "name", "Image", -1, "", true},
				{"input", "name", "Secret", "value", m6.Secret, -1, "", true},
				// TODO: Description
				// TODO: Link
				// TODO: Code
				{"input", "name", "P1", "value", fmt.Sprint(m6.P1), -1, "", true},
				{"input", "name", "P2", "value", fmt.Sprint(m6.P2), -1, "", true},
				{"input", "name", "P3", "value", fmt.Sprint(m6.P3), -1, "", true},
				{"input", "name", "P4", "value", fmt.Sprint(m6.P4), -1, "", true},
				{"input", "name", "P5", "value", fmt.Sprint(m6.P5), -1, "", true},
				{"input", "name", "P6", "value", fmt.Sprint(m6.P6), -1, "", true},
				{"input", "name", "Price", "value", fmt.Sprint(m6.Price), -1, "", true},
				{"button", "name", "save", "value", "", -1, "", true},
				{"button", "name", "save", "value", "continue", -1, "", true},
				{"button", "name", "save", "value", "another", -1, "", true},
			},
			[]Approval{},
		},
		//6
		{
			httptest.NewRequest("GET", fmt.Sprintf("/testmodelb/%d", m6.ID), nil),
			http.StatusOK,
			s2,
			map[string][]string{},
			[]attrExample{
				{"input", "name", "ID", "value", fmt.Sprint(m6.ID), -1, "", true},
				{"input", "name", "Name", "value", m6.Name, -1, "", true},
				{"input", "name", "ItemCount", "value", fmt.Sprintf("%03d", m6.ItemCount), -1, "", true},
				{"input", "name", "ItemCount", "required", "", -1, "", true},
				{"input", "name", "ItemCount", "readonly", "", -1, "", true},
				{"input", "name", "Phone", "value", m6.Phone, -1, "", true},
				{"input", "name", "Active", "name", "Active", -1, "", false},
				{"select", "name", "OtherModelID", "name", "OtherModelID", -1, "", true},
				{"select", "name", "OtherModelID", "readonly", "", -1, "", false},
				{"option", "value", "", "selected", "", 7, "", false},
				{"option", "value", fmt.Sprint(m5.ID), "selected", "", 7, "", true},
				{"select", "name", "ModelAList", "name", "ModelAList", -1, "", true},
				{"select", "name", "ModelAList", "multiple", "", -1, "", true},
				{"option", "value", fmt.Sprint(m5.ID), "selected", "", 10, "", true},
				{"select", "name", "ParentID", "name", "ParentID", -1, "", true},
				{"option", "value", "", "selected", "", 14, "", true},
				{"option", "value", fmt.Sprint(m6.ID), "selected", "", 14, "", false},
				{"input", "name", "Email", "value", m6.Email, -1, "", true},
				{"input", "name", "Email", "type", "email", -1, "", true},
				{"input", "name", "en-Greeting", "value", m6.Greeting, -1, "", true},
				{"input", "name", "File", "name", "File", -1, "", true},
				{"input", "name", "Image", "name", "Image", -1, "", true},
				{"input", "name", "Secret", "value", m6.Secret, -1, "", true},
				// TODO: Description
				// TODO: Link
				// TODO: Code
				{"input", "name", "P1", "value", fmt.Sprint(m6.P1), -1, "", true},
				{"input", "name", "P2", "value", fmt.Sprint(m6.P2), -1, "", true},
				{"input", "name", "P3", "value", fmt.Sprint(m6.P3), -1, "", true},
				{"input", "name", "P4", "value", fmt.Sprint(m6.P4), -1, "", true},
				{"input", "name", "P5", "value", fmt.Sprint(m6.P5), -1, "", true},
				{"input", "name", "P6", "value", fmt.Sprint(m6.P6), -1, "", true},
				{"input", "name", "Price", "value", fmt.Sprint(m6.Price), -1, "", true},
				{"button", "name", "save", "value", "", -1, "", false},
				{"button", "name", "save", "value", "continue", -1, "", false},
				{"button", "name", "save", "value", "another", -1, "", false},
			},
			[]Approval{},
		},
		//7
		{
			httptest.NewRequest("GET", fmt.Sprintf("/testmodela/%d", m5.ID), nil),
			http.StatusOK,
			s2,
			map[string][]string{},
			[]attrExample{
				{"input", "name", "ID", "value", fmt.Sprint(m5.ID), -1, "", true},
				{"input", "name", "Name", "value", m5.Name, -1, "", true},
				{"button", "name", "save", "value", "", -1, "", true},
				{"button", "name", "save", "value", "continue", -1, "", true},
				{"button", "name", "save", "value", "another", -1, "", true},
			},
			[]Approval{},
		},
		// 8
		{
			httptest.NewRequest("GET", fmt.Sprintf("/testmodela/%d", m5.ID), nil),
			http.StatusOK,
			s3,
			map[string][]string{},
			[]attrExample{
				{"input", "name", "ID", "value", fmt.Sprint(m5.ID), -1, "", true},
				{"input", "name", "Name", "value", m5.Name, -1, "", true},
				{"button", "name", "save", "value", "", -1, "", false},
				{"button", "name", "save", "value", "continue", -1, "", false},
				{"button", "name", "save", "value", "another", -1, "", false},
			},
			[]Approval{},
		},
		// 9
		{
			httptest.NewRequest("GET", fmt.Sprintf("/testmodela/new"), nil),
			http.StatusOK,
			s2,
			map[string][]string{},
			[]attrExample{
				{"input", "name", "ID", "value", "0", -1, "", true},
				{"input", "name", "Name", "value", "", -1, "", true},
				{"button", "name", "save", "value", "", -1, "", true},
				{"button", "name", "save", "value", "continue", -1, "", true},
				{"button", "name", "save", "value", "another", -1, "", false},
			},
			[]Approval{},
		},
		// 10
		{
			httptest.NewRequest("GET", fmt.Sprintf("/testmodela/new"), nil),
			http.StatusOK,
			s3,
			map[string][]string{},
			[]attrExample{
				{"input", "name", "ID", "value", "0", -1, "", true},
				{"input", "name", "Name", "value", "", -1, "", true},
				{"button", "name", "save", "value", "", -1, "", false},
				{"button", "name", "save", "value", "continue", -1, "", false},
				{"button", "name", "save", "value", "another", -1, "", false},
			},
			[]Approval{},
		},
		// 11
		{
			httptest.NewRequest("POST", fmt.Sprintf("/testmodelb/%d", m6.ID), nil),
			http.StatusSeeOther,
			s1,
			map[string][]string{
				"ID":   {fmt.Sprint(m6.ID)},
				"Name": {"Updated Name"},
				//"ItemCount": "34",
				"Phone":        {"+188854321"},
				"Active":       {"on"},
				"OtherModelID": {fmt.Sprint(m5.ID)},
				"ModelAList":   {fmt.Sprint(m5.ID)},
				"ParentID":     {"0"},
				"Email":        {"updated@example.com"},
				"en-Greeting":  {"Hello Updated Greeting"},
				"File":         {""},
				"Image":        {""},
				"Secret":       {"Updated Secret"},
				"P1":           {"2"},
				"P2":           {"0.2"},
				"P3":           {"0.3"},
				"P4":           {"0.4"},
				"P5":           {"0.5"},
				"P6":           {"0.6"},
				"Price":        {"100.01"},
				"save":         {"continue"},
			},
			[]attrExample{},
			[]Approval{},
		},
		// 12
		{
			httptest.NewRequest("GET", fmt.Sprintf("/testmodelb/%d", m6.ID), nil),
			http.StatusOK,
			s1,
			map[string][]string{},
			[]attrExample{
				{"input", "name", "ID", "value", fmt.Sprint(m6.ID), -1, "", true},
				{"input", "name", "Name", "value", "Updated Name", -1, "", true},
				{"input", "name", "ItemCount", "value", fmt.Sprintf("%03d", m6.ItemCount), -1, "", true},
				{"input", "name", "ItemCount", "required", "", -1, "", true},
				{"input", "name", "ItemCount", "readonly", "", -1, "", true},
				{"input", "name", "Phone", "value", "+188854321", -1, "", true},
				{"input", "name", "Active", "name", "Active", -1, "", false},
				{"select", "name", "OtherModelID", "name", "OtherModelID", -1, "", true},
				{"select", "name", "OtherModelID", "readonly", "", -1, "", false},
				{"option", "value", "", "selected", "", 7, "", false},
				{"option", "value", fmt.Sprint(m5.ID), "selected", "", 7, "", true},
				{"select", "name", "ModelAList", "name", "ModelAList", -1, "", true},
				{"select", "name", "ModelAList", "multiple", "", -1, "", true},
				{"option", "value", fmt.Sprint(m5.ID), "selected", "", 10, "", true},
				{"select", "name", "ParentID", "name", "ParentID", -1, "", true},
				{"option", "value", "", "selected", "", 14, "", true},
				{"option", "value", fmt.Sprint(m6.ID), "selected", "", 14, "", false},
				{"input", "name", "Email", "value", "updated@example.com", -1, "", true},
				{"input", "name", "Email", "type", "email", -1, "", true},
				{"input", "name", "en-Greeting", "value", "Hello Updated Greeting", -1, "", true},
				{"input", "name", "File", "name", "File", -1, "", true},
				{"input", "name", "Image", "name", "Image", -1, "", true},
				{"input", "name", "Secret", "value", "Updated Secret", -1, "", true},
				// TODO: Description
				// TODO: Link
				// TODO: Code
				{"input", "name", "P1", "value", "2", -1, "", true},
				{"input", "name", "P2", "value", "0.2", -1, "", true},
				{"input", "name", "P3", "value", "0.3", -1, "", true},
				{"input", "name", "P4", "value", "0.4", -1, "", true},
				{"input", "name", "P5", "value", "0.5", -1, "", true},
				{"input", "name", "P6", "value", "0.6", -1, "", true},
				{"input", "name", "Price", "value", fmt.Sprint("100.01"), -1, "", true},
				{"button", "name", "save", "value", "", -1, "", true},
				{"button", "name", "save", "value", "continue", -1, "", true},
				{"button", "name", "save", "value", "another", -1, "", true},
			},
			[]Approval{},
		},
		//13
		{
			fileR1,
			http.StatusSeeOther,
			s1,
			map[string][]string{
				"ID":   {fmt.Sprint(m6.ID)},
				"Name": {"Updated Name"},
				//"ItemCount": "34",
				"Phone":        {"+188854321"},
				"Active":       {"on"},
				"OtherModelID": {fmt.Sprint(m5.ID)},
				"ModelAList":   {fmt.Sprint(m5.ID)},
				"ParentID":     {"0"},
				"Email":        {"updated@example.com"},
				"en-Greeting":  {"Hello Updated Greeting"},
				//"File":         {""},
				//"Image":  {""},
				"Secret": {"Updated Secret"},
				"P1":     {"2"},
				"P2":     {"0.2"},
				"P3":     {"0.3"},
				"P4":     {"0.4"},
				"P5":     {"0.5"},
				"P6":     {"0.6"},
				"Price":  {"100.01"},
				"save":   {"continue"},
			},
			[]attrExample{},
			[]Approval{},
		},
		//14
		{
			fileR2,
			http.StatusSeeOther,
			s1,
			map[string][]string{
				"ID":   {fmt.Sprint(m6.ID)},
				"Name": {"Updated Name"},
				//"ItemCount": "34",
				"Phone":        {"+188854321"},
				"Active":       {"on"},
				"OtherModelID": {fmt.Sprint(m5.ID)},
				"ModelAList":   {fmt.Sprint(m5.ID)},
				"ParentID":     {"0"},
				"Email":        {"updated@example.com"},
				"en-Greeting":  {"Hello Updated Greeting"},
				//"File":         {""},
				//"Image":        {""},
				"Secret": {"Updated Secret"},
				"P1":     {"2"},
				"P2":     {"0.2"},
				"P3":     {"0.3"},
				"P4":     {"0.4"},
				"P5":     {"0.5"},
				"P6":     {"0.6"},
				"Price":  {"100.01"},
				"save":   {"continue"},
			},
			[]attrExample{},
			[]Approval{},
		},
		//15
		{
			fileR3,
			http.StatusSeeOther,
			s1,
			map[string][]string{
				"ID":   {fmt.Sprint(m6.ID)},
				"Name": {"Updated Name"},
				//"ItemCount": "34",
				"Phone":        {"+188854321"},
				"Active":       {"on"},
				"OtherModelID": {fmt.Sprint(m5.ID)},
				"ModelAList":   {fmt.Sprint(m5.ID)},
				"ParentID":     {"0"},
				"Email":        {"updated@example.com"},
				"en-Greeting":  {"Hello Updated Greeting"},
				//"File":         {""},
				//"Image":        {""},
				"Secret": {"Updated Secret"},
				"P1":     {"2"},
				"P2":     {"0.2"},
				"P3":     {"0.3"},
				"P4":     {"0.4"},
				"P5":     {"0.5"},
				"P6":     {"0.6"},
				"Price":  {"100.01"},
				"save":   {"continue"},
			},
			[]attrExample{},
			[]Approval{},
		},
		//16
		{
			fileR4,
			http.StatusSeeOther,
			s1,
			map[string][]string{
				"ID":   {fmt.Sprint(m6.ID)},
				"Name": {"Updated Name"},
				//"ItemCount": "34",
				"Phone":        {"+188854321"},
				"Active":       {"on"},
				"OtherModelID": {fmt.Sprint(m5.ID)},
				"ModelAList":   {fmt.Sprint(m5.ID)},
				"ParentID":     {"0"},
				"Email":        {"updated@example.com"},
				"en-Greeting":  {"Hello Updated Greeting"},
				//"File":         {""},
				//"Image":        {""},
				"Secret": {"Updated Secret"},
				"P1":     {"2"},
				"P2":     {"0.2"},
				"P3":     {"0.3"},
				"P4":     {"0.4"},
				"P5":     {"0.5"},
				"P6":     {"0.6"},
				"Price":  {"100.01"},
				"save":   {"continue"},
			},
			[]attrExample{},
			[]Approval{},
		},
		//17
		{
			httptest.NewRequest("GET", fmt.Sprintf("/teststruct/%d", m1.ID+100), nil),
			http.StatusNotFound,
			s1,
			map[string][]string{},
			[]attrExample{},
			[]Approval{},
		},
		//18
		{
			httptest.NewRequest("GET", fmt.Sprintf("/badname/%d", m1.ID), nil),
			http.StatusNotFound,
			s1,
			map[string][]string{},
			[]attrExample{},
			[]Approval{},
		},
		//19
		{
			httptest.NewRequest("POST", fmt.Sprintf("/testapproval/%d", m7.ID), nil),
			http.StatusSeeOther,
			s2,
			map[string][]string{
				"ID":          {fmt.Sprint(m6.ID)},
				"Name":        {"Test1"},
				"Start":       {tomorrow.Format("2006-01-02 15:04")},
				"End":         {""},
				"Count":       {"2"},
				"Price":       {"2"},
				"List":        {"1"},
				"TestModelID": {""},
				"save":        {"continue"},
			},
			[]attrExample{},
			[]Approval{
				{ModelName: "testapproval", ModelPK: m6.ID, ColumnName: "Name", OldValue: m7.Name, NewValue: "Test1", NewValueDescription: "Test1", ChangedBy: "u1", ApprovalAction: ApprovalAction(1)},
				{ModelName: "testapproval", ModelPK: m6.ID, ColumnName: "Start", OldValue: m7.Start.Format("2006-01-02 15:04:05-07:00"), NewValue: tomorrow.Format("2006-01-02 15:04:05-07:00"), NewValueDescription: tomorrow.Format("2006-01-02 15:04:05-07:00"), ChangedBy: "u1", ApprovalAction: ApprovalAction(1)},
				{ModelName: "testapproval", ModelPK: m6.ID, ColumnName: "End", OldValue: m7.End.Format("2006-01-02 15:04:05-07:00"), NewValue: "", NewValueDescription: "", ChangedBy: "u1", ApprovalAction: ApprovalAction(1)},
				{ModelName: "testapproval", ModelPK: m6.ID, ColumnName: "Count", OldValue: fmt.Sprint(m7.Count), NewValue: "2", NewValueDescription: "2", ChangedBy: "u1", ApprovalAction: ApprovalAction(1)},
				{ModelName: "testapproval", ModelPK: m6.ID, ColumnName: "Price", OldValue: fmt.Sprint(m7.Price), NewValue: "2", NewValueDescription: "2", ChangedBy: "u1", ApprovalAction: ApprovalAction(1)},
				{ModelName: "testapproval", ModelPK: m6.ID, ColumnName: "List", OldValue: fmt.Sprint(m7.List), NewValue: "1", NewValueDescription: GetString(testList(1)), ChangedBy: "u1", ApprovalAction: ApprovalAction(1)},
				{ModelName: "testapproval", ModelPK: m6.ID, ColumnName: "TestModel", OldValue: fmt.Sprint(m7.TestModelID), NewValue: "0", NewValueDescription: "", ChangedBy: "u1", ApprovalAction: ApprovalAction(1)},
				{ModelName: "testapproval", ModelPK: m6.ID, ColumnName: "Active", OldValue: fmt.Sprint(m7.Active), NewValue: "false", NewValueDescription: "false", ChangedBy: "u1", ApprovalAction: ApprovalAction(1)},
			},
		},
		//20
		{
			httptest.NewRequest("GET", fmt.Sprintf("/testapproval/%d", m7.ID), nil),
			http.StatusOK,
			s2,
			map[string][]string{},
			[]attrExample{
				{"input", "name", "Name", "value", "Test1", -1, "", true},
				{"input", "name", "Start", "value", tomorrow.Format("2006-01-02 15:04:05"), -1, "", true},
				{"input", "name", "End", "value", "", -1, "", true},
				{"input", "name", "Count", "value", "2", -1, "", true},
				{"input", "name", "Price", "value", "2", -1, "", true},
				{"select", "name", "List", "name", "List", -1, "", true},
				{"option", "value", "", "selected", "", 5, "", false},
				{"option", "value", "1", "selected", "", 5, "", true},
				{"select", "name", "TestModelID", "name", "TestModelID", -1, "", true},
				{"option", "value", "", "selected", "", 8, "", true},
				{"option", "value", "1", "selected", "", 8, "", false},
				{"input", "name", "Active", "checked", "", -1, "", false},
			},
			[]Approval{},
		},
		//21
		{
			httptest.NewRequest("POST", fmt.Sprintf("/testapproval/%d", m7.ID), nil),
			http.StatusSeeOther,
			s2,
			map[string][]string{
				"ID":          {fmt.Sprint(m6.ID)},
				"Name":        {"Test2"},
				"Start":       {now.Format("2006-01-02 15:04")},
				"End":         {now.Format("2006-01-02 15:04")},
				"Count":       {"3"},
				"Price":       {"3"},
				"List":        {""},
				"TestModelID": {fmt.Sprint(m7.ID)},
				"Active":      {"on"},
				"save":        {"continue"},
			},
			[]attrExample{},
			[]Approval{
				{ModelName: "testapproval", ModelPK: m6.ID, ColumnName: "Name", OldValue: "Test1", NewValue: "Test2", NewValueDescription: "Test2", ChangedBy: "u1", ApprovalAction: ApprovalAction(2)},
				{ModelName: "testapproval", ModelPK: m6.ID, ColumnName: "Start", OldValue: tomorrow.Format("2006-01-02 15:04:05-07:00"), NewValue: now.Format("2006-01-02 15:04:05-07:00"), NewValueDescription: now.Format("2006-01-02 15:04:05-07:00"), ChangedBy: "u1", ApprovalAction: ApprovalAction(2)},
				{ModelName: "testapproval", ModelPK: m6.ID, ColumnName: "End", OldValue: "", NewValue: now.Format("2006-01-02 15:04:05-07:00"), NewValueDescription: now.Format("2006-01-02 15:04:05-07:00"), ChangedBy: "u1", ApprovalAction: ApprovalAction(2)},
				{ModelName: "testapproval", ModelPK: m6.ID, ColumnName: "Count", OldValue: "2", NewValue: "3", NewValueDescription: "3", ChangedBy: "u1", ApprovalAction: ApprovalAction(2)},
				{ModelName: "testapproval", ModelPK: m6.ID, ColumnName: "Price", OldValue: "2", NewValue: "3", NewValueDescription: "3", ChangedBy: "u1", ApprovalAction: ApprovalAction(2)},
				{ModelName: "testapproval", ModelPK: m6.ID, ColumnName: "List", OldValue: "1", NewValue: "0", NewValueDescription: "0", ChangedBy: "u1", ApprovalAction: ApprovalAction(2)},
				{ModelName: "testapproval", ModelPK: m6.ID, ColumnName: "TestModel", OldValue: "0", NewValue: fmt.Sprint(m7.ID), NewValueDescription: GetString(m7), ChangedBy: "u1", ApprovalAction: ApprovalAction(2)},
				{ModelName: "testapproval", ModelPK: m6.ID, ColumnName: "Active", OldValue: "false", NewValue: "true", NewValueDescription: "true", ChangedBy: "u1", ApprovalAction: ApprovalAction(2)},
			},
		},
		//22
		{
			httptest.NewRequest("GET", fmt.Sprintf("/testapproval/%d", m7.ID), nil),
			http.StatusOK,
			s2,
			map[string][]string{},
			[]attrExample{
				{"input", "name", "Name", "value", "Test1", -1, "", true},
				{"input", "name", "Start", "value", tomorrow.Format("2006-01-02 15:04:05"), -1, "", true},
				{"input", "name", "End", "value", "", -1, "", true},
				{"input", "name", "Count", "value", "2", -1, "", true},
				{"input", "name", "Price", "value", "2", -1, "", true},
				{"select", "name", "List", "name", "List", -1, "", true},
				{"option", "value", "", "selected", "", 5, "", false},
				{"option", "value", "1", "selected", "", 5, "", true},
				{"select", "name", "TestModelID", "name", "TestModelID", -1, "", true},
				{"option", "value", "", "selected", "", 8, "", true},
				{"option", "value", "1", "selected", "", 8, "", false},
				{"input", "name", "Active", "checked", "", -1, "", false},
			},
			[]Approval{},
		},
	}

	loadSessions()
	loadPermissions()

	for i, e := range examples {
		w = httptest.NewRecorder()

		if e.r.Form == nil {
			e.r.Form = url.Values{}
		}
		for k, v := range e.postParam {
			e.r.Form[k] = v
		}

		formHandler(w, e.r, e.s)

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
					t.Errorf("formHandler returned attrribue %s=%s for attr %s. Expected(%v) %#v, got (%#v) for %s(%d-%d)", tempAttr.selectorKey, tempAttr.selectorValue, tempAttr.checkKey, tempAttr.expected, tempAttr.checkValue, tempValue, tag, i, counter)
				} else {
					if index != -1 {
						e.attr[counter].path = path[index]
					}
				}
			}
		}

		// Check for approvals
		var tempApp Approval
		q := "model_name = ? AND model_pk = ? AND column_name = ? AND old_value = ? AND new_value = ? AND new_value_description = ? AND changed_by = ?"
		for _, app := range e.approvals {
			tempApp = Approval{}
			Get(&tempApp, q, app.ModelName, app.ModelPK, app.ColumnName, app.OldValue, app.NewValue, app.NewValueDescription, app.ChangedBy)
			if tempApp.ID == 0 {
				t.Errorf("formHandler didn't create the correct approval for example %d for values. %#v", i, app)
			}

			if app.ApprovalAction != ApprovalAction(0) {
				tempApp.ApprovalAction = app.ApprovalAction
				tempApp.UpdatedBy = "admin"
				tempApp.Save()
			}
		}
	}

	// Clean up
	Delete(s1)
	Delete(s2)
	Delete(s3)
	Delete(m1)
	Delete(m2)
	Delete(m3)
	Delete(m4)
	Delete(m5)
	Delete(m6)
	Delete(m7)
	Delete(u1)
	Delete(u2)
	Delete(g1)
	Delete(up1)
	Delete(gp1)
	Delete(gp2)
	Delete(gp3)
}

func checkTagAttr(selectorKey string, selectorValue string, checkKey string, checkValue string, attr []map[string]string, path []string, pathPrefix string) (int, string) {
	for i := range attr {
		if tempName, ok := attr[i][selectorKey]; ok && tempName == selectorValue {
			if tempValue, ok := attr[i][checkKey]; ok {
				if tempValue != checkValue {
					continue
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
