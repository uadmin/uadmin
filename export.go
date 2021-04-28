package uadmin

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	// import upportd image formats to allow exporting
	// images to excel
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/jinzhu/gorm"
)

func getFilter(r *http.Request, session *Session, schema *ModelSchema) (interface{}, []interface{}) {
	queryList := []string{}
	args := []interface{}{}
	var dateRe = regexp.MustCompile(`^[0-9]{4}-[0-9]{1,2}-[0-9]{1,2}$`)
	for k, v := range r.URL.Query() {

		if k == "m" || k == "o" || k == "p" || k == "return_url" {
			continue
		}

		if len(v) > 0 {
			// Unescape '{' and '}'
			v[0], _ = url.QueryUnescape(v[0])
			// Replace placeholders
			v[0] = strings.Replace(v[0], "{USERNAME}", session.User.Username, -1)
			v[0] = strings.Replace(v[0], "{USERID}", fmt.Sprint(session.User.ID), -1)
			v[0] = strings.Replace(v[0], "{NOW}", time.Now().Format("2006-01-02 15:04:05"), -1)
		}

		if k == "q" {
			// Code for search
			searchQuery := []string{}
			for i := range schema.Fields {
				f := schema.Fields[i]
				if f.Searchable {
					for _, term := range strings.Split(v[0], " ") {
						searchQuery = append(searchQuery, fmt.Sprintf("%s LIKE ?", gorm.ToDBName(schema.Fields[i].Name)))
						args = append(args, "%"+term+"%")
					}
				}
			}

			if len(searchQuery) > 0 {
				queryList = append(queryList, fmt.Sprintf("(%s)", strings.Join(searchQuery, " OR ")))
			}
			continue
		}

		queryParts := strings.Split(k, "__")
		if SQLInjection(r, queryParts[0], "") {
			continue
		}
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
				// IN
				query += " IN (?)"
			}
			if queryParts[1] == "contains" {
				// Contains
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
			// Format dates
			dateType := false
			for i := range schema.Fields {
				if gorm.ToColumnName(schema.Fields[i].Name) == queryParts[0] {
					if schema.Fields[i].Type == cDATE {
						dateType = true
						break
					}
				}
			}
			if dateType && v[0] == "" {
				query = queryParts[0] + " IS NULL"
				//args = append(args, nil)
			} else if dateRe.MatchString(v[0]) {
				d, _ := time.Parse("2006-01-02", v[0])
				d = time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.Local)
				args = append(args, d)
			} else {
				args = append(args, v[0])
			}
		}
		queryList = append(queryList, query)
	}

	return strings.Join(queryList, " AND "), args
}

