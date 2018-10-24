package uadmin

import (
	"fmt"
	"html"
	"html/template"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/uadmin/uadmin/helper"
)

// GetListSchema returns a schema for list view
func getListData(a interface{}, PageLength int, r *http.Request, session *Session) (l *listData) {
	//r.Form.Set("getInlines", "false")
	l = &listData{}
	schema, _ := getSchema(a)
	language := getLanguage(r)
	translateSchema(&schema, language.Code)

	t := reflect.TypeOf(a)

	// For Order By and Pagination
	o := r.FormValue("o")
	asc := !strings.HasPrefix(o, "-")
	if !asc {
		o = strings.Replace(o, "-", "", 1)
	}
	p := r.FormValue("p")
	page, _ := strconv.ParseInt(p, 10, 32)

	m, ok := newModelArray(schema.ModelName, false)
	if !ok {
		log.Println("ERROR: GetListSchema.newModelArray. No model name", schema.ModelName)
		return
	}

	var _query string

	// Filter records for inlines
	if r.FormValue("inline_id") != "" {
		_query = r.FormValue("inline_id")
	}

	if r.FormValue("inline_id") != "" {
		newModel, _ := newModel(schema.ModelName, false)
		_, ok := t.MethodByName("AdminPage")
		if !ok {
			AdminPage("", asc, int(page-1)*PageLength, PageLength, m.Addr().Interface(), _query)
			l.Count = Count(m.Interface(), _query)
		} else {
			objects := make(map[int]interface{})
			objects[0] = ""
			objects[1] = asc
			objects[2] = int(page-1) * PageLength
			objects[3] = PageLength
			objects[4] = _query
			adminPage := newModel.MethodByName("AdminPage")
			in := make([]reflect.Value, adminPage.Type().NumIn())

			for i := 0; i < adminPage.Type().NumIn(); i++ {
				//l := save.Type().In(i)
				object := objects[i]
				in[i] = reflect.ValueOf(object)
			}

			j := adminPage.Call(in)
			m = j[0]
		}
	} else {
		newModel, _ := newModel(schema.ModelName, false)
		_, ok := t.MethodByName("AdminPage")
		//query, args := getFilter(r)

		var (
			query interface{}
			args  []interface{}
		)

		if r.FormValue("predefined_query") != "" {
			query = r.FormValue("predefined_query")
			args = nil
		} else {
			query, args = getFilter(r, session)
			if r.FormValue("q") != "" {
				q := "%" + r.FormValue("q") + "%"
				imodel := reflect.TypeOf(newModel.Interface())
				var arrSearchQuery []string

				for i := 0; i < imodel.NumField(); i++ {
					f := imodel.Field(i)
					if f.Tag.Get("required") == cTRUE {
						arrSearchQuery = append(arrSearchQuery, fmt.Sprintf("%s LIKE '%s'", gorm.ToDBName(f.Name), q))
					}
				}

				if len(arrSearchQuery) > 0 {
					searchQuery := strings.Join(arrSearchQuery, " OR ")
					if query == "" {
						query = searchQuery
					} else {
						query = query.(string) + " AND (" + searchQuery + ")"
					}
				}
			}
		}

		if !ok {
			AdminPage(o, asc, int(page-1)*PageLength, PageLength, m.Addr().Interface(), query, args...)
			_, HasCount := t.MethodByName("Count")
			if HasCount {
				objects := make(map[int]interface{})
				objects[0] = query
				objects[1] = args

				count := newModel.MethodByName("Count")
				countIn := make([]reflect.Value, count.Type().NumIn())

				for i := 0; i < count.Type().NumIn(); i++ {
					//l := save.Type().In(i)
					object := objects[i]
					countIn[i] = reflect.ValueOf(object)
				}

				count.Call(countIn)
			} else {
				l.Count = Count(m.Interface(), query, args...)
			}
		} else {
			objects := make(map[int]interface{})
			objects[0] = ""
			objects[1] = asc
			objects[2] = int(page-1) * PageLength
			objects[3] = PageLength
			objects[4] = query
			objects[5] = args
			save := newModel.MethodByName("AdminPage")
			in := make([]reflect.Value, save.Type().NumIn())

			for i := 0; i < save.Type().NumIn(); i++ {
				object := objects[i]
				in[i] = reflect.ValueOf(object)
			}

			j := save.Call(in)
			m = j[0]
			_, HasCount := t.MethodByName("Count")
			if HasCount {
				objects = make(map[int]interface{})
				objects[0] = query
				objects[1] = args

				count := newModel.MethodByName("Count")
				countIn := make([]reflect.Value, count.Type().NumIn())

				for i := 0; i < count.Type().NumIn(); i++ {
					object := objects[i]
					countIn[i] = reflect.ValueOf(object)
				}

				count.Call(countIn)
			} else {
				l.Count = Count(m.Interface(), query, args...)
			}
		}
	}

	for index := 0; index < len(schema.Fields); index++ {
		field := reflect.StructField{}
		method := reflect.Method{}
		if schema.Fields[index].IsMethod {
			method, _ = t.MethodByName(schema.Fields[index].Name)
			if strings.Contains(method.Name, "__List") {
				schema.Fields[index].ListDisplay = true
				continue
			} else {
				schema.Fields[index].ListDisplay = false
				continue
			}
		}
		field, _ = t.FieldByName(schema.Fields[index].Name)

		if strings.ToLower(string(field.Name[0])) == string(field.Name[0]) {
			continue
		}
		if !schema.Fields[index].ListDisplay {
			continue
		}

		schema.Fields[index].ListDisplay = true
	}
	for i := 0; i < m.Len(); i++ {
		l.Rows = append(l.Rows, evaluateObject(m.Index(i).Interface(), t, &schema, language.Code))
	}
	return
}

