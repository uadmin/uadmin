package uadmin

import (
	"net/http"
	"strings"
)

// loginHandler HTTP handeler for verifying login data and creating sessions for users
func loginHandler(w http.ResponseWriter, r *http.Request) {
	type Context struct {
		Err         string
		ErrExists   bool
		SiteName    string
		Languages   []Language
		RootURL     string
		OTPRequired bool
		Language    Language
	}

	c := Context{}
	c.SiteName = SiteName
	c.RootURL = RootURL
	c.Language = getLanguage(r)

	if r.Method == cPOST {
		if r.FormValue("save") == "Send Request" {
			email := r.FormValue("email")
			user := User{}
			Get(&user, "Email = ?", email)
			if user.ID != uint(0) {
				c.ErrExists = true
				c.Err = "Password recovery request sent. Please check email to reset your password"
				forgotPasswordHandler(&user, r)
			} else {
				c.ErrExists = true
				c.Err = "Please check email address. Email address must be associated with the account to be recovered."
			}
		} else {
			username := r.PostFormValue("username")
			username = strings.ToLower(username)
			password := r.PostFormValue("password")
			otp := r.PostFormValue("otp")
			lang := r.PostFormValue("language")
			user := User{}
			Get(&user, "username = ?", username)
			if user.ID == 0 {
				c.ErrExists = true
				c.Err = "Invalid username"
				// Store login denied to the user log
				go func() {
					log := &Log{}
					r.Form.Set("login-status", "invalid username")
					log.SignIn(username, log.Action.LoginDenied(), r)
					log.Save()
				}()
			} else {
				if user.OTPRequired {
					Trail(INFO, "User: %s OTP: %s", user.Username, user.GetOTP())
				}
				session := user.Login(password, otp)
				if session == nil || !user.Active {
					c.ErrExists = true
					c.Err = "Invalid password or inactive user"
					go func() {
						log := &Log{}
						r.Form.Set("login-status", "invalid password or inactive user")
						log.SignIn(username, log.Action.LoginDenied(), r)
						log.Save()
					}()
				} else {
					cookie, _ := r.Cookie("session")
					if cookie == nil {
						cookie = &http.Cookie{}
					}
					cookie.Name = "session"
					cookie.Value = session.Key
					//cookie.Secure = true
					cookie.Path = "/"
					http.SetCookie(w, cookie)

					// set language cookie
					cookie, _ = r.Cookie("language")
					if cookie == nil {
						cookie = &http.Cookie{}
					}
					cookie.Name = "language"
					cookie.Value = lang
					//cookie.Secure = true
					cookie.Path = "/"
					http.SetCookie(w, cookie)

					// Check for OTP
					if session.PendingOTP {
						c.OTPRequired = true
					} else {
						// Store login successful to the user log
						go func() {
							log := &Log{}
							log.SignIn(user.Username, log.Action.LoginSuccessful(), r)
							log.Save()
						}()
						if r.URL.Query().Get("next") == "" {
							nextURL := strings.TrimSuffix(r.RequestURI, "logout")
							http.Redirect(w, r, nextURL, 303)
							return
						}
						http.Redirect(w, r, r.URL.Query().Get("next"), 303)
						return
					}
				}
			}
		}
	}
	c.Languages = activeLangs

	RenderHTML(w, r, "./templates/uadmin/"+Theme+"/login.html", c)
}
