package uadmin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// processDelete is a handler for processing deleting records from a table
func processDelete(a interface{}, w http.ResponseWriter, r *http.Request, session *Session, user *User) {
	if r.FormValue("listID") == "" || r.FormValue("listID") == "," {
		return
	}
	tempID := strings.Split(r.FormValue("listID"), ",")
	var tempIDs []uint
	modelName, ok := a.(string)
	if !ok {
		pageErrorHandler(w, r, session)
		return
	}

	if !user.GetAccess(modelName).Delete {
		return
	}

	for _, v := range tempID {
		temp, _ := strconv.ParseUint(v, 10, 64)
		tempIDs = append(tempIDs, uint(temp))
	}

	if LogDelete {
		for _, v := range tempIDs {
			log := Log{}
			log.Username = user.Username
			log.Action = log.Action.Deleted()
			log.TableName = modelName
			log.TableID = int(v)

			m, ok := NewModel(modelName, false)
			if !ok {
				Trail(ERROR, "processDelete invalid model name: %s", modelName)
			}
			Get(m.Addr().Interface(), "id = ?", v)

			s, _ := getSchema(modelName)
			getFormData(m.Interface(), r, session, &s, user)
			jsonifyValue := map[string]string{}
			for _, ff := range s.Fields {
				jsonifyValue[ff.Name] = fmt.Sprint(ff.Value)
			}

			json, _ := json.Marshal(jsonifyValue)
			log.Activity = string(json)

			log.Save()
		}
	}

	m, ok := NewModel(modelName, true)
	if !ok {
		pageErrorHandler(w, r, session)
		return
	}

	type Deleter interface {
		Delete(interface{}, string, ...interface{})
	}

	deleter, ok := m.Interface().(Deleter)
	if ok {
		deleter.Delete(m.Interface(), "id IN (?)", tempIDs)
	} else {
		DeleteList(m.Interface(), "id IN (?)", tempIDs)
	}
}
