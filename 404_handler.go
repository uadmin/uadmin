package uadmin

import (
	"fmt"
	"html/template"
	"net/http"
)

// page404Handler is handler to return 404 pages
func page404Handler(w http.ResponseWriter, r *http.Request, session *Session) {
	type Context struct {
		User       string
		ID         uint
		UserExists bool
		Language   Language
		SiteName   string
		RootURL    string
	}

	c := Context{}

	c.RootURL = RootURL
	c.SiteName = SiteName
	c.Language = getLanguage(r)
	//user := GetUserFromRequest(r)
	if session != nil {
		user := session.User
		c.User = user.Username
		c.ID = user.ID
	}

	t := template.New("").Funcs(template.FuncMap{
		"Tf": Tf,
	})
	w.WriteHeader(http.StatusNotFound)
	t, err := t.ParseFiles("./templates/uadmin/" + Theme + "/404.html")

	if err != nil {
		fmt.Fprint(w, err.Error())
		fmt.Println("ERROR", err.Error())
	}
	err = t.ExecuteTemplate(w, "404.html", c)
	if err != nil {
		fmt.Println(err.Error())
	}
}
