package uadmin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// TestGetFilter is a unit testing function for getFilter() function
func TestGetFilter(t *testing.T) {
	// Setup
	var record TestStruct2
	baseDate := time.Now()
	baseDate = time.Date(baseDate.Year(), baseDate.Month(), baseDate.Day(), 0, 0, 0, 0, baseDate.Location())
	om := TestStruct1{Name: "Test Model"}
	Save(&om)
	tx := db.Begin()
	for i := 0; i < 100; i++ {
		record = TestStruct2{
			Name:         fmt.Sprintf("Record No %d", i),
			Count:        i + 1,
			Value:        float64(i) + 0.5,
			Start:        baseDate.AddDate(0, 0, -i),
			OtherModelID: om.ID,
		}
		if i%2 == 0 {
			tempDate := record.Start.AddDate(0, 0, 1)
			record.End = &tempDate
			record.AnotherModelID = om.ID
			record.Active = true
		}
		record.Type = TestType(i % 3)
		tx.Create(&record)
	}
	tx.Commit()

	examples := []struct {
		r      *http.Request
		count  int
		header []string
	}{
		{
			r:     httptest.NewRequest("GET", fmt.Sprintf("/export/?m=teststruct2&start__lte=%s&start__gte=%s", baseDate.Format("2006-01-02"), baseDate.AddDate(0, 0, -10).Format("2006-01-02")), nil),
			count: 11,
			header: []string{
				"Name", "Count", "Value", "Start", "End", "Type", "Other Model", "Another Model", "Active",
			},
		},
		{
			r:     httptest.NewRequest("GET", fmt.Sprintf("/export/?m=teststruct2&start__lt=%s&start__gt=%s", baseDate.Format("2006-01-02"), baseDate.AddDate(0, 0, -10).Format("2006-01-02")), nil),
			count: 9,
			header: []string{
				"Name", "Count", "Value", "Start", "End", "Type", "Other Model", "Another Model", "Active",
			},
		},
		{
			r:     httptest.NewRequest("GET", fmt.Sprintf("/export/?m=teststruct2&count__in=%s", "1,2,3,4"), nil),
			count: 4,
			header: []string{
				"Name", "Count", "Value", "Start", "End", "Type", "Other Model", "Another Model", "Active",
			},
		},
		{
			r:     httptest.NewRequest("GET", fmt.Sprintf("/export/?m=teststruct2&name__contains=%s", url.QueryEscape("Record No 1")), nil),
			count: 11,
			header: []string{
				"Name", "Count", "Value", "Start", "End", "Type", "Other Model", "Another Model", "Active",
			},
		},
		{
			r:     httptest.NewRequest("GET", fmt.Sprintf("/export/?m=teststruct2&active=%s", "1"), nil),
			count: 50,
			header: []string{
				"Name", "Count", "Value", "Start", "End", "Type", "Other Model", "Another Model", "Active",
			},
		},
	}

	s1 := Session{
		Active: true,
		UserID: 1,
	}
	s1.GenerateKey()
	s1.Save()
	Preload(&s1)

	for _, e := range examples {
		w := httptest.NewRecorder()
		exportHandler(w, e.r, &s1)
		// Check if 303
		if w.Code != http.StatusSeeOther {
			t.Errorf("exportHandler returned invalid code. Expected %d, got %d", 303, w.Code)
			continue
		}
		// Check if the file exists

		if _, ok := w.Header()["Location"]; !ok {
			t.Error("exportHandler returned no Location in header")
			continue
		}
		if len(w.Header()["Location"]) < 1 {
			t.Error("exportHandler returned empty Location in header")
			continue
		}
		if _, err := os.Stat("." + w.Header()["Location"][0]); os.IsNotExist(err) {
			t.Errorf("exportHandler didn't create a file. Expected %s", w.Header()["Location"][0])
			continue
		}
		f, err := excelize.OpenFile("." + w.Header()["Location"][0])
		if err != nil {
			t.Errorf("exportHandler created an invalid xlsx file. %s", err)
			continue
		}
		if f.SheetCount != 1 {
			t.Errorf("exportHandler invalid number of sheets. Expected 1, got %d", f.SheetCount)
			continue
		}
		sheetName := "Sheet1"
		for i := range e.header {
			colName, _ := excelize.ColumnNumberToName(i + 1)
			header, err := f.GetCellValue(sheetName, colName+"1")
			if err != nil {
				t.Errorf("exportHandler cell was not found. %s", err)
				continue
			}
			if header != e.header[i] {
				t.Errorf("exportHandler invalid header. Expected %s, got %s", e.header[i], header)
				continue
			}
		}
		// TODO: Test data
		// for col := range e.header {
		// 	for row := range sheet.Rows{
		// 		if sheet.Cell(i+1, 1).String() == e.header[i] {
		// 			t.Errorf("exportHandler invalid header. Expected %s, got %s", e.header[i], sheet.Cell(i+1, 1).String())
		// 		}
		// 	}
		// }
	}
	Delete(s1)
	DeleteList(&TestStruct1{}, "")
	DeleteList(&TestStruct2{}, "")
}
