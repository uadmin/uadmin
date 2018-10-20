package uadmin

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// formHandler !
func formHandler(w http.ResponseWriter, r *http.Request, session *Session) {
	r.ParseMultipartForm(32 << 20)
	type Context struct {
		User            string
		ID              uint
		Schema          ModelSchema
		SaveAndContinue bool
		IsUpdated       bool
		Demo            bool
		CanUpdate       bool
		SiteName        string
		Translation     FormTranslation
		Lang            string
		Direction       string
		RootURL         string
		ReadOnlyF       string
	}

	c := Context{}

	language := getLanguage(r)
	c.RootURL = RootURL
	c.Direction = language.Direction
	c.User = session.User.Username

	c.Translation.AddNew = translateUI(language.Code, "addnew")
	c.Translation.New = translateUI(language.Code, "new")
	c.Translation.SaveAndAddAnother = translateUI(language.Code, "saveandaddanother")
	c.Translation.Save = translateUI(language.Code, "save")
	c.Translation.SaveAndContinue = translateUI(language.Code, "saveandcontinue")
	c.Translation.Dashboard = translateUI(language.Code, "dashboard")
	c.Translation.ChangePassword = translateUI(language.Code, "changepassword")
	c.Translation.Logout = translateUI(language.Code, "logout")
	c.Translation.History = translateUI(language.Code, "history")
	c.Translation.Browse = translateUI(language.Code, "browse")

	c.SiteName = SiteName
	user := session.User

	//temp = template.HTML(fmt.Sprintf("<a class='clickable Row_id no-style bold' data-id='%d' href='/admin/%s/%d'>%s</a>", id, s.ModelName, id, obj))
	t := template.New("").Funcs(template.FuncMap{
		"TOHTML": func(data string) template.HTML {
			str, _ := strconv.Unquote(data)
			return template.HTML(str)
		},
	})
	t, err := t.ParseFiles("./templates/uadmin/" + Theme + "/form.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
		fmt.Println("ERROR", err.Error())
	}

	URLPath := strings.Split(strings.TrimPrefix(r.URL.Path, RootURL), "/")

	ModelName := URLPath[0]
	ModelID, _ := strconv.ParseUint(URLPath[1], 10, 64)
	ID := uint(ModelID)

	m, ok := newModel(ModelName, false)
	if !ok {
		page404Handler(w, r)
		return
	}

	s, _ := getSchema(m.Interface())

	up := user.HasAccess(ModelName)
	if user.UserGroupID != 0 {
		Get(&user.UserGroup, "id = ?", user.UserGroupID)
	}
	gp := user.UserGroup.HasAccess(ModelName)

	// If admin allow adding and editing
	if user.Admin {
		c.CanUpdate = true
	} else {
		// First check if there is a group permission
		if gp.ID != 0 {
			if ID > 0 {
				c.CanUpdate = gp.Edit
			} else {
				c.CanUpdate = gp.Add
			}
		}
		// Then overide it with user permission if it exists
		if up.ID != 0 {
			if ID > 0 {
				c.CanUpdate = gp.Edit
			} else {
				c.CanUpdate = gp.Add
			}
		}
	}

	r.Form.Set("ModelID", fmt.Sprint(ModelID))
	InlineModelName := ""
	if r.FormValue("listModelName") != "" {
		InlineModelName = strings.ToLower(r.FormValue("listModelName"))
	}

	//errMap := map[string]string{}
	if r.Method == cPOST {
		if r.FormValue("delete") == "delete" {
			if InlineModelName != "" {
				processDelete(InlineModelName, w, r, session)
			}
			c.IsUpdated = true
			http.Redirect(w, r, fmt.Sprint(RootURL+r.URL.Path), 303)
		} else {
			// Process the form and check for validaction errors
			processForm(ModelName, w, r, session, s)
			if r.FormValue("new_url") != "" {
				r.URL, _ = url.Parse(r.FormValue("new_url"))
			} else {
				return
			}
		}
	}

	Get(m.Addr().Interface(), "id = ?", ModelID)

	// Return 404 incase the ID doens't exist in the DB and its not in new form
	if URLPath[1] != "new" {
		if getID(m) == 0 {
			page404Handler(w, r)
			return
		}
	}

	c.Schema = getFormData(m.Interface(), r, session, s)

	c.SaveAndContinue = (URLPath[1] == "new" && len(inlines[ModelName]) > 0 && r.URL.Query().Get("return_url") == "")

	// Disable fk for inline form
	if r.URL.Query().Get("return_url") != "" {
		for k := range r.URL.Query() {
			if c.Schema.FieldMyName(k).Type == cFK {
				c.ReadOnlyF = c.Schema.FieldMyName(k).Name
			}
		}
	}

	// for _, f := range c.Schema.Fields {
	// 	c.Data.ErrMsg = append(c.Data.ErrMsg, "")
	// 	for k, v := range errMap {
	// 		if f.Name == k {
	// 			c.Data.ErrMsg[len(c.Data.ErrMsg)-1] = v
	// 		}
	// 	}
	// }

	err = t.ExecuteTemplate(w, "form.html", c)
	if err != nil {
		Trail(ERROR, "Unable to render html template file (form.html). %s", err)
	}

	// Store Read Log in a seperate go routine
	go func() {
		if ModelID > 0 {
			log := &Log{}
			log.ParseRecord(m, m.Type().Name(), uint(ModelID), &session.User, log.Action.Read(), r)
			log.Save()
		}
	}()
}
