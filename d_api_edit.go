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
	schema, _ := getSchema(modelName)
	tableName := schema.TableName

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

	// remove empty file and image fields from params to avoid deleteing existing files
	// in case the request does not have a new file
	for k, v := range params {
		for _, f := range schema.Fields {
			if len(k) > 0 && k[0] == '_' && f.ColumnName == k[1:] && (f.Type == cIMAGE || f.Type == cFILE) && v == "" {
				delete(params, k)
				break
			}
		}
	}

	// Process Upload files
	fileList, err := dAPIUpload(w, r, &schema)
	if err != nil {
		Trail(ERROR, "dAPI Add Upload error processing. %s", err)
	}
	for k, v := range fileList {
		params["_"+k] = v
	}

	// Remove the field for file and image after adding it with an underscore
	// to params
	for k := range params {
		for _, f := range schema.Fields {
			if f.ColumnName == k && (f.Type == cIMAGE || f.Type == cFILE) {
				delete(params, k)
				break
			}
		}
	}

	writeMap := getEditMap(params, &schema, &model) // map[string]interface{}

	db := GetDB()

	if len(urlParts) == 2 {
		// Edit multiple
		q, args := getFilters(params, tableName, &schema)

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
		}, params, "edit", model.Interface())
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
		}, params, "edit", model.Interface())
	} else {
		// Error: Unknown format
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "invalid format (" + r.URL.Path + ")",
		})
		return
	}
}

func getEditMap(params map[string]string, schema *ModelSchema, model *reflect.Value) map[string]interface{} {
	paramResult := map[string]interface{}{}

	for k, v := range params {
		if k[0] != '_' {
			continue
		}
		k = k[1:]

		var f *F
		var isPtr = false
		for i := range schema.Fields {
			if k == schema.Fields[i].ColumnName || ((k) == schema.Fields[i].ColumnName+"_id" && schema.Fields[i].Type == cFK) {
				f = &schema.Fields[i]
				isPtr = model.FieldByName(f.Name).Kind() == reflect.Ptr
				break
			}
		}
		if f == nil {
			continue
		}
		if v == "" && isPtr {
			paramResult[k] = nil
		} else {
			paramResult[k] = v
		}
	}

	return paramResult
}

func getWriteQueryFields(v string) string {
	if strings.HasPrefix(v, "_") {
		return "`" + strings.TrimPrefix(v, "_") + "`"
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
