package uadmin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	//"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// processForm verifies form data and stores it to DB
func processForm(modelName string, w http.ResponseWriter, r *http.Request, session *Session, s *ModelSchema) reflect.Value { //(errMap map[string]string) {
	var log *Log
	user := session.User
	//errMap = map[string]string{}
	now := time.Now()
	DType := reflect.TypeOf(now)
	DType1 := reflect.TypeOf(&now)
	tempID, _ := strconv.ParseUint(r.FormValue("ID"), 10, 64)
	ID := uint(tempID)
	isNew := ID == 0

	m, ok := NewModel(modelName, true)
	if !ok {
		Trail(ERROR, "processForm.NewModel model not found (%s)", modelName)
		pageErrorHandler(w, r, session)
		return m
	}

	// Fetch record from DB if not new
	if ID != 0 {
		Get(m.Interface(), "id = ?", ID)
	}

	if ID != 0 && LogEdit {
		func() {
			log = &Log{}
			log.ParseRecord(m, modelName, ID, &user, log.Action.Modified(), r)
			log.Save()
		}()
	}

	// Get Type
	t := reflect.TypeOf(m.Interface()).Elem()

	perm := user.GetAccess(modelName)
	appList := []Approval{}

	// Check if there is a field name Createdby
	_, hasCreatedBy := t.FieldByName("CreatedBy")
	_, hasUpdatedBy := t.FieldByName("UpdatedBy")
	_, isValidate := t.MethodByName("Validate")

	// Process Fields
	for index := 0; index < t.NumField(); index++ {
		// Ignore private fields
		if strings.ToLower(string(t.Field(index).Name[0])) == string(t.Field(index).Name[0]) {
			continue
		}
		f := s.FieldByName(t.Field(index).Name)
		if strings.HasSuffix(t.Field(index).Name, "ID") && f.Name == "" {
			f = s.FieldByName(strings.TrimSuffix(t.Field(index).Name, "ID"))
		}
		if f.ReadOnly == "true" || (strings.Contains(f.ReadOnly, "new") && isNew) || (strings.Contains(f.ReadOnly, "edit") && !isNew) {
			continue
		}

		if t.Field(index).Type.Kind() == reflect.Int {
			_v := r.FormValue(t.Field(index).Name)
			i, _ := strconv.ParseInt(_v, 10, 64)

			// Check if approval is required
			if f.Approval && m.Elem().FieldByName(t.Field(index).Name).Int() != i && !perm.Approval {
				appList = append(appList, Approval{
					ModelName:  modelName,
					ColumnName: f.Name,
					OldValue:   fmt.Sprint(m.Elem().FieldByName(t.Field(index).Name).Int()),
					NewValue:   fmt.Sprint(i),
					ChangedBy:  user.Username,
					ChangeDate: now,
				})
			} else {
				m.Elem().FieldByName(t.Field(index).Name).SetInt(i)
			}
		} else if t.Field(index).Type.Kind() == reflect.String {
			// Check if Multi lingual
			val := ""
			if f.Type == cMULTILINGUAL {
				tVal := map[string]string{}
				for _, lang := range activeLangs {
					tVal[lang.Code] = fmt.Sprint(r.FormValue(lang.Code + "-" + t.Field(index).Name))
				}
				buffer := &bytes.Buffer{}
				encoder := json.NewEncoder(buffer)
				encoder.SetEscapeHTML(false)
				_ = encoder.Encode(tVal)
				val = string(buffer.Bytes())
			} else if f.Type == "image" || f.Type == "file" {
				f.Value = m.Elem().FieldByName(f.Name)
				val = processUpload(r, f, modelName, session, s)
				if val == "" {
					continue
				}
			} else {
				val = fmt.Sprint(r.FormValue(t.Field(index).Name))
			}

			// Check if approval is required
			if f.Approval && m.Elem().FieldByName(t.Field(index).Name).String() != val && !perm.Approval {
				// Check if the field is multilingual and if it has a pending approval. If there is a pending
				// approval, add the changes to the existing approval instead of adding a new approval
				newApproval := Approval{}
				if f.Type == cMULTILINGUAL {
					Get(&newApproval, "model_name = ? AND column_name = ? AND model_pk = ? AND approval_action = 0", modelName, f.Name, ID)
					transObj := map[string]string{}
					json.Unmarshal([]byte(newApproval.NewValue), &transObj)

					oldVal := ""
					newVal := ""
					appVal := ""
					changed := false
					for _, lang := range activeLangs {
						oldVal = Translate(m.Elem().FieldByName(t.Field(index).Name).String(), lang.Code, false)
						newVal = Translate(val, lang.Code, false)
						appVal = Translate(newApproval.NewValue, lang.Code, false)
						if oldVal != newVal {
							transObj[lang.Code] = newVal
							changed = true
						} else if newVal != appVal && newApproval.ID != 0 {
							changed = true
							transObj[lang.Code] = appVal
						} else {
							transObj[lang.Code] = newVal
						}
					}
					if changed {
						buf, _ := json.Marshal(transObj)

						newApproval.ModelName = modelName
						newApproval.ColumnName = f.Name
						newApproval.OldValue = m.Elem().FieldByName(t.Field(index).Name).String()
						newApproval.NewValue = string(buf)
						newApproval.ChangedBy = user.Username
						newApproval.ChangeDate = now
						appList = append(appList, newApproval)
					} else {
						m.Elem().FieldByName(t.Field(index).Name).SetString(val)
					}
				} else {
					appList = append(appList, Approval{
						ModelName:  modelName,
						ColumnName: f.Name,
						OldValue:   m.Elem().FieldByName(t.Field(index).Name).String(),
						NewValue:   val,
						ChangedBy:  user.Username,
						ChangeDate: now,
					})
				}
			} else {
				m.Elem().FieldByName(t.Field(index).Name).SetString(val)
			}
		} else if t.Field(index).Type.Kind() == reflect.Bool {
			var val bool
			val = false
			if string(r.FormValue(t.Field(index).Name)) == "on" {
				val = true
			}
			// Check if approval is required
			if f.Approval && m.Elem().FieldByName(t.Field(index).Name).Bool() != val && !perm.Approval {
				appList = append(appList, Approval{
					ModelName:  modelName,
					ColumnName: f.Name,
					OldValue:   fmt.Sprint(m.Elem().FieldByName(t.Field(index).Name).Bool()),
					NewValue:   fmt.Sprint(val),
					ChangedBy:  user.Username,
					ChangeDate: now,
				})
			} else {
				m.Elem().FieldByName(t.Field(index).Name).SetBool(val)
			}
		} else if t.Field(index).Type.Kind() == reflect.Uint {
			_v := r.FormValue(t.Field(index).Name)
			i, _ := strconv.ParseUint(_v, 10, 64)
			val := uint(i)
			// Check if approval is required
			if f.Approval && m.Elem().FieldByName(t.Field(index).Name).Uint() != uint64(val) && !perm.Approval {
				appList = append(appList, Approval{
					ModelName:  modelName,
					ColumnName: f.Name,
					OldValue:   fmt.Sprint(m.Elem().FieldByName(t.Field(index).Name).Uint()),
					NewValue:   fmt.Sprint(val),
					ChangedBy:  user.Username,
					ChangeDate: now,
				})
			} else {
				m.Elem().FieldByName(t.Field(index).Name).Set(reflect.ValueOf(val))
			}
		} else if t.Field(index).Type.Kind() == reflect.Float64 {
			_v := r.FormValue(t.Field(index).Name)
			i, _ := strconv.ParseFloat(_v, 10)

			// Check if approval is required
			if f.Approval && m.Elem().FieldByName(t.Field(index).Name).Float() != i && !perm.Approval {
				appList = append(appList, Approval{
					ModelName:  modelName,
					ColumnName: f.Name,
					OldValue:   fmt.Sprint(m.Elem().FieldByName(t.Field(index).Name).Float()),
					NewValue:   fmt.Sprint(i),
					ChangedBy:  user.Username,
					ChangeDate: now,
				})
			} else {
				m.Elem().FieldByName(t.Field(index).Name).Set(reflect.ValueOf(i))
			}
		} else if t.Field(index).Type.Kind() == reflect.Slice {
			Field := m.Elem().Field(index)
			_v := r.Form[t.Field(index).Name]

			// Initialize the list and item
			m2mListType := reflect.TypeOf(Field.Interface())
			m2mList := reflect.New(m2mListType).Elem()
			m2mItemType := reflect.TypeOf(Field.Interface()).Elem()

			// Append the selected items from the form
			for _, m2mID := range _v {
				m2mItem := reflect.Zero(m2mItemType)
				m2mItem = reflect.New(m2mItemType).Elem()

				Get(m2mItem.Addr().Interface(), "id="+fmt.Sprint(m2mID))
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
			if m.Elem().FieldByName(t.Field(index).Name).Interface().(time.Time).IsZero() {
				tmTemp := time.Time{}
				tmTemp = time.Date(tmTemp.Year(), tmTemp.Month(), tmTemp.Day(), tmTemp.Hour(), tmTemp.Minute(), tmTemp.Second(), tmTemp.Nanosecond(), time.Local)
				m.Elem().FieldByName(t.Field(index).Name).Set(reflect.ValueOf(tmTemp))
			}
			// Check if approval is required
			if f.Approval && !tm.Equal(m.Elem().FieldByName(t.Field(index).Name).Interface().(time.Time)) && !perm.Approval {
				appList = append(appList, Approval{
					ModelName:  modelName,
					ColumnName: f.Name,
					OldValue:   m.Elem().FieldByName(t.Field(index).Name).Interface().(time.Time).Format("2006-01-02 15:04:05-07:00"),
					NewValue:   tm.Format("2006-01-02 15:04:05-07:00"),
					ChangedBy:  user.Username,
					ChangeDate: now,
				})
			} else {
				//if tm.IsZero() {
				//	tm = now
				//}
				m.Elem().FieldByName(t.Field(index).Name).Set(reflect.ValueOf(tm))
			}
		} else if t.Field(index).Type == DType1 {
			if r.FormValue(t.Field(index).Name) == "" {
				if f.Approval && !m.Elem().FieldByName(t.Field(index).Name).IsNil() && !perm.Approval {
					appList = append(appList, Approval{
						ModelName:  modelName,
						ColumnName: f.Name,
						OldValue:   m.Elem().FieldByName(t.Field(index).Name).Elem().Interface().(time.Time).Format("2006-01-02 15:04:05-07:00"),
						NewValue:   "",
						ChangedBy:  user.Username,
						ChangeDate: now,
					})
				} else {
					var tm *time.Time
					m.Elem().FieldByName(t.Field(index).Name).Set(reflect.ValueOf(tm))
				}
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
				tmOld := m.Elem().FieldByName(t.Field(index).Name)
				valChanged := !tmOld.IsNil() && !tm.Equal(tmOld.Elem().Interface().(time.Time))
				valChanged = valChanged || tmOld.IsNil()
				oldVal := ""
				if !tmOld.IsNil() {
					oldVal = tmOld.Elem().Interface().(time.Time).Format("2006-01-02 15:04:05-07:00")
				}
				if f.Approval && valChanged && !perm.Approval {
					appList = append(appList, Approval{
						ModelName:  modelName,
						ColumnName: f.Name,
						OldValue:   oldVal,
						NewValue:   tm.Format("2006-01-02 15:04:05-07:00"),
						ChangedBy:  user.Username,
						ChangeDate: now,
					})
				} else {
					m.Elem().FieldByName(t.Field(index).Name).Set(reflect.ValueOf(&tm))
				}
			}
		} else {
		}
	}

	// Create Log before changing anything
	if !isNew {
		if LogEdit {
			func() {
				log.Save()
			}()
		}
		if hasUpdatedBy {
			m.Elem().FieldByName("UpdatedBy").SetString(user.Username)
		}
	} else {
		if hasCreatedBy {
			m.Elem().FieldByName("CreatedBy").SetString(user.Username)
		}
	}

	if isValidate {
		in := []reflect.Value{}
		validate := m.MethodByName("Validate")
		ret := validate.Call(in)
		if ret[0].Len() > 0 {
			tempErrMap, _ := ret[0].Interface().(map[string]string)

			for k, v := range tempErrMap {
				for i := range s.Fields {
					if s.Fields[i].Name == k {
						s.Fields[i].ErrMsg = v
					}
				}
				//s.FieldByName(k).ErrMsg = v
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
		// ERROR OCCURRED THEN RETURN
		newURL := "new?"
		if !isNew {
			newURL = fmt.Sprintf("%d?", ID)
		}
		var val string
		for i := 0; i < t.NumField(); i++ {
			val = fmt.Sprint(m.Elem().FieldByName(t.Field(i).Name))
			if m.Elem().FieldByName(t.Field(i).Name).Type().String() == "time.Time" {
				val = m.Elem().FieldByName(t.Field(i).Name).Interface().(time.Time).Format("2006-01-02 15:04:05")
			} else if m.Elem().FieldByName(t.Field(i).Name).Type().String() == "*time.Time" {
				if m.Elem().FieldByName(t.Field(i).Name).IsNil() {
					val = ""
				} else {
					val = m.Elem().FieldByName(t.Field(i).Name).Interface().(*time.Time).Format("2006-01-02 15:04:05")
				}
			}
			newURL += t.Field(i).Name + "=" + fmt.Sprint(val) + "&"
		}
		newURL = strings.Replace(newURL, "\n", "", -1)
		r.Form.Set("new_url", newURL[0:len(newURL)-1])
		return m
	}

	// Save the record
	var saverI saver
	saverI, ok = m.Interface().(saver)
	if !ok {
		Save(m.Elem().Addr().Interface())
	} else {
		saverI.Save()
	}

	// Save Approvals
	for _, approval := range appList {
		approval.ModelPK = GetID(m)
		approval.Save()
	}

	// Store the log for a new record
	if LogAdd {
		if isNew {
			ID = GetID(m)
			log = &Log{}
			log.ParseRecord(m, modelName, ID, &user, log.Action.Added(), r)
			log.Save()
		}
	}

	// Redirect the user to the proper URL
	newURL := strings.TrimPrefix(r.URL.Path, RootURL)
	if r.FormValue("save") == "" {
		newURL = RootURL + strings.Split(newURL, "/")[0]
		if r.FormValue("return_url") != "" {
			newURL = r.FormValue("return_url")
		}
		http.Redirect(w, r, newURL, 303)
		return m
	}
	if r.FormValue("save") == "another" {
		newURL = RootURL + strings.Split(newURL, "/")[0] + "/new"
		http.Redirect(w, r, newURL, 303)
		return m
	}
	newURL = RootURL + strings.Split(newURL, "/")[0] + "/" + fmt.Sprint(ID)
	http.Redirect(w, r, newURL, 303)
	return m
}
