package uadmin

import (
	"html/template"
	"net/http"
)

func passwordResetHandler(w http.ResponseWriter, r *http.Request) {
	var err error
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
		page404Handler(w, r, nil)
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
		page404Handler(w, r, nil)
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

	t := template.New("").Funcs(template.FuncMap{
		"Tf": Tf,
	})

	t, err = t.ParseFiles("./templates/uadmin/" + Theme + "/resetpassword.html")
	if err != nil {
		Trail(ERROR, "passwordResetHandler unable to parse resetpassword.html. %s", err)
		page404Handler(w, r, nil)
		return
	}

	err = t.ExecuteTemplate(w, "resetpassword.html", c)
	if err != nil {
		Trail(ERROR, "passwordResetHandler unable to execute template. %s", err)
	}
}
