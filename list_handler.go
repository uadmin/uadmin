package uadmin

import (
	"fmt"
	"net/http"
	"strings"
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
		CanAdd         bool
		CanDelete      bool
		HasAccess      bool
		SiteName       string
		Language       Language
		RootURL        string
		HasCategorical bool
		Searchable     bool
		CSRF           string
		Logo           string
		FavIcon        string
	}

	c := Context{}
	c.RootURL = RootURL
	c.SiteName = SiteName
	c.Language = getLanguage(r)
	c.User = session.User.Username
	c.CSRF = session.Key
	c.Logo = Logo
	c.FavIcon = FavIcon
	user := session.User

	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
	ModelName := r.URL.Path

	// Check permissions
	perm := user.GetAccess(ModelName)
	if !perm.Read {
		pageErrorHandler(w, r, session)
		return
	}
	c.HasAccess = perm.Read
	c.CanAdd = perm.Add
	c.CanDelete = perm.Delete

	// Initialize the schema
	m, ok := NewModel(ModelName, false)

	// Return 404 if it is an unknown model
	if !ok {
		pageErrorHandler(w, r, session)
		return
	}

	// Process delete
	if r.Method == cPOST {
		if r.FormValue("delete") == "delete" {
			processDelete(ModelName, w, r, session, &user)
			c.IsUpdated = true
			http.Redirect(w, r, fmt.Sprint(RootURL+r.URL.Path), http.StatusSeeOther)
		}
	}

	// Get the schema for the model
	c.Schema, _ = getSchema(m.Interface())
	for i := range c.Schema.Fields {
		if c.Schema.Fields[i].CategoricalFilter {
			c.HasCategorical = true
		}
		if c.Schema.Fields[i].Filter && c.Schema.Fields[i].Type == cFK {
			c.Schema.Fields[i].Choices = getChoices(strings.ToLower(c.Schema.Fields[i].TypeName))
		}
		if c.Schema.Fields[i].Searchable {
			c.Searchable = true
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

	RenderHTML(w, r, "./templates/uadmin/"+c.Schema.GetListTheme()+"/list.html", c)
}
