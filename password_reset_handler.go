package uadmin

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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
	fmt.Println(r.FormValue("u"), r.FormValue("key"))
	userID, err := strconv.ParseUint(r.FormValue("u"), 10, 64)
	if err != nil {
		page404Handler(w, r)
		return
	}
	user := &User{}
	Get(user, "id = ?", userID)
	if user.ID == 0 {
		go func() {
			log := &Log{}
			r.Form.Set("reset-status", "invalid user id")
			log.PasswordReset(r.FormValue("u"), log.Action.PasswordResetDenied(), r)
			log.Save()
		}()
		page404Handler(w, r)

		return
	}
	otpCode := r.FormValue("key")
	if !user.VerifyOTP(otpCode) {
		go func() {
			log := &Log{}
			r.Form.Set("reset-status", "invalid otpcode: "+otpCode)
			log.PasswordReset(user.Username, log.Action.PasswordResetDenied(), r)
			log.Save()
		}()
		page404Handler(w, r)
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
		Trail(ERROR, "Unable to parse resetpassword.html. %s", err)
		page404Handler(w, r)
		return
	}

	// URL, _ := url.Parse(r.RequestURI)
	// URLPath := strings.Split(URL.Path, "/")

	err = t.ExecuteTemplate(w, "resetpassword.html", c)
	// lock.Unlock()
	if err != nil {
		fmt.Println(err.Error())
	}
}