// exportHandler handles http request for exporting data
func exportHandler(w http.ResponseWriter, r *http.Request, session *Session) {
	//http://hostname/admin/export/?m=orders&date__gte=2016-02-01&date__lte=2016-03-01
	var err error

	// TODO: Call ListSchemaModifier of the schema and use the modified one

	modelName := r.URL.Query().Get("m")
	schema, ok := getSchema(modelName)
	if !ok {
		pageErrorHandler(w, r, session)
		return
	}

	a, ok := NewModelArray(modelName, false)
	if !ok {
		pageErrorHandler(w, r, session)
		return
	}
	m, _ := NewModel(modelName, false)

	query, args := getFilter(r, session, &schema)

	ap, ok := m.Interface().(adminPager)

	if ok {
		err = ap.AdminPage("id", true, 0, -1, a.Addr().Interface(), query, args...)
	} else {
		err = AdminPage("id", true, 0, -1, a.Addr().Interface(), query, args...)
	}

	if err != nil {
		pageErrorHandler(w, r, session)
		return
	}

	file := excelize.NewFile()
	t := reflect.TypeOf(m.Interface())
	sheetName := "Sheet1"

	// Header
	/*
		row = sheet.AddRow()
		headerStyle := xlsx.NewStyle()
		headerStyle.Font.Bold = true
		headerStyle.Font.Size = 10
		headerStyle.Font.Name = "Arial"
		headerStyle.ApplyFont = true
		for i := 0; i < m.NumField(); i++ {
			if !schema.FieldByName(t.Field(i).Name).ListDisplay || m.Field(i).Type().Name() == "Model" || (m.Field(i).Type().Kind() == reflect.Uint && strings.HasSuffix(t.Field(i).Name, "ID")) {
				continue
			}
			cell = row.AddCell()
			cell.SetStyle(headerStyle)
			cell.Value = getDisplayName(t.Field(i).Name)
		}
	*/
	var colName string
	var colIndex = 0
	dateFormat := "yyyy-mm-dd HH:MM:SS"
	var preloaded bool

	headerStyle, _ := file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 7},
			{Type: "top", Color: "000000", Style: 7},
			{Type: "bottom", Color: "000000", Style: 7},
			{Type: "right", Color: "000000", Style: 7},
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"cccccc"},
			Pattern: 1,
		},
		Font: &excelize.Font{Bold: true},
	})
	bodyStyle, _ := file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 7},
			{Type: "top", Color: "000000", Style: 7},
			{Type: "bottom", Color: "000000", Style: 7},
			{Type: "right", Color: "000000", Style: 7},
		},
		Alignment: &excelize.Alignment{WrapText: true, Vertical: "top"},
	})
	dateStyle, _ := file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 7},
			{Type: "top", Color: "000000", Style: 7},
			{Type: "bottom", Color: "000000", Style: 7},
			{Type: "right", Color: "000000", Style: 7},
		},
		CustomNumFmt: &dateFormat,
	})
	codeStyle, _ := file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 7},
			{Type: "top", Color: "000000", Style: 7},
			{Type: "bottom", Color: "000000", Style: 7},
			{Type: "right", Color: "000000", Style: 7},
		},
		Font:      &excelize.Font{Family: "mono"},
		Alignment: &excelize.Alignment{WrapText: true, Vertical: "top"},
	})

	// Add header
	for i := 0; i < m.NumField(); i++ {
		if !schema.FieldByName(t.Field(i).Name).ListDisplay || m.Field(i).Type().Name() == "Model" || (m.Field(i).Type().Kind() == reflect.Uint && strings.HasSuffix(t.Field(i).Name, "ID")) || schema.FieldByName(t.Field(i).Name).Type == cLINK {
			continue
		}
		colIndex++
		colName, _ = excelize.ColumnNumberToName(colIndex)
		file.SetCellValue(sheetName, colName+"1", schema.FieldByName(t.Field(i).Name).DisplayName)
		file.SetCellStyle(sheetName, colName+"1", colName+"1", headerStyle)
		file.SetColWidth(sheetName, colName, colName, 1.3*float64(len(schema.FieldByName(t.Field(i).Name).DisplayName)))
	}

	// Add body data
	for i := 0; i < a.Len(); i++ {
		colIndex = 0
		preloaded = false
		for c := 0; c < m.NumField(); c++ {
			if !schema.FieldByName(t.Field(c).Name).ListDisplay || m.Field(c).Type().Name() == "Model" || (m.Field(c).Type().Kind() == reflect.Uint && strings.HasSuffix(t.Field(c).Name, "ID")) || schema.FieldByName(t.Field(c).Name).Type == cLINK {
				continue
			}
			colIndex++
			colName, _ = excelize.ColumnNumberToName(colIndex)
			cellName := fmt.Sprintf(colName+"%d", i+2)

			// Determine the data type
			if schema.FieldByName(t.Field(c).Name).Type == cDATE {
				// Process Date/Time
				var cDate time.Time
				if t.Field(c).Type.Kind() == reflect.Ptr {
					if a.Index(i).Field(c).IsNil() {
						continue
					}
					cDate = a.Index(i).Field(c).Elem().Interface().(time.Time)
				} else {
					cDate = a.Index(i).Field(c).Interface().(time.Time)
				}
				startDate := time.Date(1900, 1, 1, 0, 0, 0, 0, time.Now().Location())
				file.SetCellValue(sheetName, cellName, cDate.Sub(startDate).Hours()/24)
				file.SetColWidth(sheetName, colName, colName, 20)

				file.SetCellStyle(sheetName, cellName, cellName, dateStyle)
				cWidth, _ := file.GetColWidth(sheetName, colName)
				if cWidth < 20 {
					file.SetColWidth(sheetName, colName, colName, 20)
				}
			} else if t.Field(c).Type.Kind() == reflect.Struct || (t.Field(c).Type.Kind() == reflect.Ptr && t.Field(c).Type.Elem().Kind() == reflect.Struct) {
				// Process forign keys
				if !preloaded {
					Preload(a.Index(i).Addr().Interface())
				}
				file.SetCellValue(sheetName, cellName, GetString(a.Index(i).Field(c).Interface()))
				file.SetCellStyle(sheetName, cellName, cellName, bodyStyle)
				excelAdjustWidthHight(file, sheetName, colName, cellName, i+2, GetString(a.Index(i).Field(c).Interface()))
			} else if t.Field(c).Type.Kind() == reflect.Int && t.Field(c).Type != reflect.TypeOf(0) {
				// Process static list type
				value := a.Index(i).Field(c).Interface()
				file.SetCellValue(sheetName, cellName, GetString(value))
				file.SetCellStyle(sheetName, cellName, cellName, bodyStyle)
				excelAdjustWidthHight(file, sheetName, colName, cellName, i+2, GetString(value))
			} else if schema.FieldByName(t.Field(c).Name).Type == cIMAGE {
				// Process images
				if a.Index(i).Field(c).String() == "" {
					continue
				}
				file.SetRowHeight(sheetName, i+2, 100)
				file.SetColWidth(sheetName, colName, colName, 25)
				file.AddPicture(sheetName, cellName, a.Index(i).Field(c).String()[1:], `{"autofit": true, "print_obj": true, "lock_aspect_ratio": true, "locked": false, "positioning": "oneCell", "x_scale":5.0, "y_scale":5.0}`)
				file.SetCellStyle(sheetName, cellName, cellName, bodyStyle)
			} else if schema.FieldByName(t.Field(c).Name).Type == cCODE {
				file.SetCellValue(sheetName, cellName, a.Index(i).Field(c).Interface())
				file.SetCellStyle(sheetName, cellName, cellName, codeStyle)
				excelAdjustWidthHight(file, sheetName, colName, cellName, i+2, fmt.Sprint(a.Index(i).Field(c).Interface()))
			} else {
				// All other data
				file.SetCellValue(sheetName, cellName, a.Index(i).Field(c).Interface())
				file.SetCellStyle(sheetName, cellName, cellName, bodyStyle)
				excelAdjustWidthHight(file, sheetName, colName, cellName, i+2, fmt.Sprint(a.Index(i).Field(c).Interface()))
			}
		}
	}

	exportRoot := "./media/export/"
	if _, err = os.Stat(exportRoot); os.IsNotExist(err) {
		os.MkdirAll(exportRoot, 0700)
		os.Create(exportRoot + "index.html")
	}

	fileName := GenerateBase64(24)
	for _, err = os.Stat("./media/export/" + fileName + ".xlsx"); os.IsExist(err); {
		fileName = GenerateBase64(24)
	}
	err = file.SaveAs("./media/export/" + fileName + ".xlsx")
	if err != nil {
		Trail(ERROR, "exportHandler unable to save file %s. %s", "./media/export/"+fileName+".xlsx", err)
	}
	http.Redirect(w, r, "/media/export/"+fileName+".xlsx", http.StatusSeeOther)
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
}

func excelAdjustWidthHight(file *excelize.File, sheetName, colName, cellName string, rowIndex int, content string) {
	cWidth, _ := file.GetColWidth(sheetName, colName)
	sWidth := float64(len(content))
	if sWidth > 40 {
		cHight, _ := file.GetRowHeight(sheetName, rowIndex)
		sHight := sWidth / 2
		if sHight > 100 {
			sHight = 100
		}

		if cHight < sHight {
			file.SetRowHeight(sheetName, rowIndex, sHight)
		}
		sWidth = 40
	}
	if cWidth < sWidth {
		file.SetColWidth(sheetName, colName, colName, sWidth)
	}
}
