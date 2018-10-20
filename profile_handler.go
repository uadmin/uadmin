package uadmin

import (
	"fmt"
	"html/template"
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
		Translation  ProfileTranslation
		Direction    string
		RootURL      string
		ProfilePhoto string
	}

	c := Context{}
	c.RootURL = RootURL

	language := getLanguage(r)
	c.Direction = language.Direction

	c.Translation.SaveChanges = translateUI(language.Code, "savechanges")
	c.Translation.ChangePassword = translateUI(language.Code, "changepassword")
	c.Translation.Logout = translateUI(language.Code, "logout")
	c.Translation.OldPassword = translateUI(language.Code, "oldpassword")
	c.Translation.NewPassword = translateUI(language.Code, "newpassword")
	c.Translation.ConfirmPassword = translateUI(language.Code, "confirmpassword")
	c.Translation.ApplyChanges = translateUI(language.Code, "applychanges")
	c.Translation.Close = translateUI(language.Code, "close")
	c.Translation.Profile = translateUI(language.Code, "profile")

	c.SiteName = SiteName

	user := session.User
	c.ProfilePhoto = session.User.Photo

	s, _ := getSchema(user)
	r.Form.Set("ModelID", fmt.Sprint(user.ID))
	c.Schema = getFormData(user, r, session, s)

	t := template.New("") //create a new template
	t, err := t.ParseFiles("./templates/uadmin/" + Theme + "/profile.html")

	if err != nil {
		fmt.Fprint(w, err.Error())
		fmt.Println("ERROR", err.Error())
		return
	}

	if r.Method == cPOST {
		c.IsUpdated = true
		if r.FormValue("save") == "" {
			user.Username = r.FormValue("Username")
			user.FirstName = r.FormValue("FirstName")
			user.LastName = r.FormValue("LastName")
			user.Email = r.FormValue("Email")
			f := c.Schema.FieldMyName("Photo")
			user.Photo, c.Schema = processUpload(r, f, "user", session, c.Schema)
			user.Save()
		}
		if r.FormValue("save") == "password" {
			oldPassword := r.FormValue("oldPassword")
			newPassword := r.FormValue("newPassword")
			confirmPassword := r.FormValue("confirmPassword")
			// TODO: Add OTP logic
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

	err = t.ExecuteTemplate(w, "profile.html", c)
	if err != nil {
		fmt.Println(err.Error())
	}

	return
}
