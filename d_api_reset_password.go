package uadmin

import (
	"net/http"
	"time"
)

func dAPIResetPasswordHandler(w http.ResponseWriter, r *http.Request, s *Session) {
	// Get parameters
	username := r.FormValue("username")
	email := r.FormValue("email")
	otp := r.FormValue("otp")
	password := r.FormValue("password")

	// check if there is an email or a username
	if username == "" && email == "" {
		w.WriteHeader(400)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "No username nor email",
		})
		// log the request
		go func() {
			log := &Log{}
			if password != "" {
				r.Form.Set("password", "*****")
			}
			log.PasswordReset("", log.Action.PasswordResetDenied(), r)
			log.Save()
		}()
		return
	}

	// get user
	user := User{}
	if username != "" {
		Get(&user, "username = ? AND active = ?", username, true)
	} else {
		Get(&user, "email = ? AND active = ?", email, true)
	}

	// log the request
	go func() {
		log := &Log{}
		if password != "" {
			r.Form.Set("password", "*****")
		}
		log.PasswordReset(user.Username, log.Action.PasswordResetRequest(), r)
		log.Save()
	}()

	// check if the user exists and active
	if user.ID == 0 || (user.ExpiresOn != nil || user.ExpiresOn.After(time.Now())) {
		w.WriteHeader(404)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "username or email do not match any active user",
		})
		// log the request
		go func() {
			log := &Log{}
			if password != "" {
				r.Form.Set("password", "*****")
			}
			log.PasswordReset(user.Username, log.Action.PasswordResetDenied(), r)
			log.Save()
		}()
		return
	}

	// If there is no otp, then we assume this is a request to send a password
	// reset email
	if otp == "" {
		err := forgotPasswordHandler(&s.User, r, CustomResetPasswordLink, ResetPasswordMessage)

		if err != nil {
			w.WriteHeader(403)
			ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			// log the request
			go func() {
				log := &Log{}
				if password != "" {
					r.Form.Set("password", "*****")
				}
				log.PasswordReset(user.Username, log.Action.PasswordResetDenied(), r)
				log.Save()
			}()
			return
		}
		// log the request
		w.WriteHeader(http.StatusAccepted)
		go func() {
			log := &Log{}
			if password != "" {
				r.Form.Set("password", "*****")
			}
			r.Form.Set("reset-status", "Email was sent with the OTP")
			log.PasswordReset(user.Username, log.Action.PasswordResetSuccessful(), r)
			log.Save()
		}()
		ReturnJSON(w, r, map[string]interface{}{
			"status": "ok",
		})
		return
	}

	// Since there is an OTP, we can check it and reset the password
	// Check if there is a a new password
	if password == "" {
		w.WriteHeader(400)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "missing password",
		})
		// log the request
		go func() {
			log := &Log{}
			if password != "" {
				r.Form.Set("password", "*****")
			}
			log.PasswordReset("", log.Action.PasswordResetDenied(), r)
			log.Save()
		}()
		return
	}

	// check OTP
	if !user.VerifyOTP(otp) {
		incrementInvalidLogins(r)
		w.WriteHeader(401)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "invalid or expired OTP",
		})
		// log the request
		go func() {
			log := &Log{}
			if password != "" {
				r.Form.Set("password", "*****")
			}
			log.PasswordReset("", log.Action.PasswordResetDenied(), r)
			log.Save()
		}()
		return
	}

	// reset the password
	user.Password = password
	user.Save()

	// log the request
	go func() {
		log := &Log{}
		if password != "" {
			r.Form.Set("password", "*****")
		}
		r.Form.Set("reset-status", "Successfully changed the password")
		log.PasswordReset("", log.Action.PasswordResetSuccessful(), r)
		log.Save()
	}()
	ReturnJSON(w, r, map[string]interface{}{
		"status": "ok",
	})
}
