package uadmin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// processForm verifies form data and stores it to DB
func processForm(modelName string, w http.ResponseWriter, r *http.Request, session *Session, s ModelSchema) { //(errMap map[string]string) {
	user := session.User
	//errMap = map[string]string{}
	now := time.Now()
	DType := reflect.TypeOf(now)
	DType1 := reflect.TypeOf(&now)
	tempID, _ := strconv.ParseUint(r.FormValue("ID"), 10, 64)
	ID := uint(tempID)
	isNew := ID == 0

	m, ok := newModel(modelName, true)
	if !ok {
		Trail(ERROR, "processForm.newModel model not found (%s)", modelName)
		page404Handler(w, r)
		return
	}

	// Fetch record from DB if not new
	if ID != 0 {
		Get(m.Interface(), "id = ?", ID)
	}

	// Get Type
	t := reflect.TypeOf(m.Interface()).Elem()

	// Check if there is a field name Createdby
	_, hasCreatedBy := t.FieldByName("CreatedBy")
	_, hasUpdatedBy := t.FieldByName("UpdatedBy")
	_, isValidate := t.MethodByName("Validate")

	// Create Log before changing anything
	if !isNew {
		func() {
			log := &Log{}
			log.ParseRecord(m, modelName, ID, &user, log.Action.Modified(), r)
			log.Save()
		}()
		if hasUpdatedBy {
			m.Elem().FieldByName("UpdatedBy").SetString(user.Username)
		}
	} else {
		if hasCreatedBy {
			m.Elem().FieldByName("CreatedBy").SetString(user.Username)
		}
	}

	// Process Fields
	for index := 0; index < t.NumField(); index++ {
		// Ignore private fields
		if strings.ToLower(string(t.Field(index).Name[0])) == string(t.Field(index).Name[0]) {
			continue
		}
		f := s.FieldMyName(t.Field(index).Name)
		if f.ReadOnly == "true" || (strings.Contains(f.ReadOnly, "new") && isNew) || (strings.Contains(f.ReadOnly, "edit") && !isNew) {
			continue
		}
		if t.Field(index).Type.Kind() == reflect.Int {
			_v := r.FormValue(t.Field(index).Name)
			i, _ := strconv.ParseInt(_v, 10, 64)
			m.Elem().FieldByName(t.Field(index).Name).SetInt(i)
		} else if t.Field(index).Type.Kind() == reflect.String {
			// Check if Multi lingual
			val := ""
			if f.Type == cMULTILINGUAL {
				tVal := map[string]string{}
				langs := []Language{}
				Filter(&langs, "`active` = ?", true)
				for _, lang := range langs {
					tVal[lang.Code] = fmt.Sprint(r.FormValue(lang.Code + "-" + t.Field(index).Name))
				}
				buffer := &bytes.Buffer{}
				encoder := json.NewEncoder(buffer)
				encoder.SetEscapeHTML(false)
				_ = encoder.Encode(tVal)
				val = string(buffer.Bytes())
			} else if f.Type == "image" || f.Type == "file" {
				val, s = processUpload(r, f, modelName, session, s)
				if val == "" {
					continue
				}
			} else {
				val = fmt.Sprint(r.FormValue(t.Field(index).Name))
			}
			m.Elem().FieldByName(t.Field(index).Name).SetString(val)
		} else if t.Field(index).Type.Kind() == reflect.Bool {
			var val bool
			val = false
			if string(r.FormValue(t.Field(index).Name)) == "on" {
				val = true
			}
			m.Elem().FieldByName(t.Field(index).Name).SetBool(val)
		} else if t.Field(index).Type.Kind() == reflect.Uint {
			_v := r.FormValue(t.Field(index).Name)
			i, _ := strconv.ParseInt(_v, 10, 64)
			val := uint(i)
			m.Elem().FieldByName(t.Field(index).Name).Set(reflect.ValueOf(val))
		} else if t.Field(index).Type.Kind() == reflect.Float64 {
			_v := r.FormValue(t.Field(index).Name)
			i, _ := strconv.ParseFloat(_v, 10)
			//val := uint(i)
			m.Elem().FieldByName(t.Field(index).Name).Set(reflect.ValueOf(i))
		} else if t.Field(index).Type.Kind() == reflect.Slice {
			Field := m.Elem().Field(index)
			_v := r.Form[t.Field(index).Name]

			// Initialize the list and item
			m2mListType := reflect.TypeOf(Field.Interface())
			m2mList := reflect.New(m2mListType).Elem()
			m2mItemType := reflect.TypeOf(Field.Interface()).Elem()

			// DeleteM2m(m.Elem().Addr().Interface(), m2mItemType.Name())

			//DeleteM2m(newType.Addr().Interface(), m2mListType.Name())
			// Append the selected items from the form
			for _, m2mID := range _v {
				m2mItem := reflect.Zero(m2mItemType)
				m2mItem = reflect.New(m2mItemType).Elem()
				// params := []reflect.Value{
				// 	reflect.ValueOf("id = ?"),
				// 	reflect.ValueOf(m2mID),
				// }

				Get(m2mItem.Addr().Interface(), "id="+fmt.Sprint(m2mID))
				// m2mItemInstance := m2mItem.MethodByName("Get").Call(params)[0].Elem()
				m2mList = reflect.Append(m2mList, m2mItem)
			}
			// Set the list to the field
			m.Elem().FieldByName(t.Field(index).Name).Set(m2mList)
		} else if t.Field(index).Type == DType {
			if r.FormValue(t.Field(index).Name) == "" {
				continue
			}
			tm, err := time.Parse("2006-01-02 15:04", r.FormValue(t.Field(index).Name))
			if err != nil {
				tm, err = time.Parse("2006-01-02T15:04", r.FormValue(t.Field(index).Name))
			}
			if err != nil {
				tm, err = time.Parse("2006-01-02T15:04:05", r.FormValue(t.Field(index).Name))
			}
			if err != nil {
				tm, err = time.Parse("2006-01-02 15:04:05", r.FormValue(t.Field(index).Name))
			}
			if err != nil {
				Trail(WARNING, "Unable to parse date: %s (%s)", r.FormValue(t.Field(index).Name), err)
				continue
			}
			tm = time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second(), tm.Nanosecond(), time.Local)
			m.Elem().FieldByName(t.Field(index).Name).Set(reflect.ValueOf(tm))
		} else if t.Field(index).Type == DType1 {
			if r.FormValue(t.Field(index).Name) == "" {
				// TODO: Remove value on empty
				//m.Elem().FieldByName(t.Field(index).Name).Set(reflect.ValueOf(nil))
				continue
			} else {
				tm, err := time.Parse("2006-01-02 15:04", r.FormValue(t.Field(index).Name))
				if err != nil {
					tm, err = time.Parse("2006-01-02T15:04", r.FormValue(t.Field(index).Name))
				}
				if err != nil {
					tm, err = time.Parse("2006-01-02T15:04:05", r.FormValue(t.Field(index).Name))
				}
				if err != nil {
					tm, err = time.Parse("2006-01-02 15:04:05", r.FormValue(t.Field(index).Name))
				}
				if err != nil {
					Trail(WARNING, "Unable to parse date: %s (%s)", r.FormValue(t.Field(index).Name), err)
					continue
				}
				tm = time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second(), tm.Nanosecond(), time.Local)
				m.Elem().FieldByName(t.Field(index).Name).Set(reflect.ValueOf(&tm))
			}
		} else {
		}
	}

	if isValidate {
		in := []reflect.Value{}
		validate := m.MethodByName("Validate")
		ret := validate.Call(in)
		if ret[0].Len() > 0 {
			tempErrMap, _ := ret[0].Interface().(map[string]string)

			for k, v := range tempErrMap {
				s.FieldMyName(k).ErrMsg = v
			}
		}
	}

	formError := false
	for _, f := range s.Fields {
		if f.ErrMsg != "" {
			formError = true
			break
		}
	}

	if formError {
		// ERROR OCCURED THEN RETURN
		newURL := "new?"
		for i := 0; i < t.NumField(); i++ {
			newURL += t.Field(i).Name + "=" + fmt.Sprint(m.Elem().FieldByName(t.Field(i).Name)) + "&"
		}
		r.Form.Set("new_url", newURL)
		return
	}

	// Save the record
	var saverI saver
	saverI, ok = m.Interface().(saver)
	if !ok {
		Save(m.Elem().Addr().Interface())
	} else {
		saverI.Save()
	}

	// Store the log for a new record
	if isNew {
		ID = getID(m)
		log := &Log{}
		log.ParseRecord(m, modelName, ID, &user, log.Action.Added(), r)
		log.Save()
	}

	// Redirect the user to the proper URL
	newURL := strings.TrimPrefix(r.URL.Path, RootURL)
	if r.FormValue("save") == "" {
		newURL = RootURL + strings.Split(newURL, "/")[0]
		if r.FormValue("return_url") != "" {
			newURL = r.FormValue("return_url")
		}
		http.Redirect(w, r, newURL, 303)
		return
	}
	if r.FormValue("save") == "another" {
		newURL = RootURL + strings.Split(newURL, "/")[0] + "/new"
		http.Redirect(w, r, newURL, 303)
		return
	}
	newURL = RootURL + strings.Split(newURL, "/")[0] + "/" + fmt.Sprint(ID)
	http.Redirect(w, r, newURL, 303)
	return
}
