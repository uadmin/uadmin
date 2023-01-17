package uadmin

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
)

func dAPIEditHandler(w http.ResponseWriter, r *http.Request, s *Session) {
	urlParts := strings.Split(r.URL.Path, "/")
	modelName := r.Context().Value(CKey("modelName")).(string)
	model, _ := NewModel(modelName, false)
	schema, _ := getSchema(modelName)
	tableName := schema.TableName

	// Check CSRF
	if CheckCSRF(r) {
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Failed CSRF protection.",
		})
		return
	}

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

	// remove empty file and image fields from params to avoid deleting existing files
	// in case the request does not have a new file
	for k, v := range params {
		for _, f := range schema.Fields {
			if len(k) > 0 && k[0] == '_' && f.ColumnName == k[1:] && (f.Type == cIMAGE || f.Type == cFILE) {
				if v == "" && params[k+"-delete"] != "delete" {
					delete(params, k)
				}
				delete(params, k+"-delete")
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

	writeMap, m2mMap := getEditMap(params, &schema, &model)

	db := GetDB()

	if r.URL.Path == "" {
		// Edit multiple
		q, args := getFilters(r, params, tableName, &schema)

		modelArray, _ := NewModelArray(modelName, true)
		db.Model(model.Interface()).Where(q, args...).Scan(modelArray.Interface())
		db = db.Model(model.Interface()).Where(q, args...).Updates(writeMap)
		if db.Error != nil {
			ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": "Unable to update database. " + db.Error.Error(),
			})
			return
		}
		rowsAffected := db.RowsAffected

		// Process M2M
		db = GetDB().Begin()
		table1 := schema.ModelName
		for i := 0; i < modelArray.Elem().Len(); i++ {
			for k, v := range m2mMap {
				t2Schema, _ := getSchema(k)
				table2 := t2Schema.ModelName
				// First delete existing records
				sql := sqlDialect[Database.Type]["deleteM2M"]
				sql = strings.Replace(sql, "{TABLE1}", table1, -1)
				sql = strings.Replace(sql, "{TABLE2}", table2, -1)
				db = db.Exec(sql, GetID(modelArray.Elem().Index(i)))

				if v == "" {
					continue
				}

				// Now add the records
				for _, id := range strings.Split(v, ",") {
					sql = sqlDialect[Database.Type]["insertM2M"]
					sql = strings.Replace(sql, "{TABLE1}", table1, -1)
					sql = strings.Replace(sql, "{TABLE2}", table2, -1)
					db = db.Exec(sql, GetID(modelArray.Elem().Index(i)), id)
				}
			}
		}
		db.Commit()

		returnDAPIJSON(w, r, map[string]interface{}{
			"status":     "ok",
			"rows_count": rowsAffected,
		}, params, "edit", model.Interface())
		if log {
			for i := 0; i < modelArray.Elem().Len(); i++ {
				createAPIEditLog(modelName, modelArray.Elem().Index(i).Interface(), &s.User, r)
			}
		}

		// Execute business logic
		if _, ok := model.Addr().Interface().(saver); ok {
			for i := 0; i < modelArray.Elem().Len(); i++ {
				id := GetID(modelArray.Elem().Index(i))
				model, _ = NewModel(modelName, false)
				Get(model.Addr().Interface(), "id = ?", id)
				model.Addr().Interface().(saver).Save()
			}
		}
	} else if len(urlParts) == 1 {
		// Edit One
		m, _ := NewModel(modelName, true)
		db.Model(model.Interface()).Where("id = ?", urlParts[0]).Scan(m.Interface())
		db = db.Model(model.Interface()).Where("id = ?", urlParts[0]).Updates(writeMap)
		if db.Error != nil {
			ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": "Unable to update database. " + db.Error.Error(),
			})
			return
		}
		rowsAffected := db.RowsAffected

		// Process M2M
		db = GetDB().Begin()
		table1 := schema.ModelName
		for k, v := range m2mMap {
			t2Schema, _ := getSchema(k)
			table2 := t2Schema.ModelName
			// First delete existing records
			sql := sqlDialect[Database.Type]["deleteM2M"]
			sql = strings.Replace(sql, "{TABLE1}", table1, -1)
			sql = strings.Replace(sql, "{TABLE2}", table2, -1)
			db = db.Exec(sql, urlParts[0])

			if v == "" {
				continue
			}

			// Now add the records
			for _, id := range strings.Split(v, ",") {
				sql = sqlDialect[Database.Type]["insertM2M"]
				sql = strings.Replace(sql, "{TABLE1}", table1, -1)
				sql = strings.Replace(sql, "{TABLE2}", table2, -1)
				db = db.Exec(sql, urlParts[0], id)
			}
		}
		db.Commit()

		if log {
			createAPIEditLog(modelName, m.Interface(), &s.User, r)
		}

		// Execute business logic
		if modelSaver, ok := m.Interface().(saver); ok {
			modelSaver.Save()
		}

		returnDAPIJSON(w, r, map[string]interface{}{
			"status":     "ok",
			"rows_count": rowsAffected,
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

func getEditMap(params map[string]string, schema *ModelSchema, model *reflect.Value) (map[string]interface{}, map[string]string) {
	paramResult := map[string]interface{}{}
	m2mMap := map[string]string{}

	for k, v := range params {
		if k[0] != '_' {
			continue
		}
		k = k[1:]

		// Check M2M
		isM2M := false
		for _, f := range schema.Fields {
			if k == f.ColumnName && f.Type == cM2M {
				m2mMap[strings.ToLower(f.TypeName)] = v
				isM2M = true
				break
			}
		}
		if isM2M {
			continue
		}

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

	return paramResult, m2mMap
}

func getWriteQueryFields(v string) string {
	if strings.HasPrefix(v, "_") {
		return columnEnclosure() + strings.TrimPrefix(v, "_") + columnEnclosure()
	}
	return ""
}

func createAPIEditLog(modelName string, m interface{}, user *User, r *http.Request) {
	b, _ := json.Marshal(m)
	output := string(b[:len(b)-1]) + `,"_IP":"` + GetRemoteIP(r) + `"}`

	log := Log{
		Username:  user.Username,
		Action:    Action(0).Modified(),
		TableName: modelName,
		TableID:   int(GetID(reflect.ValueOf(m))),
		Activity:  output,
	}
	log.Save()
}
