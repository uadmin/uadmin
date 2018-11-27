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
		User           string
		Pagination     int
		Data           *listData
		Schema         ModelSchema
		IsUpdated      bool
		Demo           bool
		CanAdd         bool
		CanDelete      bool
		HasAccess      bool
		SiteName       string
		Language       Language
		RootURL        string
		HasCategorical bool
	}

	c := Context{}
	c.RootURL = RootURL
	c.SiteName = SiteName
	c.Language = getLanguage(r)
	c.User = session.User.Username
	user := session.User

	// Creat the template
	t := template.New("").Funcs(template.FuncMap{
		"Tf": Tf,
	})
	t, err := t.ParseFiles("./templates/uadmin/" + Theme + "/list.html")
	if err != nil {
		log.Println("ERROR: listHandler.ParseFiles", err.Error())
	}

	ModelName := r.URL.Path

	up := user.HasAccess(ModelName)
	if user.UserGroupID != 0 {
		Get(&user.UserGroup, "id = ?", user.UserGroupID)
	}
	gp := user.UserGroup.HasAccess(ModelName)
	if up.ID != 0 && !user.Admin {
		if !up.Read {
			page404Handler(w, r, session)
			return
		}
		c.HasAccess = true
		c.CanAdd = up.Add
		c.CanDelete = up.Delete
	} else if gp.ID != 0 && !user.Admin {
		if !gp.Read {
			page404Handler(w, r, session)
			return
		}
		c.HasAccess = true
		c.CanAdd = gp.Add
		c.CanDelete = gp.Delete
	} else if !user.Admin {
		page404Handler(w, r, session)
		return
	}

	if user.Admin {
		c.HasAccess = true
		c.CanAdd = true
		c.CanDelete = true
	}

	if r.Method == cPOST {
		if r.FormValue("delete") == "delete" {
			processDelete(ModelName, w, r, session, &user)
			c.IsUpdated = true
			http.Redirect(w, r, fmt.Sprint(RootURL+r.URL.Path), 303)
		}
	}

	// Initialize the schema
	m, ok := NewModel(ModelName, false)

	// Return 404 if it is an unknown model
	if !ok {
		page404Handler(w, r, session)
		return
	}

	// Get the schema for the model
	c.Schema, _ = getSchema(m.Interface())
	for i := range c.Schema.Fields {
		if c.Schema.Fields[i].CategoricalFilter {
			c.HasCategorical = true
		}
	}
	// func (*ModelSchema, *User) (string, []interface{})
	query := ""
	args := []interface{}{}
	if c.Schema.ListModifier != nil {
		query, args = c.Schema.ListModifier(&c.Schema, &user)
	}

	c.Data = getListData(m.Interface(), PageLength, r, session, query, args...)
	c.Pagination = paginationHandler(c.Data.Count, PageLength)

	err = t.ExecuteTemplate(w, "list.html", c)
	// lock.Unlock()
	if err != nil {
		log.Println("ERROR: listHandler.ExecuteTemplate", err.Error())
	}
}
