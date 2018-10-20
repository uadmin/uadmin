package uadmin

import (
	"fmt"
	"html/template"
	"net/http"
)

// homeHandler !
func homeHandler(w http.ResponseWriter, r *http.Request, session *Session) {
	type Context struct {
		User        string
		Demo        bool
		Menu        interface{}
		SiteName    string
		Translation DashboardTranslation
		Direction   string
		RootURL     string
	}

	c := Context{}
	language := getLanguage(r)

	c.RootURL = RootURL
	c.Direction = language.Direction

	c.Translation.Dashboard = translateUI(language.Code, "dashboard")
	c.Translation.ChangePassword = translateUI(language.Code, "changepassword")
	c.Translation.Logout = translateUI(language.Code, "logout")

	c.SiteName = SiteName

	c.User = session.User.Username

	menu := session.User.GetDashboardMenu()
	for ctr := range menu {
		menu[ctr].MenuName = translate(menu[ctr].MenuName, language.Code, true)
	}
	c.Menu = menu

	if session.User.Admin {
		menu := []DashboardMenu{}
		All(&menu)
		for ctr := range menu {
			menu[ctr].MenuName = translate(menu[ctr].MenuName, language.Code, true)
		}
		c.Menu = menu
	}

	t := template.New("") //create a new template
	// w.WriteHeader(http.StatusNotFound)
	t, err := t.ParseFiles("./templates/uadmin/" + Theme + "/home.html")

	if err != nil {
		fmt.Fprint(w, err.Error())
		fmt.Println("ERROR", err.Error())
	}
	// URL, _ := url.Parse(r.RequestURI)
	// URLPath := strings.Split(URL.Path, "/")
	err = t.ExecuteTemplate(w, "home.html", c)
	// lock.Unlock()
	if err != nil {
		fmt.Println(err.Error())
	}
}
