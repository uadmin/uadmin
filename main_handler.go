package uadmin

import (
	"net/http"
	"net/url"
	"strings"
)

// mainHandler is the main handler for the admin
func mainHandler(w http.ResponseWriter, r *http.Request) {
	if !CheckRateLimit(r) {
		w.Write([]byte("Slow down. You are going too fast!"))
		return
	}
	if !ValidateIP(r, AllowedIPs, BlockedIPs) {
		if r.Form == nil {
			r.Form = url.Values{}
		}
		r.Form.Set("err_msg", "Your IP Address ("+r.RemoteAddr+") is not Allowed to Access this Page")
		r.Form.Set("err_code", "403")
		pageErrorHandler(w, r, nil)
		return
	}
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
		if r.Form == nil {
			r.Form = url.Values{}
		}
		r.Form.Set("err_msg", "Remote Access Denied")
		pageErrorHandler(w, r, nil)
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
		if URLParts[0] == "settings" {
			settingsHandler(w, r, session)
			return
		}
		listHandler(w, r, session)
		return
	} else if len(URLParts) == 2 {
		formHandler(w, r, session)
		return
	}
	pageErrorHandler(w, r, session)
}
