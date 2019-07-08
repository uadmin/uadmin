package uadmin

import (
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"net/http"
	"strings"
	//"github.com/jinzhu/gorm"
)

// apiHandler !
func apiHandler(w http.ResponseWriter, r *http.Request) {
	session := IsAuthenticated(r)
	if session == nil {
		loginHandler(w, r)
		return
	}

	Path := strings.TrimPrefix(r.URL.Path, RootURL+"api")
	if strings.HasPrefix(Path, "/upload_image") {
		UploadImageHandler(w, r, session)
		return
	}
	if strings.HasPrefix(Path, "/search") {
		// Move to separate file
		modelName := r.FormValue("m")
		model, ok := NewModel(modelName, false)
		if !ok {
			pageErrorHandler(w, r, session)
			return
		}
		s, _ := getSchema(modelName)

		query := ""
		args := []interface{}{}
		if s.ListModifier != nil {
			query, args = s.ListModifier(&s, &session.User)
		}

		ld := getListData(model.Interface(), PageLength, r, session, query, args...)

		type Context struct {
			List      [][]string `json:"list"`
			PageCount int        `json:"page_count"`
		}
		context := Context{
			List: [][]string{},
		}

		for i := range ld.Rows {
			context.List = append(context.List, []string{})
			for j := range ld.Rows[i] {
				switch ld.Rows[i][j].(type) {
				case template.HTML:
					context.List[i] = append(context.List[i], fmt.Sprint(ld.Rows[i][j]))
				default:
					context.List[i] = append(context.List[i], html.EscapeString(fmt.Sprint(ld.Rows[i][j])))
				}
			}
		}
		context.PageCount = paginationHandler(ld.Count, PageLength)

		bytes, _ := json.Marshal(context)
		w.Write(bytes)
		return
	}
}
