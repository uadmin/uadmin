package uadmin

import (
	"fmt"
	"html/template"
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

	if r.Method == cPOST {
		var tempSet Setting
		for k, v := range r.Form {
			tempSet = Setting{}
			Get(&tempSet, "code = ?", k)

			if tempSet.ID == 0 {
				continue
			}

			// Process Files and Images
			if tempSet.DataType == tempSet.DataType.File() || tempSet.DataType == tempSet.DataType.Image() {
				continue
			}

			// Process Other Values
			tempSet.ParseFormValue(v)
			tempSet.Save()
		}

		boolSet := []Setting{}
		Filter(&boolSet, "data_type = ?", DataType(0).Boolean())
		for _, s := range boolSet {
			if r.FormValue(s.Code) == "" {
				// Set the value to false
				s.Value = "0"
				s.Save()
			}
		}

		fileSet := []Setting{}
		Filter(&fileSet, "data_type = ? OR data_type = ?", DataType(0).File(), DataType(0).Image())
		for _, tempSet := range fileSet {
			sParts := strings.SplitN(tempSet.Code, ".", 2)

			// Process Files and Images
			_, _, err := r.FormFile(tempSet.Code)
			if err != nil {
				continue
			}

			s, _ := getSchema(tempSet)
			s.FieldByName(sParts[1])

			f := F{Name: tempSet.Code, Type: tMap[tempSet.DataType]}

			val := processUpload(r, &f, "setting", session, &s)
			if val == "" {
				continue
			}
			tempSet.ParseFormValue([]string{val})
			tempSet.Save()
		}
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
