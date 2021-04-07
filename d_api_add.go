package uadmin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

func dAPIAddHandler(w http.ResponseWriter, r *http.Request, s *Session) {
	var rowsCount int64
	urlParts := strings.Split(r.URL.Path, "/")
	modelName := urlParts[0]
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
	if disableAdder, ok := model.Interface().(APIDisabledAdder); ok {
		allow = disableAdder.APIDisabledAdd(r)
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
	if publicAdder, ok := model.Interface().(APIPublicAdder); ok {
		allow = publicAdder.APIPublicAdd(r)
	}
	if !allow && s != nil {
		allow = s.User.GetAccess(modelName).Add
	}
	if !allow {
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Permission denied",
		})
		return
	}

	// Check if log is required
	log := APILogAdd
	if logAdder, ok := model.Interface().(APILogAdder); ok {
		log = logAdder.APILogAdd(r)
	}

	// Get parameters
	params := getURLArgs(r)
	params = customParamsAdd(params, model, s)

	createdIDs := []int{}

	// Process Upload files
	fileList, err := dAPIUpload(w, r, &schema)
	if err != nil {
		Trail(ERROR, "dAPI Add Upload error processing. %s", err)
	}
	for k, v := range fileList {
		params["_"+k] = v
	}

	if len(urlParts) == 2 {
		// Add One/Many
		q, args, m2mFields := getAddFilters(params, &schema)

		if DebugDB {
			Trail(DEBUG, "q: %s, v: %#v", q, args)
		}
		db := GetDB().Begin()

		for i := range q {
			// Build args place holder
			argsPlaceHolder := []string{}
			for range args[i] {
				argsPlaceHolder = append(argsPlaceHolder, "?")
			}

			db = db.Exec("INSERT INTO "+tableName+" ("+q[i]+") VALUES ("+strings.Join(argsPlaceHolder, ",")+")", args[i]...)
			rowsCount += db.RowsAffected
		}
		id := []int{}
		if Database.Type == "sqlite" {
			db = db.Raw("SELECT last_insert_rowid() AS lastid")
		} else if Database.Type == "mysql" {
			db = db.Raw("SELECT LAST_INSERT_ID() AS lastid")
		}
		db.Pluck("lastid", &id)
		db.Commit()

		if db.Error != nil {
			ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": "Error in add. " + db.Error.Error(),
			})
			return
		}

		intRowsCount := int(rowsCount)
		for i := 1; i <= intRowsCount; i++ {
			createdIDs = append(createdIDs, id[0]-(intRowsCount-i))
		}

		// Add M2M records
		// No need to delete existing m2m records because it
		// is a new model
		// Insert records
		db = GetDB().Begin()
		for i := range m2mFields {
			table1 := schema.ModelName
			for m2mModelName := range m2mFields[i] {
				t2Schema, _ := getSchema(m2mModelName)
				table2 := t2Schema.ModelName
				for _, id := range strings.Split(m2mFields[i][m2mModelName], ",") {
					if m2mFields[i][m2mModelName] == "" {
						continue
					}
					sql := sqlDialect[Database.Type]["insertM2M"]
					sql = strings.Replace(sql, "{TABLE1}", table1, -1)
					sql = strings.Replace(sql, "{TABLE2}", table2, -1)
					sql = strings.Replace(sql, "{TABLE1_ID}", fmt.Sprint(createdIDs[i]), -1)
					sql = strings.Replace(sql, "{TABLE2_ID}", id, -1)
					db = db.Exec(sql)
				}
			}
		}
		db.Commit()

		returnDAPIJSON(w, r, map[string]interface{}{
			"status":     "ok",
			"rows_count": rowsCount,
			"id":         createdIDs,
		}, params, "add", model.Interface())

		if log {
			for i := range createdIDs {
				createAPIAddLog(q, args, gorm.ToColumnName(model.Type().Name()), createdIDs[i], s, r)
			}
		}
	} else {
		// Error: Unknown format
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "invalid format (" + r.URL.Path + ")",
		})
		return
	}
}

func customParamsAdd(params map[string]string, m reflect.Value, s *Session) map[string]string {
	if m.FieldByName("CreatedAt").Kind() != reflect.Invalid {
		params["_created_at"] = time.Now().Format("2006-01-02 15:04:05")
	}
	if m.FieldByName("CreatedBy").Kind() != reflect.Invalid && s != nil {
		params["_created_by"] = s.User.Username
	}
	return params
}

