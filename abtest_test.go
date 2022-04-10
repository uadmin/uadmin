package uadmin

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestABTest(t *testing.T) {
	currentAbTestsCount := abTestCount
	test01 := ABTest{
		Name:   "test 01",
		Type:   ABTestType(0).Static(),
		Active: true,
		Group:  "test 1",
	}
	test01.Save()

	test01.Save()

	if test01.ID == 0 {
		t.Errorf("ABTest was not saved correctly. Got 0 ID")
	}
	if abTestCount != currentAbTestsCount+1 {
		t.Errorf("abTestCount didn't increment correctly. Expected %d but got %d", currentAbTestsCount+1, abTestCount)
	}

	test02 := ABTest{
		Name:      "test 02",
		Type:      ABTestType(0).Model(),
		ModelName: ModelList(0),
		Field:     FieldList(0),
		Active:    true,
		Group:     "test 1",
	}
	test02.Save()

	if test02.ID == 0 {
		t.Errorf("ABTest was not saved correctly. Got 0 ID")
	}
	if abTestCount != currentAbTestsCount+2 {
		t.Errorf("abTestCount didn't increment correctly. Expected %d but got %d", currentAbTestsCount+1, abTestCount)
	}

	// Add Values
	value1A := ABTestValue{
		ABTestID: test01.ID,
		Value:    "A",
		Active:   true,
	}
	value1B := ABTestValue{
		ABTestID: test01.ID,
		Value:    "B",
		Active:   true,
	}
	Save(&value1A)
	Save(&value1B)
	value2A := ABTestValue{
		ABTestID: test02.ID,
		Value:    "A",
		Active:   true,
	}
	value2B := ABTestValue{
		ABTestID: test02.ID,
		Value:    "B",
		Active:   true,
	}
	Save(&value2A)
	Save(&value2B)
	syncABTests()
	r := &http.Request{}
	r.Header = http.Header{}
	abt := getABT(r)
	r.AddCookie(&http.Cookie{
		Name:  "abt",
		Value: fmt.Sprint(abt),
		Path:  "/",
	})
	ABTestClick(r, "test 1")
	time.Sleep(time.Millisecond * 10)
	// We lock and unlock to ensure that ABTestClick is done
	abTestsMutex.Lock()
	time.Sleep(time.Millisecond)
	abTestsMutex.Unlock()
	syncABTests()

	// get abt

	// Check click
	clicksRight := 0
	clicksWrong := 0
	if abt%2 == 0 {
		GetValueSorted("ab_test_values", "clicks", "", true, &clicksRight, "id = ?", value1A.ID)
		GetValueSorted("ab_test_values", "clicks", "", true, &clicksWrong, "id = ?", value1B.ID)
	} else {
		GetValueSorted("ab_test_values", "clicks", "", true, &clicksWrong, "id = ?", value1A.ID)
		GetValueSorted("ab_test_values", "clicks", "", true, &clicksRight, "id = ?", value1B.ID)
	}
	if clicksRight != 1 {
		t.Errorf("Test 1 Expected 1 click for the right value got %d", clicksRight)
	}
	if clicksWrong != 0 {
		t.Errorf("Test 1 Expected 0 click for the wrong value got %d", clicksWrong)
	}
	clicksRight = 0
	clicksWrong = 0
	if abt%2 == 0 {
		GetValueSorted("ab_test_values", "clicks", "", true, &clicksRight, "id = ?", value2A.ID)
		GetValueSorted("ab_test_values", "clicks", "", true, &clicksWrong, "id = ?", value2B.ID)
	} else {
		GetValueSorted("ab_test_values", "clicks", "", true, &clicksWrong, "id = ?", value2A.ID)
		GetValueSorted("ab_test_values", "clicks", "", true, &clicksRight, "id = ?", value2B.ID)
	}
	if clicksRight != 1 {
		t.Errorf("Test 2 Expected 1 click for the right value got %d", clicksRight)
	}
	if clicksWrong != 0 {
		t.Errorf("Test 2 Expected 0 click for the wrong value got %d", clicksWrong)
	}

	// test reset
	test01.Reset()
	clicksRight = 0
	clicksWrong = 0
	GetValueSorted("ab_test_values", "clicks", "", true, &clicksRight, "id = ?", value1A.ID)
	GetValueSorted("ab_test_values", "clicks", "", true, &clicksWrong, "id = ?", value1B.ID)
	if clicksRight != 0 {
		t.Errorf("Expected 0 click for the right value got %d", clicksRight)
	}
	if clicksWrong != 0 {
		t.Errorf("Expected 0 click for the wrong value got %d", clicksWrong)
	}

	// Clean up
	Delete(&test01)
}

func TestLoadModels(t *testing.T) {
	choiceList := loadModels(nil, nil)
	if len(modelList) != len(choiceList) {
		t.Errorf("loadModels didn't return the correct number of choices, expected %d but got %d", len(modelList), len(choiceList))
	}
}

func TestLoadFields(t *testing.T) {
	test := ABTest{
		Type: ABTestType(0).Static(),
	}
	choiceList := loadFields(test, nil)
	if len(choiceList) != 0 {
		t.Errorf("loadFields didn't return the correct number of choices, expected %d but got %d", 0, len(choiceList))
	}
	choiceList = loadFields(&test, nil)
	if len(choiceList) != 0 {
		t.Errorf("loadFields on pointer didn't return the correct number of choices, expected %d but got %d", 0, len(choiceList))
	}
	choiceList = loadFields(struct{}{}, nil)
	if len(choiceList) != 0 {
		t.Errorf("loadFields on invalid didn't return the correct number of choices, expected %d but got %d", 0, len(choiceList))
	}

	test.Type = ABTestType(0).Model()
	test.ModelName = ModelList(0)
	choiceList = loadFields(test, nil)
	s := Schema[getModelName(modelList[0])]
	if len(choiceList) != len(s.Fields) {
		t.Errorf("loadFields didn't return the correct number of choices, expected %d but got %d", len(s.Fields), len(choiceList))
	}
}
