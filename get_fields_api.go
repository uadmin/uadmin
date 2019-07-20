package uadmin

import (
	"net/http"
)

func GetFieldsAPI(w http.ResponseWriter, r *http.Request, session *Session) {
	modelName := r.FormValue("m")

	response := []string{}
	s := ModelSchema{}
	for _, v := range Schema {
		if v.ModelName == modelName {
			s = v
			break
		}
	}

	for _, f := range s.Fields {
		response = append(response, f.Name)
	}
	ReturnJSON(w, r, response)
}
