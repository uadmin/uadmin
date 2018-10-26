package uadmin

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func getFormData(a interface{}, r *http.Request, session *Session, s ModelSchema) (schema ModelSchema) {
	// This holds the formatted value of the field
	var value interface{}
	var f *F

	// Get the type of model
	t := reflect.TypeOf(a)

	// Get the value of the model
	modelValue := reflect.ValueOf(a)

	// Get the primary key
	newForm := r.FormValue("ModelID") == "0"
	ModelID64, _ := strconv.ParseUint(r.FormValue("ModelID"), 10, 64)

	// Loop over the fields of the model and get schema information
	for index := 0; index < t.NumField(); index++ {
		// Read field value
		fieldValue := modelValue.Field(index)

		// Get the field from schema
		f = &F{}
		fName := t.Field(index).Name
		if t.Field(index).Anonymous {
			fName = "ID"
		}
		for i := range s.Fields {
			if s.Fields[i].Name == fName {
				f = &s.Fields[i]
				break
			}
		}
		if f.Type == cFK {
			fieldValue = modelValue.FieldByName(f.Name + "ID")
		}

		// Check if the field was not found in schema
		if f.Name == "" {
			continue
		}

		// For new records:
		// Overide field value with any values passed in request
		// If not available check if there is a default value for the field
		if r.FormValue(t.Field(index).Name) != "" {
			fieldValue = reflect.ValueOf(r.FormValue(t.Field(index).Name))
		} else if f.DefaultValue != "" && newForm {
			DefaultValue := f.DefaultValue
			DefaultValue = strings.Replace(DefaultValue, "{NOW}", time.Now().Format("2006-01-02 15:04:05"), -1)
			fieldValue = reflect.ValueOf(DefaultValue)
		}

		if f.Type == cID {
			m, ok := fieldValue.Interface().(Model)
			if !ok {
				Trail(ERROR, "Unable tp parse value of ID for %s.%s (%#v)", t.Name(), f.Name, fieldValue.Interface())
			}
			value = m.ID
		} else if f.Type == cNUMBER {
			if f.Format != "" {
				value = fmt.Sprintf(f.Format, value)
			} else {
				value = fieldValue.Interface()
			}
		} else if f.Type == cFK {
			// Get selected items's ID
			fkValue, _ := strconv.ParseUint(fmt.Sprint(fieldValue.Interface()), 10, 64)
			value = fkValue

			if f.LimitChoicesTo == nil {
				fkList, _ := NewModelArray(strings.ToLower(t.Field(index).Type.Name()), false)
				All(fkList.Addr().Interface())

				// Build choices
				f.Choices = []Choice{
					Choice{
						K: 0,
						V: "-",
					},
				}

				for i := 0; i < fkList.Len(); i++ {
					f.Choices = append(f.Choices, Choice{
						K:        GetID(fkList.Index(i)),
						V:        GetString(fkList.Index(i).Interface()),
						Selected: uint(fkValue) == GetID(fkList.Index(i)),
					})
				}
			} else {
				f.Choices = f.LimitChoicesTo(a, &session.User)

				for i := 0; i < len(f.Choices); i++ {
					f.Choices[i].Selected = uint(fkValue) == f.Choices[i].K
				}
			}

		} else if f.Type == cM2M {
			if fmt.Sprint(reflect.TypeOf(fieldValue.Interface())) == "string" {
				continue
			}
			fKType := reflect.TypeOf(fieldValue.Interface()).Elem()
			m, ok := NewModelArray(strings.ToLower(fKType.Name()), false)

			if !ok {
				Trail(ERROR, "GetListSchema.NewModelArray. No model name (%s)", s.ModelName)
			}

			All(m.Addr().Interface())
			f.Choices = []Choice{}
			for i := 0; i < m.Len(); i++ {
				item := m.Index(i).Interface()
				id := GetID(m.Index(i))
				// if id == myID {
				// 	continue
				// }
				f.Choices = append(f.Choices, Choice{
					K: id,
					V: GetString(item),
				})
			}

			for i := 0; i < fieldValue.Len(); i++ {
				for counter, val := range f.Choices {
					itemID := GetID(fieldValue.Index(i))
					if val.K == itemID {
						f.Choices[counter].Selected = true
					}
				}
			}
		} else if f.Type == cDATE {
			if newForm && t.Field(index).Type.Kind() != reflect.Ptr {
				value = time.Now().Format("2006-01-02 15:04:05")
			} else {
				var d *time.Time
				// If the date is not a pointer to date make it a pointer
				if t.Field(index).Type.Kind() != reflect.Ptr {
					tempD, _ := fieldValue.Interface().(time.Time)
					d = &tempD
				} else {
					d, _ = fieldValue.Interface().(*time.Time)
				}
				if d == nil {
					value = ""
				} else {
					value = d.Format("2006-01-02 15:04:05") //2006-01-02 15:04:05
				}
			}
		} else if f.Type == cBOOL {
			d, ok := fieldValue.Interface().(bool)
			if !ok {
				Trail(ERROR, "Unable tp parse bool value for %s.%s (%#v)", t.Name(), f.Name, fieldValue.Interface())
			}
			value = d
		} else if f.Type == cLIST {
			value = fieldValue.Int()
			if f.LimitChoicesTo != nil {
				f.Choices = append([]Choice{Choice{"-", 0, false}}, f.LimitChoicesTo(a, &session.User)...)
			}
			for i := range f.Choices {
				f.Choices[i].Selected = f.Choices[i].K == uint(fieldValue.Int())
			}
		} else if f.Type == cMULTILINGUAL {
			value = fieldValue.Interface()
			for i := range activeLangs {
				f.Translations[i].Value = Translate(fmt.Sprint(value), activeLangs[i].Code, false)
			}
		} else {
			value = fieldValue.Interface()
		}
		f.Value = value
	}

	// Get data from method fields
	for index := 0; index < t.NumMethod(); index++ {
		// Check if the method should be included in the field list
		if strings.Contains(t.Method(index).Name, "__Form") {
			if strings.ToLower(string(t.Method(index).Name[0])) == string(t.Method(index).Name[0]) {
				continue
			}

			in := []reflect.Value{}
			ret := modelValue.Method(index).Call(in)
			s.FieldByName(t.Method(index).Name).Value = ret[0].Interface()
		}
	}

	inlineData := []listData{}
	if uint(ModelID64) != 0 {
		for _, inlineS := range s.Inlines {
			inlineModel, _ := NewModel(strings.ToLower(inlineS.ModelName), false)
			inlineQ := fmt.Sprintf("%s = %d", foreignKeys[s.ModelName][strings.ToLower(inlineS.ModelName)], ModelID64)
			r.Form.Set("inline_id", inlineQ)
			rows := getListData(inlineModel.Interface(), PageLength, r, session)
			r.Form.Del("inline_id")
			if rows.Count == 0 {
				rows.Rows = [][]interface{}{}
			}
			inlineData = append(inlineData, *rows)
		}
	}
	s.InlinesData = inlineData
	s.ModelID = uint(ModelID64)
	return s
}
