package uadmin

import (
	"net/http"
	"strings"
)

func dAPIAuthHandler(w http.ResponseWriter, r *http.Request, s *Session) {
	// Trim leading path
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "auth")
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

	switch r.URL.Path {
	case "login":
		dAPILoginHandler(w, r, s)
	case "logout":
		dAPILogoutHandler(w, r, s)
	case "signup":
		dAPISignupHandler(w, r, s)
	case "resetpassword":
		dAPIResetPasswordHandler(w, r, s)
	case "changepassword":
		dAPIChangePasswordHandler(w, r, s)
	default:
		w.WriteHeader(http.StatusNotFound)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Unknown auth command: (" + r.URL.Path + ")",
		})
	}
}
