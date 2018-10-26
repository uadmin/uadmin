package uadmin

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

func getFilter(r *http.Request, session *Session) (interface{}, []interface{}) {
	queryList := []string{}
	args := []interface{}{}
	for k, v := range r.URL.Query() {

		if k == "m" || k == "o" || k == "p" {
			continue
		}

		if k == "q" {
			continue
		}

		user := session.User

		if len(v) > 0 {
			v[0] = strings.Replace(v[0], "{username}", user.Username, -1)
			v[0] = strings.Replace(v[0], "{me}", user.Username, -1)
			v[0] = strings.Replace(v[0], "{userid}", fmt.Sprint(user.ID), -1)
		}

		queryParts := strings.Split(k, "__")
		query := "`" + queryParts[0] + "`"
		if len(queryParts) > 1 {
			if queryParts[1] == "lt" {
				// Less than
				query += " < ?"
			}
			if queryParts[1] == "lte" {
				// Less than or equal to
				query += " <= ?"
			}
			if queryParts[1] == "gt" {
				// Greater than
				query += " > ?"
			}
			if queryParts[1] == "gte" {
				// Greater than or equal to
				query += " >= ?"
			}
			if queryParts[1] == "in" {
				// Greater than or equal to
				query += " IN (?)"
			}
			if queryParts[1] == "contains" {
				// Greater than or equal to
				query += " LIKE ?"
			}
		} else {
			query += " = ?"
		}
		if len(queryParts) > 1 && queryParts[1] == "in" {
			args = append(args, strings.Split(v[0], ","))
		} else if len(queryParts) > 1 && queryParts[1] == "contains" {
			args = append(args, "%"+v[0]+"%")
		} else {
			args = append(args, v[0])
		}
		queryList = append(queryList, query)
	}
	return strings.Join(queryList, " AND "), args
}

// exportHandler !
func exportHandler(w http.ResponseWriter, r *http.Request, session *Session) {
	//http://hostname/admin/export/?m=orders&date__gte=2016-02-01&date__lte=2016-03-01
	modelName := r.URL.Query().Get("m")
	a, ok := NewModelArray(modelName, false)
	if !ok {
		page404Handler(w, r, session)
		return
	}

	query, args := getFilter(r, session)

	if modelName == "warehouseline" {
		whl, ok := NewModel(modelName, true)
		if !ok {
			page404Handler(w, r, session)
			return
		}

		objects := make(map[int]interface{})
		objects[0] = ""
		objects[1] = true
		objects[2] = 0
		objects[3] = -1
		objects[4] = query
		objects[5] = args
		save := whl.MethodByName("AdminPage")
		in := make([]reflect.Value, save.Type().NumIn())

		for i := 0; i < save.Type().NumIn(); i++ {
			object := objects[i]
			in[i] = reflect.ValueOf(object)
		}

		j := save.Call(in)
		a = j[0]
	} else {
		AdminPage("id", true, 0, -1, a.Addr().Interface(), query, args...)
	}

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	m, _ := NewModel(modelName, false)
	t := reflect.TypeOf(m.Interface())

	// Header
	row = sheet.AddRow()
	for i := 0; i < m.NumField(); i++ {
		if t.Field(i).Tag.Get("listExlude") == cTRUE || m.Field(i).Type().Name() == "Model" || m.Field(i).Type().Kind() == reflect.Uint {
			continue
		}
		cell = row.AddCell()
		cell.Value = t.Field(i).Name
	}

	// TODO: Redo this part with schema and data
	now := time.Now()
	var fkModel reflect.Value
	for i := 0; i < a.Len(); i++ {
		row = sheet.AddRow()
		for c := 0; c < m.NumField(); c++ {
			if t.Field(c).Tag.Get("listExlude") == cTRUE || m.Field(c).Type().Name() == "Model" { //|| m.Field(c).Type().Kind() == reflect.Struct {
				continue
			}

			if t.Field(c).Type.Kind() == reflect.Struct {
				fkModel = reflect.New(t.Field(c).Type)
				continue
			}
			cell = row.AddCell()
			if t.Field(c).Type.Kind() == reflect.Float64 {
				cell.Value = fmt.Sprintf("%.2f", a.Index(i).Field(c).Float())
			} else if t.Field(c).Type == reflect.TypeOf(&now) {
				dt, ok := a.Index(i).Field(c).Interface().(*time.Time)
				if ok && dt != nil {
					//cell.Value = dt.Format("2006-01-02 15:04:05")
					cell.SetDateTime(*dt)
					cell.NumFmt = "YYYY-MM-DD HH:MM AM/PM"
				}
			} else if t.Field(c).Type == reflect.TypeOf(now) {
				dt, ok := a.Index(i).Field(c).Interface().(time.Time)
				if ok {
					//cell.Value = dt.Format("2006-01-02 15:04:05")
					cell.SetDateTime(dt)
					cell.NumFmt = "YYYY-MM-DD HH:MM AM/PM"
				}
			} else if t.Field(c).Type.Kind() == reflect.Uint {
				ID := a.Index(i).Field(c).Uint()
				Get(fkModel.Interface(), "`id` = ?", ID)
				cell.Value = fmt.Sprint(fkModel.Interface())
			} else if t.Field(c).Type.Kind() == reflect.Int {
				if t.Field(c).Type == reflect.TypeOf(0) {
					//fmt.Println("INT", t.Field(c).Name)
					cell.Value = fmt.Sprint(a.Index(i).Field(c).Int())
				} else {
					value := a.Index(i).Field(c).Int()
					//fmt.Println("FAKE INT", t.Field(c).Name, value)
					for mIndex := 0; mIndex < t.Field(c).Type.NumMethod(); mIndex++ {
						rValue := a.Index(i).Field(c).Method(mIndex).Call([]reflect.Value{})[0].Int()
						//fmt.Println("FAKE INT", t.Field(c).Name, value, rValue, t.Field(c).Type.Method(mIndex).Name)
						if rValue == value {
							cell.Value = t.Field(c).Type.Method(mIndex).Name
							break
						}
					}
				}
			} else if t.Field(c).Type.Kind() == reflect.Bool {
				cell.Value = fmt.Sprint(a.Index(i).Field(c).Bool())
			} else {
				fmt.Println("STRING", t.Field(c).Name)
				cell.Value = a.Index(i).Field(c).String()
			}

		}
	}
	exportRoot := "./media/export/"
	if _, err = os.Stat(exportRoot); os.IsNotExist(err) {
		os.MkdirAll(exportRoot, 0700)
		os.Create(exportRoot + "index.html")
	}

	fileName := GenerateBase64(24)
	err = file.Save("./media/export/" + fileName + ".xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
	http.Redirect(w, r, "/media/export/"+fileName+".xlsx", 303)
}
