package uadmin

import (
	"testing"
)

// TestSyncCustomTranslation is a unit testing function for syncCustomTranslation() function
func TestSyncCustomTranslation(t *testing.T) {
	// Activate a second language
	ar := Language{}
	Get(&ar, "code = ?", "ar")
	ar.Active = true
	ar.Save()

	results := syncCustomTranslation("uadmin/system")
	if len(results) != len(activeLangs) {
		t.Errorf("syncCustomTranslation didn't return status for all active languages. Got %d, expected %d", len(results), len(activeLangs))
	}

	results = syncCustomTranslation("uadmin/system/tags")
	if len(results) != 0 {
		t.Errorf("syncCustomTranslation didn't return 0 status invalid path. Got %d, expected %d", len(results), 0)
	}

	zh := Language{}
	Get(&zh, "code = ?", "zh")
	zh.Active = true
	zh.Save()

	results = syncCustomTranslation("uadmin/system")
	if len(results) != len(activeLangs) {
		t.Errorf("syncCustomTranslation didn't return status for all active languages. Got %d, expected %d", len(results), len(activeLangs))
	}

	// Clean up
	ar.Active = false
	ar.Save()
	zh.Active = false
	zh.Save()
}

// TestSyncModelTranslation is a unit testing function for syncModelTranslation() function
func TestSyncModelTranslation(t *testing.T) {
	// Activate a second language
	ar := Language{}
	Get(&ar, "code = ?", "ar")
	ar.Active = true
	ar.Save()

	results := syncModelTranslation(Schema["testmodelb"])
	if len(results) != len(activeLangs) {
		t.Errorf("syncModelTranslation didn't return status for all active languages. Got %d, expected %d", len(results), len(activeLangs))
	}

	zh := Language{}
	Get(&zh, "code = ?", "zh")
	zh.Active = true
	zh.Save()

	results = syncModelTranslation(Schema["testmodelb"])
	if len(results) != len(activeLangs) {
		t.Errorf("syncCustomTranslation didn't return status for all active languages. Got %d, expected %d", len(results), len(activeLangs))
	}

	s := Schema["testmodelb"]
	s.Fields = append(s.Fields, F{
		Name:        "TestField",
		DisplayName: "Test Field",
		Help:        "Help for test field",
		PatternMsg:  "test message",
		Choices:     []Choice{},
		ErrMsg:      "",
	})
	Schema["testmodelb"] = s
	results = syncModelTranslation(Schema["testmodelb"])
	if len(results) != len(activeLangs) {
		t.Errorf("syncCustomTranslation didn't return status for all active languages. Got %d, expected %d", len(results), len(activeLangs))
	}

	s.FieldByName("TestField").DisplayName = "`Updated Test Field"
	s.FieldByName("TestField").Help = "Updated Help for test field"
	s.FieldByName("TestField").PatternMsg = "Updated test message"
	Schema["testmodelb"] = s
	results = syncModelTranslation(Schema["testmodelb"])
	if len(results) != len(activeLangs) {
		t.Errorf("syncCustomTranslation didn't return status for all active languages. Got %d, expected %d", len(results), len(activeLangs))
	}

	// Clean up
	ar.Active = false
	ar.Save()
	zh.Active = false
	zh.Save()
	delete(Schema, "testmodelb")
	Schema["testmodelb"], _ = getSchema(TestModelB{})
}
