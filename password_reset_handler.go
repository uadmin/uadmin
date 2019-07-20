package uadmin

import (
	"net/http"
)

func passwordResetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	type Context struct {
		Err       string
		ErrExists bool
		SiteName  string
		Language  Language
		RootURL   string
	}

	c := Context{}
	c.SiteName = SiteName
	c.RootURL = RootURL
	c.Language = getLanguage(r)

	// Get the user and the code and verify them
	userID := r.FormValue("u")

	user := &User{}
	Get(user, "id = ?", userID)
	if user.ID == 0 {
		go func() {
			log := &Log{}
			r.Form.Set("reset-status", "invalid user id")
			log.PasswordReset(userID, log.Action.PasswordResetDenied(), r)
			log.Save()
		}()
		pageErrorHandler(w, r, nil)
		return
	}
	otpCode := r.FormValue("key")
	if !user.VerifyOTP(otpCode) {
		go func() {
			log := &Log{}
			r.Form.Set("reset-status", "invalid otp code: "+otpCode)
			log.PasswordReset(user.Username, log.Action.PasswordResetDenied(), r)
			log.Save()
		}()
		pageErrorHandler(w, r, nil)
		return
	}

	if r.Method == cPOST {
		if r.FormValue("password") != r.FormValue("confirm_password") {
			c.ErrExists = true
			c.Err = "Password does not match the confirm password"
		} else {
			user.Password = hashPass(r.FormValue("password"))
			user.Save()
			//log successful password reset
			go func() {
				log := &Log{}
				r.Form.Set("reset-status", "Successfully changed the password")
				log.PasswordReset(user.Username, log.Action.PasswordResetSuccessful(), r)
				log.Save()
			}()
			http.Redirect(w, r, RootURL, 303)
			return
		}
	}

	RenderHTML(w, r, "./templates/uadmin/"+Theme+"/resetpassword.html", c)
}
