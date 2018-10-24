package uadmin

import (
	"net/http"
	"strings"
)

// mainHandler is the main handler for the admin
func mainHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, RootURL)
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
	URLParts := strings.Split(r.URL.Path, "/")

	if URLParts[0] == "resetpassword" {
		passwordResetHandler(w, r)
		return
	}

	// Authentecation
	// This session is preloaded with a user
	session := IsAuthenticated(r)
	if session == nil {
		loginHandler(w, r)
		return
	}

	// Check remote access
	if !(isLocal(r.RemoteAddr) || session.User.RemoteAccess) {
		w.Write([]byte("REMOTE ACCESS DENIED"))
		return
	}

	if r.URL.Path == "" {
		homeHandler(w, r, session)
		return
	}
	if len(URLParts) == 1 {
		if URLParts[0] == "logout" {
			logoutHandler(w, r, session)
			return
		}
		if URLParts[0] == "export" {
			exportHandler(w, r, session)
			return
		}
		if URLParts[0] == "cropper" {
			cropImageHandler(w, r, session)
			return
		}
		if URLParts[0] == "profile" {
			profileHandler(w, r, session)
			return
		}
		listHandler(w, r, session)
		return
	} else if len(URLParts) == 2 {
		formHandler(w, r, session)
		return
	}
	page404Handler(w, r, session)
}
