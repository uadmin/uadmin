package uadmin

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
)

func dAPIEditHandler(w http.ResponseWriter, r *http.Request, s *Session) {
	urlParts := strings.Split(r.URL.Path, "/")
	modelName := urlParts[0]
	model, _ := NewModel(modelName, false)
	tableName := Schema[modelName].TableName

	// Check permission
	allow := false
	if disableEditor, ok := model.Interface().(APIDisabledEditor); ok {
		allow = disableEditor.APIDisabledEdit(r)
		// This is a "Disable" method
		allow = !allow
		if !allow {
			ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": "Permission denied",
			})
			return
		}
	}
	if publicEditor, ok := model.Interface().(APIPublicEditor); ok {
		allow = publicEditor.APIPublicEdit(r)
	}
	if !allow && s != nil {
		allow = s.User.GetAccess(modelName).Edit
	}
	if !allow {
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Permission denied",
		})
		return
	}

	// Check if log is required
	log := APILogEdit
	if logEditor, ok := model.Interface().(APILogEditor); ok {
		log = logEditor.APILogEdit(r)
	}

	// Get parameters
	params := getURLArgs(r)

	writeMap := getEditMap(params) // map[string]interface{}

	db := GetDB()

	if len(urlParts) == 2 {
		// Edit multiple
		q, args := getFilters(params, tableName)

		modelArray, _ := NewModelArray(modelName, true)
		if log {
			db.Model(model.Interface()).Where(q, args...).Scan(modelArray.Interface())
		}
		db = db.Model(model.Interface()).Where(q, args...).Updates(writeMap)
		if db.Error != nil {
			ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": "Unable to update database. " + db.Error.Error(),
			})
			return
		}
		returnDAPIJSON(w, r, map[string]interface{}{
			"status":     "ok",
			"rows_count": db.RowsAffected,
		}, params)
		if log {
			for i := 0; i < modelArray.Elem().Len(); i++ {
				createAPIEditLog(modelName, modelArray.Elem().Index(i).Interface(), &s.User, r)
			}
		}
	} else if len(urlParts) == 3 {
		// Edit One
		m, _ := NewModel(modelName, true)
		if log {
			db.Model(model.Interface()).Where("id = ?", urlParts[2]).Scan(m.Interface())
		}
		db = db.Model(model.Interface()).Where("id = ?", urlParts[2]).Updates(writeMap)
		if db.Error != nil {
			ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": "Unable to update database. " + db.Error.Error(),
			})
			return
		}

		if log {
			createAPIEditLog(modelName, m.Interface(), &s.User, r)
		}

		returnDAPIJSON(w, r, map[string]interface{}{
			"status":     "ok",
			"rows_count": db.RowsAffected,
		}, params)
	} else {
		// Error: Unknown format
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "invalid format (" + r.URL.Path + ")",
		})
		return
	}
}

func getEditMap(params map[string]string) map[string]interface{} {
	paramResult := map[string]interface{}{}

	for k, v := range params {
		if k[0] != '_' {
			continue
		}
		paramResult[strings.TrimPrefix(k, "_")] = v
	}

	return paramResult
}

func getWriteQueryFields(v string) string {
	if strings.HasPrefix(v, "_") {
		return strings.TrimPrefix(v, "_")
	}
	return ""
}

func createAPIEditLog(modelName string, m interface{}, user *User, r *http.Request) {
	b, _ := json.Marshal(m)
	output := string(b[:len(b)-1]) + `,"_IP":"` + r.RemoteAddr + `"}`

	log := Log{
		Username:  user.Username,
		Action:    Action(0).Modified(),
		TableName: modelName,
		TableID:   int(GetID(reflect.ValueOf(m))),
		Activity:  output,
	}
	log.Save()
}
