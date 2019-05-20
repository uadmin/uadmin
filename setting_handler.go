package uadmin

import (
	"fmt"
	"html/template"
	"net/http"
)

func settingsHandler(w http.ResponseWriter, r *http.Request, session *Session) {
	type SCat struct {
		ID       uint
		Name     string
		Icon     string
		Settings []Setting
	}
	type Context struct {
		User     string
		SiteName string
		Language Language
		RootURL  string
		SCat     []SCat
	}

	c := Context{}

	c.RootURL = RootURL
	c.Language = getLanguage(r)
	c.SiteName = SiteName
	c.User = session.User.Username
	c.SCat = []SCat{}

	catList := []SettingCategory{}
	All(&catList)

	for _, cat := range catList {
		c.SCat = append(c.SCat, SCat{
			ID:       cat.ID,
			Name:     cat.Name,
			Icon:     cat.Icon,
			Settings: []Setting{},
		})
		Filter(&c.SCat[len(c.SCat)-1].Settings, "category_id = ?", cat.ID)
	}

	t := template.New("").Funcs(template.FuncMap{
		"Tf": Tf,
	})
	//create a new template
	t, err := t.ParseFiles("./templates/uadmin/" + Theme + "/setting.html")

	if err != nil {
		fmt.Fprint(w, err.Error())
		Trail(ERROR, "Unable to render settings page: %s", err)
		return
	}

	err = t.ExecuteTemplate(w, "setting.html", c)
	if err != nil {
		Trail(ERROR, "Unable to execute settings template: %s", err)
	}
}
