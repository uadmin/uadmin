package uadmin

import (
	"net/http"
	"strings"
)

// apiHandler !
func apiHandler(w http.ResponseWriter, r *http.Request) {
	session := IsAuthenticated(r)
	Path := strings.TrimPrefix(r.URL.Path, RootURL+"api")

	// Handle requests for dAPI
	if strings.HasPrefix(Path, "/d/") || Path == "/d" {
		if session != nil {
			session.ThroughAPI = true
		}
		dAPIHandler(w, r, session)
		return
	}

	if DisableAdminUI {
		return
	}

	// For all other APIs, if the user is not authenticated
	// then send them to login page
	if session == nil {
		loginHandler(w, r)
		return
	}

	if strings.HasPrefix(Path, "/trail") {
		trailAPIHandler(w, r)
		return
	}

	if strings.HasPrefix(Path, "/upload_image") {
		UploadImageHandler(w, r, session)
		return
	}
	if strings.HasPrefix(Path, "/search") {
		searchApiHandler(w, r, session)
		return
	}
	if strings.HasPrefix(Path, "/get_models") {
		GetModelsAPI(w, r, session)
		return
	}
	if strings.HasPrefix(Path, "/get_fields") {
		GetFieldsAPI(w, r, session)
		return
	}
}
