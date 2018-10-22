package uadmin

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// listHandler !
func listHandler(w http.ResponseWriter, r *http.Request, session *Session) {
	r.ParseMultipartForm(32 << 20)

	type Context struct {
		User       string
		Pagination int
		Data       *listData
		Schema     ModelSchema
		IsUpdated  bool
		Demo       bool
		CanAdd     bool
		CanDelete  bool
		HasAccess  bool
		SiteName   string
		Language   Language
		RootURL    string
	}

	c := Context{}
	c.RootURL = RootURL
	c.SiteName = SiteName

	c.Language = getLanguage(r)
	c.User = session.User.Username

	// c.Translation.AddNew = translateUI(language.Code, "addnew")
	// c.Translation.Filter = translateUI(language.Code, "filter")
	// c.Translation.DeleteSelected = translateUI(language.Code, "deleteselected")
	// c.Translation.Excel = translateUI(language.Code, "excel")
	// c.Translation.Dashboard = translateUI(language.Code, "dashboard")
	// c.Translation.ChangePassword = translateUI(language.Code, "changepassword")
	// c.Translation.Logout = translateUI(language.Code, "logout")

	// Check if the user is logged in
	user := session.User

	// Creat the template
	t := template.New("").Funcs(template.FuncMap{
		"Tf": Tf,
	})
	t, err := t.ParseFiles("./templates/uadmin/" + Theme + "/list.html")

	if err != nil {
		log.Println("ERROR: listHandler.ParseFiles", err.Error())
	}

	// Parse the URL
	/*
		URL, _ := url.Parse(r.RequestURI)
		URLPath := strings.Split(URL.Path, "/")

		if len(URLPath) < 3 {
			log.Println("ERROR listHandler.Split: URL TOO SHORT")
		}*/
	ModelName := r.URL.Path // URLPath[2]

	up := user.HasAccess(ModelName)
	if user.UserGroupID != 0 {
		Get(&user.UserGroup, "id = ?", user.UserGroupID)
	}
	gp := user.UserGroup.HasAccess(ModelName)
	if up.ID != 0 && !user.Admin {
		if !up.Read {
			page404Handler(w, r)
			return
		}
		c.HasAccess = true
		c.CanAdd = up.Add
		c.CanDelete = up.Delete
	} else if gp.ID != 0 && !user.Admin {
		if !gp.Read {
			page404Handler(w, r)
			return
		}
		c.HasAccess = true
		c.CanAdd = gp.Add
		c.CanDelete = gp.Delete
	} else if !user.Admin {
		page404Handler(w, r)
		return
	}

	if user.Admin {
		c.HasAccess = true
		c.CanAdd = true
		c.CanDelete = true
	}

	if r.Method == cPOST {
		if r.FormValue("delete") == "delete" {
			processDelete(ModelName, w, r, session)
			c.IsUpdated = true
			http.Redirect(w, r, fmt.Sprint(RootURL+r.URL.Path), 303)
		}
	}

	// Initialize the schema
	m, ok := newModel(ModelName, false)

	// Return 404 if it is an unknown model
	if !ok {
		page404Handler(w, r)
		return
	}

	// Get the schema for the model
	c.Schema, _ = getSchema(m.Interface())
	c.Data = getListData(m.Interface(), PageLength, r, session)
	c.Pagination = paginationHandler(c.Data.Count, PageLength)

	err = t.ExecuteTemplate(w, "list.html", c)
	// lock.Unlock()
	if err != nil {
		log.Println("ERROR: listHandler.ExecuteTemplate", err.Error())
	}
}
