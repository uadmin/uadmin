package uadmin

import "net/http"

func trailHandler(w http.ResponseWriter, r *http.Request) {
	session := IsAuthenticated(r)
	type Context struct {
		User     string
		Demo     bool
		Menu     []DashboardMenu
		SiteName string
		Language Language
		RootURL  string
		Logo     string
		FavIcon  string
	}

	c := Context{}

	c.RootURL = RootURL
	c.Language = getLanguage(r)
	c.SiteName = SiteName
	c.User = session.User.Username
	c.Logo = Logo
	c.FavIcon = FavIcon

	RenderHTML(w, r, "./templates/uadmin/default/trail.html", c)
}
