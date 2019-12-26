package uadmin

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// TestSettingsHandler is a unit testing function for settingsHandler() function
func TestSettingsHandler(t *testing.T) {
	CacheSessions = false
	CachePermissions = false
	r := httptest.NewRequest("GET", "http://0.0.0.0:5000/settings", nil)
	w := httptest.NewRecorder()

	settingsList := []Setting{}
	Filter(&settingsList, "data_type <> ?", DataType(0).Boolean())

	boolSettings := []Setting{}
	Filter(&boolSettings, "data_type = ? AND value = ?", DataType(0).Boolean(), "1")

	// First check if the request is sent to 404 page if not authenticated
	settingsHandler(w, r, nil)
	if w.Code != http.StatusNotFound {
		t.Errorf("Invalid code on requesting /settings. %d, expected %d", w.Code, http.StatusNotFound)
	}

	// Add a session to the request authenticated as admin
	s1 := &Session{
		Active: true,
		UserID: 1,
	}
	s1.GenerateKey()
	s1.Save()
	Preload(s1)

	w = httptest.NewRecorder()
	settingsHandler(w, r, s1)

	if w.Code != http.StatusOK {
		t.Errorf("Invalid code on requesting /settings. %d expected %d", w.Code, http.StatusOK)
	}

	// Add a user with no permissions to settings model and check if
	// the request returns a 404
	u1 := &User{
		Username: "u1",
		Password: "u1",
		Active:   true,
	}
	u1.Save()

	s2 := &Session{
		Active: true,
		UserID: u1.ID,
	}
	s2.GenerateKey()
	s2.Save()
	Preload(s2)

	w = httptest.NewRecorder()
	settingsHandler(w, r, s2)
	if w.Code != http.StatusNotFound {
		t.Errorf("Invalid code on requesting /settings. %d, expected %d", w.Code, http.StatusNotFound)
	}

	// Send a request with a user who has read permissions. It should return the settings page
	dash := DashboardMenu{}
	Get(&dash, "url = ?", "setting")
	perm1 := &UserPermission{
		DashboardMenuID: dash.ID,
		UserID:          u1.ID,
		Read:            true,
	}
	Save(perm1)

	w = httptest.NewRecorder()
	settingsHandler(w, r, s2)
	if w.Code != http.StatusOK {
		t.Errorf("Invalid code on requesting /settings. %d expected %d", w.Code, http.StatusOK)
	}

	// Send a request with a user read permissions and no edit. The request is to save.
	// This should return a 404

	r.Method = "POST"
	w = httptest.NewRecorder()
	settingsHandler(w, r, s2)
	if w.Code != http.StatusNotFound {
		t.Errorf("Invalid code on requesting /settings. %d expected %d", w.Code, http.StatusNotFound)
	}

	// Add edit permission to the user and test again. This time it should return 200
	perm1.Edit = true
	Save(perm1)
	r.Form = url.Values{}
	for _, val := range boolSettings {
		r.Form[val.Code] = []string{"on"}
	}
	for _, val := range settingsList {
		r.Form[val.Code] = []string{val.Value}
	}
	r.Form["uAdmin.SiteName"] = []string{"Test Site"}

	w = httptest.NewRecorder()
	settingsHandler(w, r, s2)
	if w.Code != http.StatusOK {
		t.Errorf("Invalid code on requesting /settings. %d expected %d", w.Code, http.StatusOK)
	}
	setting := Setting{}
	Get(&setting, "code = ?", "uAdmin.SiteName")
	if setting.Value != "Test Site" {
		t.Errorf("Setting was not saved. expected %s, got %s", "Test Site", setting.Value)
	}

	// Clearn up
	Delete(s1)
	Delete(s2)
	Delete(u1)
	Delete(perm1)

	CacheSessions = true
	CachePermissions = true
}
