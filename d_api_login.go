package uadmin

import "net/http"

func dAPILoginHandler(w http.ResponseWriter, r *http.Request, s *Session) {
	_ = s

	// Get request variables
	username := r.FormValue("username")
	password := r.FormValue("password")
	otp := r.FormValue("otp")
	session := r.FormValue("session")

	optRequired := false
	if otp != "" {
		// Check if there is username and password or a session key
		if session != "" {
			s = Login2FAKey(r, session, otp)
		} else {
			s = Login2FA(r, username, password, otp)
		}
	} else {
		s, optRequired = Login(r, username, password)
	}

	if optRequired {
		w.WriteHeader(http.StatusAccepted)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "OTP Required",
			"session": s.Key,
		})
		return
	}

	if s == nil {
		w.WriteHeader(http.StatusUnauthorized)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Invalid credentials",
		})
		return
	}

	// Preload the user to get the group name
	Preload(&s.User)

	jwt := SetSessionCookie(w, r, s)
	res := map[string]interface{}{
		"status":  "ok",
		"session": s.Key,
		"jwt":     jwt,
		"user": map[string]interface{}{
			"username":   s.User.Username,
			"first_name": s.User.FirstName,
			"last_name":  s.User.LastName,
			"group_name": s.User.UserGroup.GroupName,
			"admin":      s.User.Admin,
		},
	}
	if CustomDAPILoginHandler != nil {
		res = CustomDAPILoginHandler(r, &s.User, res)
	}
	ReturnJSON(w, r, res)
}
