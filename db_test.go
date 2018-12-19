package uadmin

import (
	"testing"
	"time"
)

type TestStruct struct {
	Model
	Name         string
	Children     []TestStruct
	Parent       *TestStruct
	ParentID     uint
	OtherModel   TestStruct1
	OtherModelID uint
}

type TestStruct1 struct {
	Model
	Name  string `uadmin:"search"`
	Value int
}

type TestType int

func (TestType) Active() TestType {
	return 1
}

func (TestType) Inactive() TestType {
	return 2
}

type TestStruct2 struct {
	Model
	Name           string
	Count          int
	Value          float64
	Start          time.Time
	End            *time.Time
	Type           TestType
	OtherModel     TestStruct1
	OtherModelID   uint
	AnotherModel   *TestStruct1
	AnotherModelID uint
	Active         bool
	Hidden         string `uadmin:"list_exclude"`
}

// TestInitializeDB is a unit testing function for initializeDB() function
func TestInitializeDB(t *testing.T) {
	examples := []struct {
		m    interface{}
		name string
		m2m  []string
	}{
		{TestStruct{}, "test_structs", []string{"teststruct_teststruct"}},
		{TestStruct2{}, "test_struct2", []string{}},
	}

	for _, e := range examples {
		initializeDB(e.m)

		if !db.HasTable(e.name) {
			t.Errorf("initializeDB didn't create table %s", e.name)
		}
		for _, name := range e.m2m {
			if !db.HasTable(name) {
				t.Errorf("initializeDB didn't create table %s", name)
			}
		}
	}
}

