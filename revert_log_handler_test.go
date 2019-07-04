package uadmin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestRevertLogHandler(t *testing.T) {
	mA1 := TestModelA{}
	Save(&mA1)
	mB1 := TestModelB{}
	Save(&mB1)
	mB2 := TestModelB{
		Name:         "Test",
		ItemCount:    1,
		Phone:        "09999999999",
		Active:       true,
		OtherModelID: mA1.ID,
		ModelAList:   []TestModelA{mA1},
		ParentID:     mB1.ID,
		Email:        "test@example.com",
		Greeting:     `"en":"Hello"`,
		//TODO File and Image
		Secret:      "pass1",
		Description: "<div>Test</div>",
		URL:         "/test/1",
		Code:        "10 PRINT HELLO",
		P1:          10,
		P2:          0.1,
		P3:          0.2,
		P4:          0.3,
		P5:          0.4,
		P6:          0.5,
		Price:       100.0,
		List:        testList(1),
	}
	Save(&mB2)
	backup := deepCopy(mB2).(TestModelB)
	mB2.Name += "1"
	mB2.ItemCount++
	mB2.Phone += "1"
	mB2.Active = false
	mB2.OtherModelID = 0
	mB2.ModelAList = []TestModelA{}
	mB2.ParentID = 0
	mB2.Email += "1"
	mB2.Greeting = `"en":"Hello1"`
	mB2.Secret += "1"
	mB2.Description += "1"
	mB2.URL += "1"
	mB2.Code += "1"
	mB2.P1++
	mB2.P2++
	mB2.P3++
	mB2.P4++
	mB2.P5++
	mB2.P6++
	mB2.Price += 1.0
	mB2.List = testList(0)
	_ = backup

	// Send POST request to update mB2
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", fmt.Sprintf("/testmodelb/%d", mB2.ID), nil)
	r.Form = url.Values{}
	r.Form["ID"] = []string{fmt.Sprint(mB2.ID)}
	r.Form["Name"] = []string{mB2.Name}
	r.Form["ItemCount"] = []string{fmt.Sprint(mB2.ItemCount)}
	r.Form["Phone"] = []string{mB2.Phone}
	r.Form["OtherModelID"] = []string{fmt.Sprint(mB2.OtherModelID)}
	r.Form["ModelAList"] = func() []string {
		v := []string{}
		for i := range mB2.ModelAList {
			v = append(v, fmt.Sprint(mB2.ModelAList[i].ID))
		}
		return v
	}()
	r.Form["ParentID"] = []string{fmt.Sprint(mB2.ParentID)}
	r.Form["Email"] = []string{mB2.Email}
	r.Form["en-Greeting"] = []string{"Hello1"}
	r.Form["Secret"] = []string{mB2.Secret}
	r.Form["Description"] = []string{mB2.Description}
	r.Form["URL"] = []string{mB2.URL}
	r.Form["Code"] = []string{mB2.Code}
	r.Form["P1"] = []string{fmt.Sprint(mB2.P1)}
	r.Form["P2"] = []string{fmt.Sprint(mB2.P2)}
	r.Form["P3"] = []string{fmt.Sprint(mB2.P3)}
	r.Form["P4"] = []string{fmt.Sprint(mB2.P4)}
	r.Form["P5"] = []string{fmt.Sprint(mB2.P5)}
	r.Form["P6"] = []string{fmt.Sprint(mB2.P6)}
	r.Form["Price"] = []string{fmt.Sprint(mB2.Price)}
	r.Form["List"] = []string{fmt.Sprint(mB2.List)}

	s1 := &Session{
		UserID: 1,
		Active: true,
	}
	s1.GenerateKey()
	s1.Save()
	Preload(s1)

	formHandler(w, r, s1)

	log := Log{}
	Get(&log, "table_name = ? AND table_id = ? AND action = ?", "testmodelb", mB2.ID, log.Action.Modified())

	// Send a request with no session, it should return 404
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", RootURL+"revertHandler/?log_id="+fmt.Sprint(log.ID), nil)

	revertLogHandler(w, r)

	// Check if we are getting code 404
	if w.Code != http.StatusNotFound {
		t.Errorf("revertLogHandler return invalid code. Got %d expected %d", w.Code, http.StatusNotFound)
	}

	// Send a request from a user who does not have permission to logs model
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", RootURL+"revertHandler/?log_id="+fmt.Sprint(log.ID), nil)
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
	r.AddCookie(&http.Cookie{Name: "session", Value: s2.Key})

	revertLogHandler(w, r)

	// Check if we are getting code 404
	if w.Code != http.StatusNotFound {
		t.Errorf("revertLogHandler return invalid code. Got %d expected %d", w.Code, http.StatusNotFound)
	}

	// Send a request from a user with permission to logs but no log ID
	// This should return a 404
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", RootURL+"revertHandler/", nil)
	r.AddCookie(&http.Cookie{Name: "session", Value: s1.Key})
	r.ParseForm()

	revertLogHandler(w, r)

	// Check if we are getting code 404
	if w.Code != http.StatusNotFound {
		t.Errorf("revertLogHandler return invalid code. Got %d expected %d", w.Code, http.StatusNotFound)
	}

	// Send a request from a user with permission
	// This should return a 200
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", RootURL+"revertHandler/?log_id="+fmt.Sprint(log.ID), nil)
	r.AddCookie(&http.Cookie{Name: "session", Value: s1.Key})
	r.ParseForm()

	revertLogHandler(w, r)

	// Check if we are getting code 200
	if w.Code != http.StatusOK {
		t.Errorf("revertLogHandler return invalid code. Got %d expected %d", w.Code, http.StatusOK)
	}

	// Check if the values in the DB are updated
	mB2DB := TestModelB{}
	Get(&mB2DB, "id = ?", mB2.ID)
	if mB2DB.Name != backup.Name {
		t.Errorf("revertLogHandler didn't return value after edit for Name. Got (%s) expected (%s)", mB2DB.Name, backup.Name)
	}
	//TODO: Check the rest of the fields

	// Delete the record
	r = httptest.NewRequest("POST", "/testmodelb/", nil)
	r.Form = url.Values{}
	r.Form["delete"] = []string{"delete"}
	r.Form["listID"] = []string{fmt.Sprint(mB2.ID)}

	listHandler(w, r, s1)

	// get the log of the delete action
	log = Log{}
	Get(&log, "table_name = ? AND table_id = ? AND action = ?", "testmodelb", mB2.ID, log.Action.Deleted())

	// Send a request to undelete the record
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", RootURL+"revertHandler/?log_id="+fmt.Sprint(log.ID), nil)
	r.AddCookie(&http.Cookie{Name: "session", Value: s1.Key})
	r.ParseForm()

	revertLogHandler(w, r)

	// Check if we are getting code 200
	if w.Code != http.StatusOK {
		t.Errorf("revertLogHandler return invalid code. Got %d expected %d", w.Code, http.StatusOK)
	}

	// Check if the record undeleted
	mB2DB = TestModelB{}
	Get(&mB2DB, "id = ?", mB2.ID)

	if mB2DB.ID == 0 {
		t.Errorf("revertLogHandler didn't undelete the record")
	}

	//Clean Up
	Delete(s1)
	Delete(s2)
	Delete(u1)
	Delete(mA1)
	Delete(mB1)
	Delete(mB2)

}
