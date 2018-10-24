package uadmin

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func getFormData(a interface{}, r *http.Request, session *Session, s ModelSchema) (schema ModelSchema) {
	//data := []interface{}{}

	t := reflect.TypeOf(a)

	// Get the value of the model
	modelValue := reflect.ValueOf(a)
	// This holds the formatted value of the field
	var value interface{}
	var f *F

	// Get the primary key
	newForm := r.FormValue("ModelID") == "0"
	ModelID64, _ := strconv.ParseUint(r.FormValue("ModelID"), 10, 64)

	// Loop over the fields of the model and get schema information
	for index := 0; index < t.NumField(); index++ {
		//for index, f := range s.F {
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
			DefaultValue = strings.Replace(DefaultValue, "{NOW}", time.Now().Format(time.RFC822), -1)
			fieldValue = reflect.ValueOf(DefaultValue)
		}

		if f.Type == cID {
			m, ok := fieldValue.Interface().(Model)
			if !ok {
				fmt.Println("ID NOT OK")
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
				fkList, _ := newModelArray(strings.ToLower(t.Field(index).Type.Name()), false)
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
						K:        getID(fkList.Index(i)),
						V:        GetString(fkList.Index(i).Interface()),
						Selected: uint(fkValue) == getID(fkList.Index(i)),
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
			// fkModel := reflect.New(fKType).Elem()
			m, ok := newModelArray(strings.ToLower(fKType.Name()), false)

			if !ok {
				log.Println("ERROR: GetListSchema.newModelArray. No model name", s.ModelName)
			}

			All(m.Addr().Interface())
			//	all := fkModel.MethodByName("All")
			//choicesList := all.Call([]reflect.Value{})

			//fmt.Println(t.Name(), fkModel.Type().Name())
			// fmt.Println(all, choicesList[0])
			// myID := uint(0)
			// base, _ := newModel(strings.ToLower(t.Name()), true)
			// m2m, _ := newModelArray(strings.ToLower(fKType.Name()), true)
			// _query := "id = " + fmt.Sprint(ModelID64)
			// Get(base.Interface(), _query)
			// if t.Name() == fkModel.Type().Name() {
			// 	myID = getID(base)
			// }
			f.Choices = []Choice{}
			for i := 0; i < m.Len(); i++ {
				item := m.Index(i).Interface()
				id := getID(m.Index(i))
				// if id == myID {
				// 	continue
				// }
				f.Choices = append(f.Choices, Choice{
					K: id,
					V: GetString(item),
				})
			}

			// newModel := base
			// if t.Field(index).Tag.Get("self_reference") == cTRUE {
			// 	_, ok = t.MethodByName("PreloadM2m")
			// 	if !ok {
			// 		PreloadM2m(base.Elem().Addr().Interface(), m2m.Addr().Interface(), fKType.Name())
			// 	} else {
			// 		save := newModel.MethodByName("PreloadM2m")
			// 		in := make([]reflect.Value, save.Type().NumIn())
			//
			// 		j := save.Call(in)
			// 		m2m = j[0]
			// 	}
			// } else {
			// 	PreloadM2m(base.Interface(), m2m.Interface(), fKType.Name())
			// }

			// m2m := reflect.ValueOf(a).MethodByName("PreloadM2M")
			//
			// preloaded := m2m.Call([]reflect.Value{})[0]
			// // fmt.Println("M2M Func", m2m, a)
			//
			for i := 0; i < fieldValue.Len(); i++ {
				for counter, val := range f.Choices {
					// _ = counter
					//	item := m.Index(i).Interface()
					//fmt.Println(item.Interface())
					itemID := getID(fieldValue.Index(i))
					// fmt.Println(itemID)
					//
					//.Interface().(uint)
					// if !ok {
					// 	fmt.Println("ITEM ID IS NOT OK")
					// }
					if val.K == itemID {
						f.Choices[counter].Selected = true
					}
					// fmt.Println(val)
				}
			}
			// fmt.Println(f.Choices)
			// f.Choices[i]
			// for i := 0; i < len(f.Choices); i++ {
			// 	if f.Choices[i].K == v {
			// 		f.Choices[i].Selected = true
			// 	}
			// }

		} else if f.Type == cDATE {
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
		} else if f.Type == cBOOL {
			d, ok := fieldValue.Interface().(bool)
			if !ok {
				fmt.Println("BOOL NOT OK")
			}
			value = d
		} else if f.Type == cLIST {
			value = fieldValue.Int()
			for i := range f.Choices {
				f.Choices[i].Selected = f.Choices[i].K == uint(fieldValue.Int())
			}
		} else if f.Type == cMULTILINGUAL {
			value = fieldValue.Interface()
			for i := range activeLangs {
				f.Translations[i].Value = translate(fmt.Sprint(value), activeLangs[i].Code, false)
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
			inlineModel, _ := newModel(strings.ToLower(inlineS.ModelName), false)
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