// evaluateObject !
func evaluateObject(obj interface{}, t reflect.Type, s *ModelSchema, lang string) (y []interface{}) {
	value := reflect.ValueOf(obj)
	for index := 0; index < len(s.Fields); index++ {
		if s.Fields[index].IsMethod {
			if strings.Contains(s.Fields[index].Name, "__List") {
				in := []reflect.Value{}
				method := value.MethodByName(s.Fields[index].Name)
				ret := method.Call(in)
				y = append(y, ret[0].String())
			}
			continue
		}

		field, _ := t.FieldByName(s.Fields[index].Name)
		if strings.ToLower(string(field.Name[0])) == string(field.Name[0]) {
			continue
		}
		if !s.Fields[index].ListDisplay {
			continue
		}

		v := value.FieldByName(field.Name)
		if s.Fields[index].Type == cID {
			id, ok := v.Interface().(uint)
			if !ok {
				log.Println("ERROR: EvaluateObject.Interface.(uadmin.Model) ID NOT OK", v.Interface())
			}
			var temp interface{}
			temp = template.HTML(fmt.Sprintf("<a class='clickable Row_id no-style bold' data-id='%d' href='%s%s/%d'>%s</a>", id, RootURL, s.ModelName, id, html.EscapeString(GetString(obj))))
			y = append(y, temp)
		} else if s.Fields[index].Type == cNUMBER {
			temp := v.Interface()
			y = append(y, temp)

		} else if s.Fields[index].Type == cPROGRESSBAR {
			tempProgressValue, _ := strconv.ParseFloat(fmt.Sprint(v.Interface()), 64) // 10
			tempProgressColor := ""
			maxThreshold := 0.0
			var maxColor string
			if len(s.Fields[index].ProgressBar) == 1 {
				for tempThreshold, tempColor := range s.Fields[index].ProgressBar {
					tempProgressColor = tempColor
					maxThreshold = tempThreshold
				}
			} else {
				for tempThreshold, tempColor := range s.Fields[index].ProgressBar {
					if tempThreshold > maxThreshold {
						maxThreshold = tempThreshold
						maxColor = tempColor
					}
				}
				currentThreshold := maxThreshold
				tempProgressColor = maxColor
				for tempThreshold, tempColor := range s.Fields[index].ProgressBar {
					if tempThreshold > tempProgressValue && tempThreshold < currentThreshold {
						tempProgressColor = tempColor
						currentThreshold = tempThreshold
					}
				}
			}
			tempProgressWidth := tempProgressValue / maxThreshold * 100.0
			if tempProgressValue > maxThreshold {
				tempProgressWidth = 100.0
				tempProgressColor = maxColor
			}

			tempColor := helper.GetRGB(tempProgressColor)
			DarkerFactor := 0.67
			tempDarker := []int{
				int(float64(tempColor[0]) * DarkerFactor),
				int(float64(tempColor[1]) * DarkerFactor),
				int(float64(tempColor[2]) * DarkerFactor),
			}

			tempGradient1 := fmt.Sprintf("#%02x%02x%02x", tempColor[0], tempColor[1], tempColor[2])
			tempGradient2 := fmt.Sprintf("#%02x%02x%02x", tempDarker[0], tempDarker[1], tempDarker[2])

			temp := template.HTML(fmt.Sprintf("<div style='border:solid 1px;'><div style='width:%d%%; background-image:linear-gradient(%s,%s); text-align:center;'>%.2f</div></div>", int(tempProgressWidth), tempGradient1, tempGradient2, tempProgressValue))
			y = append(y, temp)
		} else if s.Fields[index].Type == cMONEY {
			temp := commaf(v.Interface())
			y = append(y, temp)
		} else if s.Fields[index].Type == cLINK {
			if fmt.Sprint(v) != "" {
				temp := template.HTML(fmt.Sprintf("<a class='btn btn-primary' href='%s'>%s</a>", v, s.Fields[index].Name))
				y = append(y, temp)
			} else {
				temp := template.HTML(fmt.Sprintf("<span></span>"))
				y = append(y, temp)
			}
		} else if s.Fields[index].Type == cDATE {
			if fmt.Sprint(v.Type())[0] == '*' {
				if v.IsNil() {
					y = append(y, "")
				} else {
					v = v.Elem()
					d, _ := v.Interface().(time.Time)
					y = append(y, d.Format("2006-01-02 15:04:05"))
				}
			} else {
				d, _ := v.Interface().(time.Time)
				y = append(y, d.Format("2006-01-02 15:04:05"))
			}

		} else if s.Fields[index].Type == cFK {
			vID := value.FieldByName(field.Name + "ID")
			cIndex, _ := vID.Interface().(uint)
			fkFieldName := s.Fields[index].Name
			fkFieldName = strings.ToLower(value.FieldByName(fkFieldName).Type().String())
			fkFieldName = strings.Split(fkFieldName, ".")[1]

			// Fetch that record from DB
			fkModel, _ := newModel(fkFieldName, true)
			Get(fkModel.Interface(), "id = ?", cIndex)
			temp := template.HTML(fmt.Sprintf("<a class='clickable no-style bold' href='%s%s/%d'>%s</a>", RootURL, fkFieldName, cIndex, html.EscapeString(GetString(fkModel.Interface()))))
			y = append(y, temp)
		} else if s.Fields[index].Type == cBOOL {
			var temp template.HTML
			tempValue, _ := v.Interface().(bool)
			if tempValue {
				temp = template.HTML(`<i class="fa fa-check-circle" aria-hidden=TRUE style="color:green;"></i>`)
			} else {
				temp = template.HTML(`<i class="fa fa-times-circle" aria-hidden=TRUE style="color:red;"></i>`)
			}
			y = append(y, temp)
		} else if s.Fields[index].Type == cLIST {
			cIndex := v.Int()
			for cCounter := 0; cCounter < len(s.Fields[index].Choices); cCounter++ {
				if uint(cIndex) == s.Fields[index].Choices[cCounter].K {
					y = append(y, s.Fields[index].Choices[uint(cCounter)].V)
					break
				}
			}
		} else if s.Fields[index].Type == cIMAGE {
			temp := template.HTML(fmt.Sprintf(`<img class="hvr-grow pointer image_trigger" style="max-width: 50px; height: auto;" src="%s" />`, v.Interface()))
			y = append(y, temp)
		} else if s.Fields[index].Type == cCODE {
			temp := template.HTML(fmt.Sprintf(`<pre style="width: 200px; white-space: pre-wrap;">%s</pre>`, v.Interface()))
			y = append(y, temp)
		} else if s.Fields[index].Type == cMULTILINGUAL {
			y = append(y, Translate(fmt.Sprint(v), lang, true))
		} else if s.Fields[index].Type == cHTML {
			str := helper.StripTags(fmt.Sprint(v))
			y = append(y, str)
		} else {

			y = append(y, v)
		}

	}
	return
}
