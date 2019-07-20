package uadmin

import (
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

	RenderHTML(w, r, "./templates/uadmin/"+Theme+"/home.html", c)
}
