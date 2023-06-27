package uadmin

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
)

func dAPIReadHandler(w http.ResponseWriter, r *http.Request, s *Session) {
	var rowsCount int64
	var err error

	urlParts := strings.Split(r.URL.Path, "/")
	modelName := r.Context().Value(CKey("modelName")).(string)
	model, _ := NewModel(modelName, false)
	params := getURLArgs(r)
	schema, _ := getSchema(modelName)

	// Check permission
	allow := false
	if disableReader, ok := model.Interface().(APIDisabledReader); ok {
		allow = disableReader.APIDisabledRead(r)
		// This is a "Disable" method
		allow = !allow
		if !allow {
			w.WriteHeader(401)
			ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": "Permission denied",
			})
			return
		}
	}
	if publicReader, ok := model.Interface().(APIPublicReader); ok {
		allow = publicReader.APIPublicRead(r)
	}
	if !allow && s != nil {
		allow = s.User.GetAccess(modelName).Read
	}
	if !allow {
		w.WriteHeader(401)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Permission denied",
		})
		return
	}

	// Check if log is required
	log := APILogRead
	if logReader, ok := model.Interface().(APILogReader); ok {
		log = logReader.APILogRead(r)
	}

	if r.URL.Path == "" {
		// Read Multiple
		var m interface{}

		SQL := "SELECT {FIELDS} FROM {TABLE_NAME}"
		if val, ok := params["$distinct"]; ok && val == "1" {
			SQL = "SELECT DISTINCT {FIELDS} FROM {TABLE_NAME}"
		}

		tableName := schema.TableName
		SQL = strings.Replace(SQL, "{TABLE_NAME}", tableName, -1)

		f, customSchema := getQueryFields(r, params, tableName)
		if f != "" {
			SQL = strings.Replace(SQL, "{FIELDS}", f, -1)
		} else {
			SQL = strings.Replace(SQL, "{FIELDS}", tableName+".*", -1)
		}

		join := getQueryJoin(r, params, tableName)
		if join != "" {
			SQL += " " + join
		}

		// Get filters from request
		q, args := getFilters(r, params, tableName, &schema)

		// Apply List Modifier from Schema
		if schema.ListModifier != nil {
			lmQ, lmArgs := schema.ListModifier(&schema, &s.User)

			if lmQ != "" {
				if q != "" {
					q += " AND "
				}

				// Add extra filters from list modifier
				q += lmQ
				args = append(args, lmArgs...)
			}
		}
		if r.Context().Value(CKey("WHERE")) != nil {
			if q != "" {
				q += " AND "
			}
			q += r.Context().Value(CKey("WHERE")).(string)
		}
		if q != "" {
			SQL += " WHERE " + q
		}

		groupBy := getQueryGroupBy(r, params)
		if groupBy != "" {
			SQL += " GROUP BY " + groupBy
		}
		order := getQueryOrder(r, params)
		if order != "" {
			SQL += " ORDER BY " + order
		}
		limit := getQueryLimit(r, params)
		if limit != "" {
			SQL += " LIMIT " + limit
		}
		offset := getQueryOffset(r, params)
		if offset != "" {
			SQL += " OFFSET " + offset
		}

		if DebugDB {
			Trail(DEBUG, SQL)
			Trail(DEBUG, "%#v", args)
		}

		if !customSchema {
			mArray, _ := NewModelArray(modelName, true)
			m = mArray.Interface()
		} // else {
		// 	m = []map[string]interface{}{}
		// }

		if Database.Type == "mysql" {
			db := GetDB()
			if !customSchema {
				err = db.Raw(SQL, args...).Scan(m).Error
			} else {
				var rec []map[string]interface{}
				err = db.Raw(SQL, args...).Scan(&rec).Error
				m = rec
			}
			if a, ok := m.([]map[string]interface{}); ok {
				rowsCount = int64(len(a))
			} else {
				rowsCount = int64(reflect.ValueOf(m).Elem().Len())
			}
		} else if Database.Type == "sqlite" {
			db := GetDB().Begin()
			db.Exec("PRAGMA case_sensitive_like=ON;")
			if !customSchema {
				err = db.Raw(SQL, args...).Scan(m).Error
			} else {
				var rec []map[string]interface{}
				err = db.Raw(SQL, args...).Scan(&rec).Error
				m = rec
			}
			db.Exec("PRAGMA case_sensitive_like=OFF;")
			db.Commit()
			if a, ok := m.([]map[string]interface{}); ok {
				rowsCount = int64(len(a))
			} else {
				rowsCount = int64(reflect.ValueOf(m).Elem().Len())
			}
		} else if Database.Type == "postgres" {
			db := GetDB()
			if !customSchema {
				err = db.Raw(SQL, args...).Scan(m).Error
			} else {
				var rec []map[string]interface{}
				err = db.Raw(SQL, args...).Scan(&rec).Error
				m = rec
			}
			if a, ok := m.([]map[string]interface{}); ok {
				rowsCount = int64(len(a))
			} else {
				rowsCount = int64(reflect.ValueOf(m).Elem().Len())
			}
		}

		// Check for errors
		if err != nil {
			w.WriteHeader(400)
			ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": "Error in read query. " + err.Error(),
			})
			return
		}

		// Preload
		if !customSchema && (params["$preload"] == "1" || params["$preload"] == "true") {
			mList := reflect.ValueOf(m)
			for i := 0; i < mList.Elem().Len(); i++ {
				Preload(mList.Elem().Index(i).Addr().Interface())
			}
		}

		// Process M2M
		getQueryM2M(params, m, customSchema, modelName)

		// Process Full Media URL
		// Mask passwords
		if !customSchema {
			for i := 0; i < reflect.ValueOf(m).Elem().Len(); i++ {
				// Search for media fields
				record := reflect.ValueOf(m).Elem().Index(i)
				for j := range schema.Fields {
					if FullMediaURL && (schema.Fields[j].Type == cIMAGE || schema.Fields[j].Type == cFILE) {
						// Check if there is a file
						if record.FieldByName(schema.Fields[j].Name).String() != "" && record.FieldByName(schema.Fields[j].Name).String()[0] == '/' {
							record.FieldByName(schema.Fields[j].Name).SetString(GetSchema(r) + "://" + GetHostName(r) + record.FieldByName(schema.Fields[j].Name).String())
						}
					}
					if MaskPasswordInAPI && schema.Fields[j].Type == cPASSWORD {
						record.FieldByName(schema.Fields[j].Name).SetString("***")
					}
				}
			}
		}

		returnDAPIJSON(w, r, map[string]interface{}{
			"status": "ok",
			"result": m,
		}, params, "read", model.Interface())
		go func() {
			if log {
				createAPIReadLog(modelName, 0, rowsCount, params, &s.User, r)
			}
		}()
		return
	} else if len(urlParts) == 1 {
		// Read One
		m, _ := NewModel(modelName, true)
		q := "id = ?"
		if r.Context().Value(CKey("WHERE")) != nil {
			q += " AND " + r.Context().Value(CKey("WHERE")).(string)
		}
		Get(m.Interface(), q, urlParts[0])
		rowsCount = 0

		var i interface{}
		if int(GetID(m)) != 0 {
			i = m.Interface()
			rowsCount = 1
		} else {
			w.WriteHeader(404)
		}

		if params["$preload"] == "1" || params["$preload"] == "true" {
			Preload(m.Interface())
		}

		// Process Full Media URL
		// Mask passwords
		// Search for media fields
		record := m.Elem()
		for j := range schema.Fields {
			if FullMediaURL && (schema.Fields[j].Type == cIMAGE || schema.Fields[j].Type == cFILE) {
				// Check if there is a file
				if record.FieldByName(schema.Fields[j].Name).String() != "" && record.FieldByName(schema.Fields[j].Name).String()[0] == '/' {
					record.FieldByName(schema.Fields[j].Name).SetString(GetSchema(r) + "://" + GetHostName(r) + record.FieldByName(schema.Fields[j].Name).String())
				}
			}
			if MaskPasswordInAPI && schema.Fields[j].Type == cPASSWORD {
				record.FieldByName(schema.Fields[j].Name).SetString("***")
			}
		}

		returnDAPIJSON(w, r, map[string]interface{}{
			"status": "ok",
			"result": i,
		}, params, "read", model.Interface())
		go func() {
			if log {
				createAPIReadLog(modelName, int(GetID(m)), rowsCount, map[string]string{"id": urlParts[0]}, &s.User, r)
			}
		}()
	} else {
		// Error: Unknown format
		w.WriteHeader(404)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "invalid format (" + r.URL.Path + ")",
		})
		return
	}
}

func createAPIReadLog(modelName string, ID int, rowsCount int64, params map[string]string, user *User, r *http.Request) {
	vals := map[string]interface{}{
		"params":     params,
		"rows_count": rowsCount,
		"_IP":        GetRemoteIP(r),
	}
	output, _ := json.Marshal(vals)

	log := Log{
		Username:  user.Username,
		Action:    Action(0).Read(),
		TableName: modelName,
		TableID:   ID,
		Activity:  string(output),
	}
	log.Save()

}
