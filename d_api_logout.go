package uadmin

import "net/http"

func dAPILogoutHandler(w http.ResponseWriter, r *http.Request, s *Session) {
	if s == nil {
		w.WriteHeader(http.StatusUnauthorized)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "User not logged in",
		})
		return
	}

	if CheckCSRF(r) {
		w.WriteHeader(http.StatusUnauthorized)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Missing CSRF token",
		})
		return
	}

	Logout(r)
	ReturnJSON(w, r, map[string]interface{}{
		"status": "ok",
	})
}
