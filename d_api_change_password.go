package uadmin

import "net/http"

func dAPIChangePasswordHandler(w http.ResponseWriter, r *http.Request, s *Session) {
	if s == nil {
		w.WriteHeader(http.StatusForbidden)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "",
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

	oldPassword := r.FormValue("old_password")
	newPassword := r.FormValue("new_password")

	// Check if there is a new password
	if newPassword == "" {
		w.WriteHeader(http.StatusBadRequest)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Missing new password",
		})
		return
	}

	// Verify old password
	err := verifyPassword(s.User.Password, oldPassword)
	if err != nil {
		incrementInvalidLogins(r)
		w.WriteHeader(401)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "invalid password",
		})
		return
	}

	s.User.Password = newPassword
	s.User.Save()

	ReturnJSON(w, r, map[string]interface{}{
		"status": "ok",
	})
}
