package uadmin

import (
	"encoding/json"
	"net/http"
	"reflect"
)

func dAPIDeleteHandler(w http.ResponseWriter, r *http.Request, s *Session) {
	var rowsCount int64
	modelKV := r.Context().Value(CKey("modelName")).(DApiModelKeyVal)
	modelName := modelKV.CommandName
	model, _ := NewModel(modelName, false)
	schema, _ := GetModelSchema(modelName)
	tableName := schema.TableName
	params := getURLArgs(r)

	// Check CSRF
	if CheckCSRF(r) {
		w.WriteHeader(http.StatusForbidden)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Failed CSRF protection.",
		})
		return
	}

	// Check permission
	allow := false
	if disableDeleter, ok := model.Interface().(APIDisabledDeleter); ok {
		allow = disableDeleter.APIDisabledDelete(r)
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
	if publicDeleter, ok := model.Interface().(APIPublicDeleter); ok {
		allow = publicDeleter.APIPublicDelete(r)
	}
	if !allow && s != nil {
		allow = s.User.GetAccess(modelName).Delete
	}
	if !allow {
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Permission denied",
		})
		return
	}

	// Check if log is required
	log := APILogDelete
	if logDeleter, ok := model.Interface().(APILogDeleter); ok {
		log = logDeleter.APILogDelete(r)
	}

	if modelKV.DataCommand == "" {
		// Delete Multiple
		q, args := getFilters(r, params, tableName, &schema)

		modelArray, _ := NewModelArray(modelName, true)

		// Block Delete All
		if q == "deleted_at IS NULL" {
			ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": "Delete all is blocked",
			})
			return
		}

		if Database.Type == "mysql" {
			db := GetDB()

			if log {
				db.Model(model.Interface()).Where(q, args...).Scan(modelArray.Interface())
			}

			db = db.Where(q, args...).Delete(model.Addr().Interface())
			if db.Error != nil {
				ReturnJSON(w, r, map[string]interface{}{
					"status":  "error",
					"err_msg": "Unable to execute DELETE SQL. " + db.Error.Error(),
				})
				return
			}
			rowsCount = db.RowsAffected
			if log {
				for i := 0; i < modelArray.Elem().Len(); i++ {
					createAPIDeleteLog(modelName, modelArray.Elem().Index(i).Interface(), &s.User, r)
				}
			}

		} else if Database.Type == "sqlite" {
			db := GetDB().Begin()

			if log {
				db.Model(model.Interface()).Where(q, args...).Scan(modelArray.Interface())
			}

			db = db.Exec("PRAGMA case_sensitive_like=ON;")
			db = db.Where(q, args...).Delete(model.Addr().Interface())
			db = db.Exec("PRAGMA case_sensitive_like=OFF;")
			db.Commit()
			if db.Error != nil {
				ReturnJSON(w, r, map[string]interface{}{
					"status":  "error",
					"err_msg": "Unable to COMMIT SQL. " + db.Error.Error(),
				})
				return
			}
			rowsCount = db.RowsAffected
			if log {
				for i := 0; i < modelArray.Elem().Len(); i++ {
					createAPIDeleteLog(modelName, modelArray.Elem().Index(i).Interface(), &s.User, r)
				}
			}
		}
		returnDAPIJSON(w, r, map[string]interface{}{
			"status":     "ok",
			"rows_count": rowsCount,
		}, params, "delete", model.Interface())
	} else if modelKV.DataCommand != "" {
		// Delete One
		m, _ := NewModel(modelName, true)

		db := GetDB()
		if log {
			db.Model(model.Interface()).Where("id = ?", modelKV.DataCommand).Scan(m.Interface())
		}
		db = db.Where("id = ?", modelKV.DataCommand).Delete(model.Addr().Interface())
		if db.Error != nil {
			ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": "Unable to execute DELETE SQL. " + db.Error.Error(),
			})
			return
		}

		if log {
			createAPIDeleteLog(modelName, m.Interface(), &s.User, r)
		}

		returnDAPIJSON(w, r, map[string]interface{}{
			"status":     "ok",
			"rows_count": db.RowsAffected,
		}, params, "delete", model.Interface())
	} else {
		// Error: Unknown format
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "invalid format (" + r.URL.Path + ")",
		})
		return
	}
}

func createAPIDeleteLog(modelName string, m interface{}, user *User, r *http.Request) {
	b, _ := json.Marshal(m)
	output := string(b[:len(b)-1]) + `,"_IP":"` + GetRemoteIP(r) + `"}`

	log := Log{
		Username:  user.Username,
		Action:    Action(0).Deleted(),
		TableName: modelName,
		TableID:   int(GetID(reflect.ValueOf(m))),
		Activity:  output,
	}
	log.Save()
}
