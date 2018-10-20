package uadmin

import (
	"fmt"
	"html/template"
	"net/http"
)

// page404Handler is handler to return 404 pages
func page404Handler(w http.ResponseWriter, r *http.Request) {
	type Context struct {
		User       string
		ID         uint
		UserExists bool
		SiteName   string
		RootURL    string
	}

	c := Context{}

	c.RootURL = RootURL
	c.SiteName = SiteName

	user := GetUserFromRequest(r)
	c.User = user.Username
	c.ID = user.ID

	t := template.New("")
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
