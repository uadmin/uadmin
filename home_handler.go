package uadmin

import (
	"fmt"
	"html/template"
	"net/http"
)

// homeHandler !
func homeHandler(w http.ResponseWriter, r *http.Request, session *Session) {
	type Context struct {
		User     string
		Demo     bool
		Menu     []DashboardMenu
		SiteName string
		Language Language
		RootURL  string
	}

	c := Context{}

	c.RootURL = RootURL
	c.Language = getLanguage(r)
	c.SiteName = SiteName
	c.User = session.User.Username

	c.Menu = session.User.GetDashboardMenu()
	for i := range c.Menu {
		c.Menu[i].MenuName = Translate(c.Menu[i].MenuName, c.Language.Code, true)
	}

	t := template.New("").Funcs(template.FuncMap{
		"Tf": Tf,
	})
	//create a new template
	t, err := t.ParseFiles("./templates/uadmin/" + Theme + "/home.html")

	if err != nil {
		fmt.Fprint(w, err.Error())
		fmt.Println("ERROR", err.Error())
	}

	err = t.ExecuteTemplate(w, "home.html", c)
	if err != nil {
		fmt.Println(err.Error())
	}
}
