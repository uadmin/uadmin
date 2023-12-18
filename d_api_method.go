package uadmin

import (
	"fmt"
	"net/http"
	"reflect"
)

func dAPIMethodHandler(w http.ResponseWriter, r *http.Request, s *Session) {
	modelKV := r.Context().Value(CKey("modelName")).(DApiModelKeyVal)
	methodName := modelKV.DataCommand
	methodID := modelKV.DataForMethod
	modelName := modelKV.CommandName
	model, _ := NewModel(modelName, true)

	params := getURLArgs(r)

	if modelKV.DataForMethod == "" {
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

	f := model.MethodByName(methodName)
	if !f.IsValid() {
		f = model.Elem().MethodByName(methodName)
	}

	if !f.IsValid() {
		w.WriteHeader(404)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Method (" + methodName + ") doesn't exist.",
		})
		return
	}

	Get(model.Interface(), "id = ?", methodID)
	if GetID(model) == 0 {
		w.WriteHeader(404)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "ID doesn't exist (" + methodID + ").",
		})
		return
	}

	ret := model.MethodByName(methodName).Call([]reflect.Value{})

	// Return if the method has a return value
	if len(ret) != 0 {
		returnDAPIJSON(w, r, map[string]interface{}{
			"status": "ok",
			"value":  fmt.Sprint(ret[0]),
		}, params, "method", model.Interface())
	}
}
