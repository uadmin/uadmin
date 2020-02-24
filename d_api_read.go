package uadmin

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
)

func dAPIReadHandler(w http.ResponseWriter, r *http.Request, s *Session) {
	var err error
	var rowsCount int64

	urlParts := strings.Split(r.URL.Path, "/")
	modelName := urlParts[0]
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

	// Run prequery handler
	if preQueryReader, ok := model.Interface().(APIPreQueryReader); ok {
		if !preQueryReader.APIPreQueryRead(w, r) {
			return
		}
	}

	if len(urlParts) == 2 {
		// Read Multiple
		var m interface{}

		SQL := "SELECT {FIELDS} FROM {TABLE_NAME}"

		tableName := schema.TableName
		SQL = strings.Replace(SQL, "{TABLE_NAME}", tableName, -1)

		f, customSchema := getQueryFields(params, tableName)
		if f != "" {
			SQL = strings.Replace(SQL, "{FIELDS}", f, -1)
		} else {
			SQL = strings.Replace(SQL, "{FIELDS}", "*", -1)
		}

		join := getQueryJoin(params, tableName)
		if join != "" {
			SQL += " " + join
		}

		q, args := getFilters(params, tableName, &schema)
		if q != "" {
			SQL += " WHERE " + q
		}

		groupBy := getQueryGroupBy(params)
		if groupBy != "" {
			SQL += " GROUP BY " + groupBy
		}
		order := getQueryOrder(params)
		if order != "" {
			SQL += " ORDER BY " + order
		}
		limit := getQueryLimit(params)
		if limit != "" {
			SQL += " LIMIT " + limit
		}
		offset := getQueryOffset(params)
		if offset != "" {
			SQL += " OFFSET " + offset
		}

		if DebugDB {
			Trail(DEBUG, SQL)
			Trail(DEBUG, "%#v", args)
		}

		var rows *sql.Rows

		if !customSchema {
			mArray, _ := NewModelArray(modelName, true)
			m = mArray.Interface()
		} else {
			m = []map[string]interface{}{}
		}

		if Database.Type == "mysql" {
			db := GetDB()
			if !customSchema {
				db.Raw(SQL, args...).Scan(m)

				// Preload
				/*
					if params["$preload"] == "1" {
						mList := reflect.ValueOf(m)
						for i := 0; i < mList.Elem().Len(); i++ {
							Preload(mList.Elem().Index(i).Addr().Interface())
						}
					}
				*/
			} else {
				rows, err = db.Raw(SQL, args...).Rows()
				if err != nil {
					w.WriteHeader(500)
					ReturnJSON(w, r, map[string]interface{}{
						"status":  "error",
						"err_msg": "Unable to execute SQL. " + err.Error(),
					})
					return
				}
				m = parseCustomDBSchema(rows)
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
				db.Raw(SQL, args...).Scan(m)
			} else {
				rows, err = db.Raw(SQL, args...).Rows()
				if err != nil {
					w.WriteHeader(500)
					ReturnJSON(w, r, map[string]interface{}{
						"status":  "error",
						"err_msg": "Unable to execute SQL. " + err.Error(),
					})
					return
				}
				m = parseCustomDBSchema(rows)
			}
			db.Exec("PRAGMA case_sensitive_like=OFF;")
			db.Commit()
			if a, ok := m.([]map[string]interface{}); ok {
				rowsCount = int64(len(a))
			} else {
				rowsCount = int64(reflect.ValueOf(m).Elem().Len())
			}
		}
		// Preload
		if params["$preload"] == "1" {
			mList := reflect.ValueOf(m)
			for i := 0; i < mList.Elem().Len(); i++ {
				Preload(mList.Elem().Index(i).Addr().Interface())
			}
		}

		// Process M2M
		getQueryM2M(params, m, customSchema, modelName)

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
	} else if len(urlParts) == 3 {
		// Read One
		m, _ := NewModel(modelName, true)
		Get(m.Interface(), "id = ?", urlParts[2])
		rowsCount = 0

		var i interface{}
		if int(GetID(m)) != 0 {
			i = m.Interface()
			rowsCount = 1
		}

		if params["$preload"] == "1" {
			Preload(m.Interface())
		}

		returnDAPIJSON(w, r, map[string]interface{}{
			"status": "ok",
			"result": i,
		}, params, "read", model.Interface())
		go func() {
			if log {
				createAPIReadLog(modelName, int(GetID(m)), rowsCount, map[string]string{"id": urlParts[2]}, &s.User, r)
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
		"_IP":        r.RemoteAddr,
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
