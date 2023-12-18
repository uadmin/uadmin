package uadmin

import (
	"net/http"
)

func dAPIAuthHandler(w http.ResponseWriter, r *http.Request, s *Session) {

	if DisableDAPIAuth {
		w.WriteHeader(http.StatusForbidden)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "dAPI auth is disabled",
		})
		return
	}

	modelKV := r.Context().Value(CKey("modelName")).(DApiModelKeyVal)
	command := modelKV.CommandName

	switch command {
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
	case "openidlogin":
		dAPIOpenIDLoginHandler(w, r, s)
	case "certs":
		dAPIOpenIDCertHandler(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Unknown auth command: (" + r.URL.Path + ")",
		})
	}
}
