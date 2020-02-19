package uadmin

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

func dAPIAddHandler(w http.ResponseWriter, r *http.Request, s *Session) {
	var rowsCount int64
	urlParts := strings.Split(r.URL.Path, "/")
	modelName := urlParts[0]
	model, _ := NewModel(modelName, false)
	schema, _ := getSchema(modelName)
	tableName := schema.TableName

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
		// Add One
		q, args := getAddFilters(params)

		if DebugDB {
			Trail(DEBUG, "q: %s, v: %#v", q, args)
		}
		db := GetDB().Begin()

		for i := range q {
			// Build args place holder
			argsPlaceHolder := []string{}
			for _ = range args[i] {
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

func getAddFilters(params map[string]string) (query []string, args [][]interface{}) {
	query = []string{}
	args = [][]interface{}{}

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
		for k, v := range params {
			if k[0] != '_' {
				continue
			}

			itemQ = append(itemQ, getWriteQueryFields(k))
			itemArgs = append(itemArgs, getAddQueryArg(v))
		}
		query = append(query, strings.Join(itemQ, ", "))
		args = append(args, itemArgs)
	} else {
		// Add Multiple
		index := 0
		var indexExists bool
		var itemArgs []interface{}
		var itemQ []string
		for {
			indexExists = false
			itemArgs = []interface{}{}
			itemQ = []string{}

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
				if strings.Contains(k[1:], fmt.Sprintf("__%d", index)) {
					// Add it
					k = strings.TrimSuffix(k, fmt.Sprintf("__%d", index))
					itemQ = append(itemQ, getWriteQueryFields(k))
					itemArgs = append(itemArgs, getAddQueryArg(v))
				} else if !strings.Contains(k[1:], "__") {
					// Add it
					itemQ = append(itemQ, getWriteQueryFields(k))
					itemArgs = append(itemArgs, getAddQueryArg(v))
				}
			}
			query = append(query, strings.Join(itemQ, ", "))
			args = append(args, itemArgs)

			index++
		}
	}

	return query, args
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
