package uadmin

import (
	"net/http"
	"strings"
)

func settingsHandler(w http.ResponseWriter, r *http.Request, session *Session) {
	r.ParseMultipartForm(32 << 20)
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
	tMap := map[DataType]string{
		DataType(0).File():  cFILE,
		DataType(0).Image(): cIMAGE,
	}

	if session == nil {
		pageErrorHandler(w, r, session)
		return
	}

	// Check if the user has permission to settings models
	perm := session.User.GetAccess("setting")
	if !perm.Read {
		pageErrorHandler(w, r, session)
		return
	}

	settings := []Setting{}
	All(&settings)
	if r.Method == cPOST {
		if !perm.Edit {
			pageErrorHandler(w, r, session)
			return
		}
		var tempSet Setting
		tx := db.Begin()

		for _, s := range settings {
			v, ok := r.Form[s.Code]
			if s.DataType == s.DataType.Image() || s.DataType == s.DataType.Image() {
				sParts := strings.SplitN(s.Code, ".", 2)

				// Process Files and Images
				_, _, err := r.FormFile(s.Code)
				if err != nil {
					continue
				}

				schema, _ := getSchema(s)
				schema.FieldByName(sParts[1])

				f := F{Name: s.Code, Type: tMap[tempSet.DataType]}

				val := processUpload(r, &f, "setting", session, &schema)
				if val == "" {
					continue
				}
				s.ParseFormValue([]string{val})
			} else if s.DataType == s.DataType.Boolean() {
				if ok {
					s.Value = "1"
				} else {
					s.Value = "0"
				}
			} else {
				s.ParseFormValue(v)
			}
			s.ApplyValue()
			tx.Save(&s)
		}
		tx.Commit()
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

	RenderHTML(w, r, "./templates/uadmin/"+Theme+"/setting.html", c)
}