// TestSave is a unit testing function for Save() and customSave() function
func TestSave(t *testing.T) {
	Schema["teststruct"], _ = getSchema(TestStruct{})
	models["teststruct"] = TestStruct{}

	// Schema["teststruct1"], _ = getSchema(TestStruct1{})
	// models["teststruct1"] = TestStruct1{}

	Schema["teststruct2"], _ = getSchema(TestStruct2{})
	models["teststruct2"] = TestStruct2{}

	r1 := TestStruct{
		Name: "",
	}
	r2 := TestStruct{
		Name: "abc",
	}
	r3 := TestStruct{
		Name: "ABC",
	}
	m1 := TestStruct1{
		Name: "abc",
	}

	examples := []struct {
		m interface{}
	}{
		{&r1},
		{&r2},
		{&r3},
	}

	mExamples := []struct {
		m interface{}
	}{
		{&m1},
	}

	for _, e := range examples {
		Save(e.m)
	}
	if Count(TestStruct{}, "") != len(examples) {
		t.Errorf("Count is invalid after saving. Got %d expected %d", Count(TestStruct{}, ""), len(examples))
	}

	for _, e := range mExamples {
		Save(e.m)
	}
	if Count(TestStruct{}, "") != len(examples) {
		t.Errorf("Count is invalid after saving. Got %d expected %d", Count(TestStruct1{}, ""), len(mExamples))
	}

	r4 := TestStruct{
		Name:         "test",
		Children:     []TestStruct{r1, r2},
		ParentID:     r3.ID,
		Parent:       &r3,
		OtherModelID: m1.ID,
		OtherModel:   m1,
	}

	examples2 := []struct {
		m interface{}
	}{
		{&r4},
	}

	for _, e := range examples2 {
		Save(e.m)
	}
	if Count(TestStruct{}, "") != len(examples)+len(examples2) {
		t.Errorf("Count is invalid after saving. Got %d expected %d", Count(TestStruct{}, ""), len(examples)+len(examples2))
	}
	var count int
	db.Table("teststruct_teststruct").Count(&count)
	if count != len(r4.Children) {
		t.Errorf("M2M count is invalid after saving. Got %d expected %d", count, len(r4.Children))
	}

	r4.Children = []TestStruct{r1}
	Save(&r4)
	db.Table("teststruct_teststruct").Count(&count)
	if count != len(r4.Children) {
		t.Errorf("M2M count is invalid after saving. Got %d expected %d", count, len(r4.Children))
	}

	r5 := TestStruct{}
	Get(&r5, "name = ?", "test")
	if r4.ID != r5.ID {
		t.Errorf("Get didn't return the correct record.ID. Got %#v expected %#v", r5, r4)
	}
	if r4.Name != r5.Name {
		t.Errorf("Get didn't return the correct record.Name. Got %#v expected %#v", r5, r4)
	}
	if len(r4.Children) != len(r5.Children) {
		t.Errorf("Get didn't return the correct record.Children. Got %#v expected %#v", r5, r4)
	}
	if r4.ParentID != r5.ParentID {
		t.Errorf("Get didn't return the correct record.ParentID. Got %#v expected %#v", r5, r4)
	}

	// Test FilterBuilder
	f := map[string]interface{}{
		"name = ?":      "abc",
		"parent_id = ?": 0,
	}
	q, args := FilterBuilder(f)
	expectedArgs := []interface{}{"abc", 0}
	expectedQ1 := "name = ? AND parent_id = ?"
	expectedQ2 := "parent_id = ? AND name = ?"
	if q != expectedQ1 {
		if q != expectedQ2 {
			t.Errorf("FilterBuilder didn't return the correct q. Got %#v expected %#v", q, expectedQ2)
		} else {
			expectedArgs = []interface{}{0, "abc"}
		}
	}
	if len(args) != len(expectedArgs) {
		t.Errorf("FilterBuilder didn't return the correct len(args). Got %#v expected %#v", len(args), len(expectedArgs))
	} else {
		for i := range args {
			if args[i] != expectedArgs[i] {
				t.Errorf("FilterBuilder didn't return the correct args[%d]. Got %#v expected %#v", i, args[i], expectedArgs[i])
			}
		}
	}

	rows := []TestStruct{}
	Filter(&rows, q, args...)
	if len(rows) != 1 {
		t.Errorf("Filter didn't return the correct number of records. Got %d expected %d", len(rows), 1)
	}

	Preload(&r5)
	if (r4.Parent == nil) != (r5.Parent == nil) {
		t.Errorf("Preload didn't return the correct record.Parent. Got %#v expected %#v", r5, r4)
	} else {
		if (r4.Parent != nil) && (r5.Parent != nil) {
			if r4.Parent.ID != r5.Parent.ID {
				t.Errorf("Preload didn't return the correct record.Parent. Got %#v expected %#v", r5, r4)
			}
		}
	}
	if r4.OtherModel.ID != r5.OtherModel.ID {
		t.Errorf("Preload didn't return the correct record.OtherModel. Got %#v expected %#v", r5, r4)
	}

	// testing AdminPage
	rows = []TestStruct{}
	AdminPage("id", true, 0, -1, &rows, "name = ?", "abc")
	if len(rows) != 1 {
		t.Errorf("AdminPage didn't return the correct number of records. Got %d expected %d", len(rows), 1)
	}

	rows = []TestStruct{}
	AdminPage("id", true, 0, 1, &rows, "name = ?", "abc")
	if len(rows) != 1 {
		t.Errorf("AdminPage didn't return the correct number of records. Got %d expected %d", len(rows), 1)
	}

	Update(TestStruct{}, "name", "abc", "")
	rows = []TestStruct{}
	AdminPage("id", true, 0, -1, &rows, "name = ?", "abc")
	if len(rows) != len(examples)+len(examples2) {
		t.Errorf("AdminPage didn't return the correct number of records after Update. Got %d expected %d", len(rows), len(examples)+len(examples2))
	}

	Delete(r1)
	if Count(TestStruct{}, "") != len(examples)+len(examples2)-1 {
		t.Errorf("Count is invalid after Delete. Got %d expected %d", Count(TestStruct{}, ""), len(examples)+len(examples2)-1)
	}

	DeleteList(TestStruct{}, "")
	if Count(TestStruct{}, "") != 0 {
		t.Errorf("Count is invalid after DeleteList. Got %d expected %d", Count(TestStruct{}, ""), 0)
	}
}