func getAddFilters(params map[string]string, schema *ModelSchema) (query []string, args [][]interface{}, m2m []map[string]string) {
	query = []string{}
	args = [][]interface{}{}
	m2m = []map[string]string{}

	// Check if we have to add one or multiple
	addOne := true
	for k := range params {
		if k[0] != '_' {
			continue
		}
		if strings.Contains(k[1:], "__") {
			addOne = false
			break
		}
	}

	if addOne {
		// Add one
		itemArgs := []interface{}{}
		itemQ := []string{}
		itemM2M := map[string]string{}
		for k, v := range params {
			if k[0] != '_' {
				continue
			}

			// Process M2M
			fDBName := getWriteQueryFields(k)
			fDBName = fDBName[1 : len(fDBName)-1]
			isM2M := false
			for _, f := range schema.Fields {
				if f.ColumnName == fDBName && f.Type == cM2M {
					itemM2M[strings.ToLower(f.TypeName)] = v
					isM2M = true
					break
				}
			}
			if isM2M {
				continue
			}

			itemQ = append(itemQ, getWriteQueryFields(k))
			itemArgs = append(itemArgs, getAddQueryArg(v))
		}
		query = append(query, strings.Join(itemQ, ", "))
		args = append(args, itemArgs)
		m2m = append(m2m, itemM2M)
	} else {
		// Add Multiple
		index := 0
		var indexExists bool
		var itemArgs []interface{}
		var itemQ []string
		var itemM2M map[string]string
		for {
			indexExists = false
			itemArgs = []interface{}{}
			itemQ = []string{}
			itemM2M = map[string]string{}

			// Check if index exists
			for k := range params {
				if k[0] != '_' {
					continue
				}
				if strings.Contains(k[1:], fmt.Sprintf("__%d", index)) {
					indexExists = true
					break
				}
			}
			if !indexExists {
				break
			}

			// build query and args
			for k, v := range params {
				if k[0] != '_' {
					continue
				}
				queryParts := strings.Split(k[1:], "__")
				paramIndex := 0
				if len(queryParts) == 2 {
					tmp, _ := strconv.ParseInt(queryParts[1], 10, 64)
					paramIndex = int(tmp)
				}
				// if strings.Contains(k[1:], fmt.Sprintf("__%d", index)) {
				if paramIndex == index {
					// Add it
					k = strings.TrimSuffix(k, fmt.Sprintf("__%d", index))

					// Process M2M
					fDBName := getWriteQueryFields(k)
					fDBName = fDBName[1 : len(fDBName)-1]
					isM2M := false
					for _, f := range schema.Fields {
						if f.ColumnName == fDBName && f.Type == cM2M {
							itemM2M[strings.ToLower(f.TypeName)] = v
							isM2M = true
							break
						}
					}
					if isM2M {
						continue
					}

					itemQ = append(itemQ, getWriteQueryFields(k))
					itemArgs = append(itemArgs, getAddQueryArg(v))
				} else if !strings.Contains(k[1:], "__") {
					// Add it

					// Process M2M
					fDBName := getWriteQueryFields(k)
					fDBName = fDBName[1 : len(fDBName)-1]
					isM2M := false
					for _, f := range schema.Fields {
						if f.ColumnName == fDBName && f.Type == cM2M {
							itemM2M[strings.ToLower(f.TypeName)] = v
							isM2M = true
							break
						}
					}
					if isM2M {
						continue
					}

					itemQ = append(itemQ, getWriteQueryFields(k))
					itemArgs = append(itemArgs, getAddQueryArg(v))
				}
			}
			query = append(query, strings.Join(itemQ, ", "))
			args = append(args, itemArgs)
			m2m = append(m2m, itemM2M)

			index++
		}
	}

	return query, args, m2m
}

func getAddQueryArg(v string) interface{} {
	var err error
	v, err = url.QueryUnescape(v)
	if err != nil {
		Trail(WARNING, "getAddQueryArg url.QueryUnescape unable to unescape value. %s", err)
		return []interface{}{v}
	}

	return v
}

func createAPIAddLog(q []string, args [][]interface{}, tableName string, ID int, session *Session, r *http.Request) {
	// TODO: Fix mismatch field name and value assignment
	// in JSON object for Activity field in Logs
	nameMap := map[string]string{}
	for _, f := range Schema[tableName].Fields {
		nameMap[f.ColumnName] = f.Name
	}

	for counter := range q {
		q1 := q[counter]
		args1 := args[counter]
		qParts := strings.Split(q1, ", ")
		vals := map[string]interface{}{
			"_IP": r.RemoteAddr,
		}
		index := 0
		for k, v := range nameMap {
			exists := false
			for i := range qParts {
				if qParts[i] == k {
					exists = true
					break
				}
			}
			if exists {
				vals[v] = args1[index]
				index++
			} else {
				vals[v] = ""
			}
		}
		b, _ := json.Marshal(vals)

		username := ""
		if session != nil {
			username = session.User.Username
		}
		log := Log{
			Username:  username,
			Action:    Action(0).Added(),
			TableName: tableName,
			TableID:   ID,
			Activity:  string(b),
		}
		log.Save()
	}
}
