package uadmin

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// pageErrorHandler is handler to return 404 pages
func pageErrorHandler(w http.ResponseWriter, r *http.Request, session *Session) {
	type Context struct {
		User       string
		ID         uint
		UserExists bool
		Language   Language
		SiteName   string
		ErrMsg     string
		ErrCode    int
		RootURL    string
	}

	c := Context{}

	c.RootURL = RootURL
	c.SiteName = SiteName
	c.Language = getLanguage(r)
	c.ErrMsg = "Page Not Found"
	c.ErrCode = 404
	if r.Form.Get("err_msg") != "" {
		c.ErrMsg = r.Form.Get("err_msg")
	}
	if code, err := strconv.ParseUint(r.Form.Get("err_code"), 10, 16); err == nil {
		c.ErrCode = int(code)
	}
	if session != nil {
		user := session.User
		c.User = user.Username
		c.ID = user.ID
	}

	t := template.New("").Funcs(template.FuncMap{
		"Tf": Tf,
	})
	w.WriteHeader(c.ErrCode)
	t, err := t.ParseFiles("./templates/uadmin/" + Theme + "/404.html")

	if err != nil {
		fmt.Fprint(w, err.Error())
		Trail(ERROR, "pageErrorHandler unable to parse HTML Page. %s", err)
	}
	err = t.ExecuteTemplate(w, "404.html", c)
	if err != nil {
		Trail(ERROR, "pageErrorHandler unable to execute template. %s", err.Error())
	}
}
