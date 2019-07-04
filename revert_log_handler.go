package uadmin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

// RevertLogHandler !
func revertLogHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request is coming from an authenticated user
	session := IsAuthenticated(r)
	if session == nil {
		pageErrorHandler(w, r, nil)
		return
	}

	// check if the user has permission to Logs
	Preload(session)
	if !session.User.GetAccess("log").Read {
		pageErrorHandler(w, r, nil)
		return
	}

	// retrieve the log
	log := Log{}
	Get(&log, "id = ?", r.FormValue("log_id"))

	if log.ID == 0 {
		pageErrorHandler(w, r, nil)
		return
	}

	// Check if the user has perission to edit the model
	if !session.User.GetAccess(log.TableName).Edit {
		pageErrorHandler(w, r, nil)
		return
	}

	if log.Action == log.Action.Deleted() {
		s := db.NewScope(models[log.TableName])
		tableName := s.TableName()
		sql := fmt.Sprintf("update %s set deleted_at = null where id = %d", tableName, log.TableID)
		db.Exec(sql)
	} else if log.Action == log.Action.Modified() {
		now := time.Now()
		DType := reflect.TypeOf(now)
		var langParser map[string]string
		err := json.Unmarshal([]byte(log.Activity), &langParser)
		if err != nil {
			Trail(ERROR, "revertLogHandler.Unmarash Unable to parse JSON from a log: %s", err.Error())
			return
		}
		model, ok := NewModel(log.TableName, true)
		if !ok {
			Trail(ERROR, "revertLogHandler.NewModel Invalid model name: %s", log.TableName)
		}
		Get(model.Interface(), "id = ?", log.TableID)
		var t reflect.Type
		t = reflect.TypeOf(model.Interface()).Elem()
		for index := 0; index < t.NumField(); index++ {
			if t.Field(index).Type.Kind() == reflect.Int {
				_v := string(langParser[t.Field(index).Name])
				//_v = fmt.Sprintf("%+v", _v)
				i, _ := strconv.ParseInt(_v, 10, 64)

				model.Elem().FieldByName(t.Field(index).Name).SetInt(i)
			} else if t.Field(index).Type.Kind() == reflect.String {
				// Check if Multilingual
				val := ""
				if t.Field(index).Tag.Get("multilingual") == cTRUE {
					tVal := map[string]string{}
					langs := []Language{}
					Filter(&langs, "`active` = ?", true)
					for _, lang := range langs {
						tVal[lang.Code] = fmt.Sprint(langParser[lang.Code+"-"+t.Field(index).Name])
					}
					b, _ := json.Marshal(tVal)
					val = string(b)
				} else {
					val = string(langParser[t.Field(index).Name])
				}

				model.Elem().FieldByName(t.Field(index).Name).SetString(val)
			} else if t.Field(index).Type.Kind() == reflect.Bool {
				var val bool
				val = false
				if string(langParser[t.Field(index).Name]) == "true" {
					val = true
				}
				model.Elem().FieldByName(t.Field(index).Name).SetBool(val)
			} else if t.Field(index).Type.Kind() == reflect.Uint {
				_v := string(langParser[t.Field(index).Name])
				i, _ := strconv.ParseInt(_v, 10, 64)
				val := uint(i)
				model.Elem().FieldByName(t.Field(index).Name).Set(reflect.ValueOf(val))
			} else if t.Field(index).Type.Kind() == reflect.Float64 {
				_v := string(langParser[t.Field(index).Name])
				i, _ := strconv.ParseFloat(_v, 64)
				model.Elem().FieldByName(t.Field(index).Name).Set(reflect.ValueOf(i))
			} else if t.Field(index).Type.Kind() == reflect.Float32 {
				_v := string(langParser[t.Field(index).Name])
				i, _ := strconv.ParseFloat(_v, 32)
				model.Elem().FieldByName(t.Field(index).Name).Set(reflect.ValueOf(float32(i)))
			} else if t.Field(index).Type.Kind() == reflect.Int32 {
				_v := string(langParser[t.Field(index).Name])
				i, _ := strconv.ParseInt(_v, 10, 32)
				model.Elem().FieldByName(t.Field(index).Name).Set(reflect.ValueOf(int32(i)))
			} else if t.Field(index).Type.Kind() == reflect.Int64 {
				_v := string(langParser[t.Field(index).Name])
				i, _ := strconv.ParseInt(_v, 10, 64)
				model.Elem().FieldByName(t.Field(index).Name).Set(reflect.ValueOf(i))
			} else if t.Field(index).Type.Kind() == reflect.Ptr && t.Field(index).Type.Elem() == DType {
				if fmt.Sprint(langParser[t.Field(index).Name]) != "" {
					tm, _ := time.Parse("2006-01-02 15:04:05 -0700", string(langParser[t.Field(index).Name]))
					model.Elem().FieldByName(t.Field(index).Name).Set(reflect.ValueOf(&tm))
				}
			} else if t.Field(index).Type == DType {
				_v := string(langParser[t.Field(index).Name])
				tm, _ := time.Parse("2006-01-02 15:04:05 -0700", _v)
				model.Elem().FieldByName(t.Field(index).Name).Set(reflect.ValueOf(tm))
			} else {
			}
		}

		Save(model.Elem().Addr().Interface())
	}
}
