package uadmin

import (
	"fmt"
	"html"
	"html/template"
	"net/http"
	"path"
	"reflect"
	"strconv"
	"strings"
	"time"

	//"github.com/jinzhu/gorm"
	"github.com/uadmin/uadmin/helper"
)

// GetListSchema returns a schema for list view
func getListData(a interface{}, PageLength int, r *http.Request, session *Session, query string, args ...interface{}) (l *listData) {
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

	m, ok := NewModelArray(schema.ModelName, false)
	if !ok {
		Trail(ERROR, "getListSchema.NewModelArray. No model name", schema.ModelName)
		return
	}

	newModel, _ := NewModel(schema.ModelName, false)
	iPager, isPager := newModel.Interface().(adminPager)
	iCounter, isCounter := newModel.Interface().(counter)

	var (
		_query interface{}
		_args  []interface{}
	)

	_query, _args = getFilter(r, session, &schema)
	if _query.(string) != "" {
		if query == "" {
			query = _query.(string)
		} else {
			query += " AND " + _query.(string)
		}
		args = append(args, _args...)
	}
	if !isPager {
		if OptimizeSQLQuery {
			FilterList(&schema, o, asc, int(page-1)*PageLength, PageLength, m.Addr().Interface(), query, args...)
		} else {
			AdminPage(o, asc, int(page-1)*PageLength, PageLength, m.Addr().Interface(), query, args...)
		}
	} else {
		iPager.AdminPage(o, asc, int(page-1)*PageLength, PageLength, m.Addr().Interface(), query, args...)
	}
	if !isCounter {
		l.Count = Count(m.Interface(), query, args...)
	} else {
		l.Count = iCounter.Count(m.Interface(), query, args...)
	}
	for i := 0; i < m.Len(); i++ {
		l.Rows = append(l.Rows, evaluateObject(m.Index(i).Interface(), t, &schema, language.Code, session))
	}
	return
}

// evaluateObject !
func evaluateObject(obj interface{}, t reflect.Type, s *ModelSchema, lang string, session *Session) (y []interface{}) {
	value := reflect.ValueOf(obj)
	for index := 0; index < len(s.Fields); index++ {
		if s.Fields[index].IsMethod {
			if strings.Contains(s.Fields[index].Name, "__List") {
				in := []reflect.Value{}
				method := value.MethodByName(s.Fields[index].Name)
				ret := method.Call(in)
				y = append(y, template.HTML(stripHTMLScriptTag(fmt.Sprint(ret[0].Interface()))))
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
				Trail(ERROR, "evaluateObject.Interface.(uadmin.Model) ID NOT OK. %#v", v.Interface())
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
			fkFieldName := strings.ToLower(value.FieldByName(s.Fields[index].Name).Type().Name())
			if value.FieldByName(s.Fields[index].Name).Type().Kind() == reflect.Ptr {
				fkFieldName = strings.ToLower(value.FieldByName(s.Fields[index].Name).Type().Elem().Name())
			}

			// Fetch that record from DB
			fkModel, _ := NewModel(fkFieldName, true)
			GetStringer(fkModel.Interface(), "id = ?", cIndex)
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
			choiceAdded := false
			if s.Fields[index].LimitChoicesTo != nil {
				s.Fields[index].Choices = s.Fields[index].LimitChoicesTo(obj, &session.User)
			}
			for cCounter := 0; cCounter < len(s.Fields[index].Choices); cCounter++ {
				if uint(cIndex) == s.Fields[index].Choices[cCounter].K {
					y = append(y, s.Fields[index].Choices[uint(cCounter)].V)
					choiceAdded = true
					break
				}
			}
			if !choiceAdded {
				y = append(y, cIndex)
			}
		} else if s.Fields[index].Type == cIMAGE {
			temp := template.HTML(fmt.Sprintf(`<img class="hvr-grow pointer image_trigger" style="max-width: 50px; height: auto;" src="%s" />`, v.Interface()))
			y = append(y, temp)

		} else if s.Fields[index].Type == cFILE {
			if v.Interface() != "" {
				fileLocation := v.Interface().(string)
				fileName := path.Base(fileLocation)
				temp := template.HTML(fmt.Sprintf(`<a href="%s">%s</a>`, v.Interface(), fileName))
				y = append(y, temp)
			} else {
				y = append(y, "")
			}

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
