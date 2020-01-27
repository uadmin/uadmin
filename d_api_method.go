package uadmin

import (
	"net/http"
	"reflect"
	"strings"
)

func dAPIMethodHandler(w http.ResponseWriter, r *http.Request, s *Session) {
	urlParts := strings.Split(r.URL.Path, "/")
	modelName := urlParts[0]
	model, _ := NewModel(modelName, true)

	if len(urlParts) < 4 {
		w.WriteHeader(400)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Bad request, URL format should be api/d/model/method/{METHOD_NAME}/{ID}",
		})
		return
	}

	f := model.MethodByName(urlParts[2])
	if !f.IsValid() {
		f = model.Elem().MethodByName(urlParts[2])
	}

	if !f.IsValid() {
		w.WriteHeader(404)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Method (" + urlParts[2] + ") doesn't exist.",
		})
		return
	}

	Get(model.Interface(), "id = ?", urlParts[3])
	if GetID(model) == 0 {
		w.WriteHeader(404)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "ID doesn't exist (" + urlParts[3] + ").",
		})
		return
	}

	model.MethodByName(urlParts[2]).Call([]reflect.Value{})
}
