package uadmin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// processDelete is a handler for processing deleting records from a table
func processDelete(a interface{}, w http.ResponseWriter, r *http.Request, session *Session) {
	if r.FormValue("listID") == "" || r.FormValue("listID") == "," {
		return
	}
	tempID := strings.Split(r.FormValue("listID"), ",")
	var tempIDs []uint
	modelName, ok := a.(string)

	user := GetUserFromRequest(r)
	for _, v := range tempID {
		temp, _ := strconv.ParseUint(v, 10, 32)
		tempIDs = append(tempIDs, uint(temp))

		log := Log{}
		log.Username = user.Username
		log.Action = log.Action.Deleted()
		log.TableName = modelName
		log.TableID = int(temp)

		m, _ := NewModel(modelName, false)
		Get(m.Addr().Interface(), "id = ?", temp)

		s, _ := getSchema(modelName)
		s = getFormData(m.Interface(), r, session, s)
		jsonifyValue := map[string]string{}
		for _, ff := range s.Fields {
			jsonifyValue[ff.Name] = fmt.Sprint(ff.Value)
		}

		json, _ := json.Marshal(jsonifyValue)
		log.Activity = string(json)

		log.Save()
	}

	for _, v := range tempID {
		temp, _ := strconv.ParseUint(v, 10, 32)
		tempIDs = append(tempIDs, uint(temp))
	}

	if !ok {
		page404Handler(w, r, session)
		return
	}
	m, ok := NewModel(modelName, true)

	_, HasCount := reflect.TypeOf(m.Interface()).MethodByName("Delete")
	if HasCount {
		objects := make(map[int]interface{})
		objects[0] = "id IN (?)"
		objects[1] = tempIDs

		count := m.MethodByName("Delete")
		countIn := make([]reflect.Value, count.Type().NumIn())

		for i := 0; i < count.Type().NumIn(); i++ {
			object := objects[i]
			countIn[i] = reflect.ValueOf(object)
		}
		count.Call(countIn)
	} else {
		DeleteList(m.Interface(), "id IN (?)", tempIDs)
	}

	if !ok {
		page404Handler(w, r, session)
		return
	}
}
