package uadmin

import (
	"fmt"
	"net/http"
)

// profileHandler !
func profileHandler(w http.ResponseWriter, r *http.Request, session *Session) {
	/*
		http://domain.com/admin/profile/
	*/
	r.ParseMultipartForm(32 << 20)
	type Context struct {
		User         string
		ID           uint
		Schema       ModelSchema
		Status       bool
		IsUpdated    bool
		Notif        string
		Demo         bool
		SiteName     string
		Language     Language
		RootURL      string
		ProfilePhoto string
		OTPImage     string
		OTPRequired  bool
	}

	c := Context{}
	c.RootURL = RootURL
	c.Language = getLanguage(r)
	c.SiteName = SiteName
	user := session.User
	c.User = user.Username
	c.ProfilePhoto = session.User.Photo
	c.OTPImage = "/media/otp/" + session.User.OTPSeed + ".png"

	// Check if OTP Required has been changed
	if r.URL.Query().Get("otp_required") != "" {
		if r.URL.Query().Get("otp_required") == "1" {
			user.OTPRequired = true
		} else if r.URL.Query().Get("otp_required") == "0" {
			user.OTPRequired = false
		}
		r.URL.RawQuery = ""
		(&user).Save()
		c.OTPImage = "/media/otp/" + user.OTPSeed + ".png"
	}

	c.OTPRequired = user.OTPRequired

	c.Schema, _ = getSchema(user)
	r.Form.Set("ModelID", fmt.Sprint(user.ID))
	getFormData(user, r, session, &c.Schema, &user)

	if r.Method == cPOST {
		c.IsUpdated = true
		if r.FormValue("save") == "" {
			user.Username = r.FormValue("Username")
			user.FirstName = r.FormValue("FirstName")
			user.LastName = r.FormValue("LastName")
			user.Email = r.FormValue("Email")
			f := c.Schema.FieldByName("Photo")
			if _, _, err := r.FormFile("Photo"); err == nil {
				user.Photo = processUpload(r, f, "user", session, &c.Schema)
			}
			(&user).Save()
			c.ProfilePhoto = user.Photo
		}
		if r.FormValue("save") == "password" {
			oldPassword := r.FormValue("oldPassword")
			newPassword := r.FormValue("newPassword")
			confirmPassword := r.FormValue("confirmPassword")
			_session := user.Login(oldPassword, "")

			if _session == nil || !user.Active {
				c.Status = true
				c.Notif = "Incorrent old password."
			} else if newPassword != confirmPassword {
				c.Status = true
				c.Notif = "New password and confirm password do not match."
			} else {
				user.Password = hashPass(newPassword)
				user.Save()

				// To logout
				Logout(r)

				return
			}
		}
	}

	RenderHTML(w, r, "./templates/uadmin/"+Theme+"/profile.html", c)
	return
}
