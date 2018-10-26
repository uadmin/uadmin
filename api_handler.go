package uadmin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
)

// apiHandler !
func apiHandler(w http.ResponseWriter, r *http.Request) {
	session := IsAuthenticated(r)
	if session == nil {
		loginHandler(w, r)
		return
	}
	switch r.FormValue("method") {
	case "searchTable":
		modelName := r.FormValue("model")
		q := "%" + r.FormValue("q") + "%"

		delete(r.Form, "method")
		delete(r.Form, "model")
		delete(r.Form, "q")

		model, ok := NewModel(modelName, false)
		s, _ := getSchema(modelName)
		if !ok {
			// error
		}

		arrFilterQuery := []string{}
		query := ""
		arrSearchQuery := []string{}
		for i := range s.Fields {
			f := s.Fields[i]
			if f.Searchable {
				arrSearchQuery = append(arrSearchQuery, fmt.Sprintf("%s LIKE '%s'", gorm.ToDBName(s.Fields[i].Name), q))
			}
		}

		if len(arrSearchQuery) > 0 {
			query = query + "(" + strings.Join(arrSearchQuery, " OR ") + ")"
		}

		for k, v := range r.Form {
			if k == "m" || k == "o" || k == "p" {
				continue
			}
			arrFilterQuery = append(arrFilterQuery, fmt.Sprintf("%s = '%s'", k, v[0]))
		}

		if len(arrFilterQuery) > 0 {
			if query != "" {
				query = query + " AND "
			}
			query = query + strings.Join(arrFilterQuery, " AND ")
		}

		if query != "" {
			r.Form.Set("predefined_query", query)
		}

		ld := getListData(model.Interface(), PageLength, r, session)

		type Context struct {
			List      [][]string `json:"list"`
			PageCount int        `json:"page_count"`
		}
		context := Context{}

		for i := range ld.Rows {
			context.List = append(context.List, []string{})
			for j := range ld.Rows[i] {
				context.List[i] = append(context.List[i], fmt.Sprint(ld.Rows[i][j]))
			}
		}
		context.PageCount = paginationHandler(ld.Count, PageLength)

		bytes, _ := json.Marshal(context)
		w.Write(bytes)
	}
}
