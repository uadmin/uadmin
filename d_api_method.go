package uadmin

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

func dAPIMethodHandler(w http.ResponseWriter, r *http.Request, s *Session) {
	urlParts := strings.Split(r.URL.Path, "/")
	modelName := r.Context().Value(CKey("modelName")).(string)
	model, _ := NewModel(modelName, true)

	params := getURLArgs(r)

	if len(urlParts) < 2 {
		w.WriteHeader(400)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Bad request, URL format should be api/d/model/method/{METHOD_NAME}/{ID}",
		})
		return
	}

	if CheckCSRF(r) {
		w.WriteHeader(http.StatusForbidden)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Failed CSRF protection.",
		})
		return
	}

	f := model.MethodByName(urlParts[0])
	if !f.IsValid() {
		f = model.Elem().MethodByName(urlParts[0])
	}

	if !f.IsValid() {
		w.WriteHeader(404)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Method (" + urlParts[0] + ") doesn't exist.",
		})
		return
	}

	Get(model.Interface(), "id = ?", urlParts[1])
	if GetID(model) == 0 {
		w.WriteHeader(404)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "ID doesn't exist (" + urlParts[1] + ").",
		})
		return
	}

	ret := model.MethodByName(urlParts[0]).Call([]reflect.Value{})

	// Return if the method has a return value
	if len(ret) != 0 {
		returnDAPIJSON(w, r, map[string]interface{}{
			"status": "ok",
			"value":  fmt.Sprint(ret[0]),
		}, params, "method", model.Interface())
	}
}
