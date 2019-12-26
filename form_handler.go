package uadmin

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// formHandler handles form view requests to render forms and process POST requests to edit
// the form content. It also handles delete requests for inlines of the form.
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
		Language        Language
		Direction       string
		RootURL         string
		ReadOnlyF       string
	}
	var err error
	c := Context{}

	c.RootURL = RootURL
	c.Language = getLanguage(r)
	c.User = session.User.Username
	c.SiteName = SiteName
	user := session.User

	URLPath := strings.Split(strings.TrimPrefix(r.URL.Path, RootURL), "/")

	ModelName := URLPath[0]
	ModelID, _ := strconv.ParseUint(URLPath[1], 10, 64)
	ID := uint(ModelID)
	_ = ID

	m, ok := NewModel(ModelName, false)
	if !ok {
		pageErrorHandler(w, r, session)
		return
	}

	// Check user permissions
	perm := user.GetAccess(ModelName)
	if !perm.Read {
		pageErrorHandler(w, r, session)
		return
	}
	c.CanUpdate = perm.Add || perm.Edit

	c.Schema, _ = getSchema(m.Interface())

	// Filter inlines that the user does not have permission to
	inlinesList := []*ModelSchema{}
	for i := range c.Schema.Inlines {
		if user.GetAccess(c.Schema.Inlines[i].ModelName).Read {
			inlinesList = append(inlinesList, c.Schema.Inlines[i])
		}
	}
	c.Schema.Inlines = inlinesList

	r.Form.Set("ModelID", fmt.Sprint(ModelID))
	InlineModelName := ""
	if r.FormValue("listModelName") != "" {
		InlineModelName = strings.ToLower(r.FormValue("listModelName"))
	}

	//errMap := map[string]string{}
	if r.Method == cPOST {
		if r.FormValue("delete") == "delete" {
			if InlineModelName != "" {
				processDelete(InlineModelName, w, r, session, &user)
			}
			c.IsUpdated = true
			http.Redirect(w, r, fmt.Sprint(RootURL+r.URL.Path), 303)
		} else {
			// Process the form and check for validation errors
			m = processForm(ModelName, w, r, session, &c.Schema)
			m = m.Elem()
			if r.FormValue("new_url") != "" {
				r.URL, err = url.Parse(r.FormValue("new_url"))
				if err != nil {
					Trail(ERROR, "formHandler unable to parse new_url(%s). %s", r.FormValue("new_url"), err)
					return
				}
			} // else {
			//	return
			//}
		}
	}

	if r.FormValue("new_url") == "" {
		if OptimizeSQLQuery {
			GetForm(m.Addr().Interface(), &c.Schema, "id = ?", ModelID)
		} else {
			Get(m.Addr().Interface(), "id = ?", ModelID)
		}
	}

	// Return 404 incase the ID doens't exist in the DB and its not in new form
	if URLPath[1] != "new" {
		if GetID(m) == 0 {
			pageErrorHandler(w, r, session)
			return
		}
	}

	// Check if Save and Continue
	c.SaveAndContinue = (URLPath[1] == "new" && len(inlines[ModelName]) > 0 && r.URL.Query().Get("return_url") == "")

	// Disable fk for inline form
	if r.URL.Query().Get("return_url") != "" {
		for k := range r.URL.Query() {
			if c.Schema.FieldByName(k).Type == cFK {
				c.ReadOnlyF = c.Schema.FieldByName(k).Name
			}
		}
	}

	// Process User Custom Schema Logic
	if c.Schema.FormModifier != nil {
		c.Schema.FormModifier(&c.Schema, m.Addr().Interface(), &user)
	}

	// Add data to Schema
	getFormData(m.Interface(), r, session, &c.Schema, &user)
	translateSchema(&c.Schema, c.Language.Code)

	RenderHTML(w, r, "./templates/uadmin/"+c.Schema.GetFormTheme()+"/form.html", c)

	// Store Read Log in a separate go routine
	if LogRead {
		go func() {
			if ModelID > 0 {
				log := &Log{}
				log.ParseRecord(m, m.Type().Name(), uint(ModelID), &session.User, log.Action.Read(), r)
				log.Save()
			}
		}()
	}
}
