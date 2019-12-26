package uadmin

import (
	"net/http"
)

func dAPIAllModelsHandler(w http.ResponseWriter, r *http.Request, session *Session) {
	response := []interface{}{}
	for _, v := range modelList {
		response = append(response, Schema[getModelName(v)])
	}
	ReturnJSON(w, r, map[string]interface{}{
		"status": "ok",
		"result": response,
	})
}
