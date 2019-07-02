package uadmin

import (
	"testing"
)

// TestGetSchema is a unit testing function for getSchema() function
func TestGetSchema(t *testing.T) {
	var schema ModelSchema
	var expectedSchema ModelSchema
	var ok bool

	schema, ok = getSchema(TestModelA{})
	if !ok {
		t.Errorf("getSchema could not parse Model1: %#v", TestModelA{})
	}

	expectedSchema = ModelSchema{
		Name:        "TestModelA",
		ModelName:   "testmodela",
		DisplayName: "Test Model A",
		Inlines:     []*ModelSchema{},
		InlinesData: []listData{},
		Fields: []F{
			{
				Name:              "ID",
				DisplayName:       "ID",
				Type:              cID,
				TypeName:          "Model",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			{
				Name:              "Name",
				DisplayName:       "Name",
				Type:              cSTRING,
				TypeName:          "string",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
		},
	}
	compareSchema("TestModelA", schema, expectedSchema, t)

	schema, ok = getSchema(TestModelB{})
	if !ok {
		t.Errorf("getSchema could not parse Model1: %#v", TestModelA{})
	}

	expectedSchema = ModelSchema{
		Name:        "TestModelB",
		ModelName:   "testmodelb",
		DisplayName: "Test Model B",
		Inlines:     []*ModelSchema{},
		InlinesData: []listData{},
		Fields: []F{
			// Model
			{
				Name:              "ID",
				DisplayName:       "ID",
				Type:              cID,
				TypeName:          "Model",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			// Name         string     `uadmin:"help:This is a test help message;search;list_exclude"`
			{
				Name:              "Name",
				DisplayName:       "Name",
				Type:              cSTRING,
				TypeName:          "string",
				Value:             nil,
				Help:              "This is a test help message",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        true,
				Filter:            false,
				ListDisplay:       false,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			// ItemCount    int        `uadmin:"max:5;min:1;format:%03d;required;read_only:edit"`
			{
				Name:              "ItemCount",
				DisplayName:       "Item Count",
				Type:              cNUMBER,
				TypeName:          "int",
				Value:             nil,
				Help:              "",
				Max:               "5",
				Min:               "1",
				Format:            "%03d",
				DefaultValue:      "",
				Required:          true,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "true,edit",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			// Phone        string     `uadmin:"default_value:09;pattern:[0-9+]{7,15};pattern_msg:invalid phone number;encrypt"`
			{
				Name:              "Phone",
				DisplayName:       "Phone",
				Type:              cSTRING,
				TypeName:          "string",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "09",
				Required:          false,
				Pattern:           "[0-9+]{7,15}",
				PatternMsg:        "invalid phone number",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           true,
			},
			// Active       bool       `uadmin:"hidden;read_only"`
			{
				Name:              "Active",
				DisplayName:       "Active",
				Type:              cBOOL,
				TypeName:          "bool",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            true,
				ReadOnly:          cTRUE,
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			// OtherModel   TestModelA `uadmin:"categorical_filter;filter;read_only:new"`
			{
				Name:              "OtherModel",
				DisplayName:       "Other Model",
				Type:              cFK,
				TypeName:          "TestModelA",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "true,new",
				Searchable:        false,
				Filter:            true,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: true,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			// ModelAList   []TestModelA
			{
				Name:              "ModelAList",
				DisplayName:       "Model A List",
				Type:              cM2M,
				TypeName:          "TestModelA",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			// Parent       *TestModelB
			{
				Name:              "Parent",
				DisplayName:       "Parent",
				Type:              cFK,
				TypeName:          "TestModelB",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			// Email        string  `uadmin:"email"`
			{
				Name:              "Email",
				DisplayName:       "Email",
				Type:              cEMAIL,
				TypeName:          "string",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			// Greeting     string  `uadmin:"multilingual"`
			{
				Name:              "Greeting",
				DisplayName:       "Greeting",
				Type:              cMULTILINGUAL,
				TypeName:          "string",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation{{Name: "English (English)", Code: "en", Flag: "", Default: true, Active: true, Value: ""}},
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			// Image        string  `uadmin:"image;upload_to:/media/home/me/images/"`
			{
				Name:              "Image",
				DisplayName:       "Image",
				Type:              cIMAGE,
				TypeName:          "string",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "/media/home/me/images/",
				Encrypt:           false,
			},
			// File         string  `uadmin:"file;upload_to:/media/home/me/files"`
			{
				Name:              "File",
				DisplayName:       "File",
				Type:              cFILE,
				TypeName:          "string",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "/media/home/me/files/",
				Encrypt:           false,
			},
			// Secret       string  `uadmin:"password"`
			{
				Name:              "Secret",
				DisplayName:       "Secret",
				Type:              cPASSWORD,
				TypeName:          "string",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			// Description  string  `uadmin:"html"`
			{
				Name:              "Description",
				DisplayName:       "Description",
				Type:              cHTML,
				TypeName:          "string",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			// URL          string  `uadmin:"link"`
			{
				Name:              "URL",
				DisplayName:       "URL",
				Type:              cLINK,
				TypeName:          "string",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			// Code         string  `uadmin:"code"`
			{
				Name:              "Code",
				DisplayName:       "Code",
				Type:              cCODE,
				TypeName:          "string",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			// P1           int     `uadmin:"progress_bar"`
			{
				Name:              "P1",
				DisplayName:       "P 1",
				Type:              cPROGRESSBAR,
				TypeName:          "int",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string{100.0: defaultProgressBarColor},
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			// P2           float64 `uadmin:"progress_bar"`
			{
				Name:              "P2",
				DisplayName:       "P 2",
				Type:              cPROGRESSBAR,
				TypeName:          "float64",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string{100.0: defaultProgressBarColor},
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			// P3           float64 `uadmin:"progress_bar:1.0"`
			{
				Name:              "P3",
				DisplayName:       "P 3",
				Type:              cPROGRESSBAR,
				TypeName:          "float64",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string{1.0: defaultProgressBarColor},
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			// P4           float64 `uadmin:"progress_bar:1.0:red"`
			{
				Name:              "P4",
				DisplayName:       "P 4",
				Type:              cPROGRESSBAR,
				TypeName:          "float64",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string{1.0: "red"},
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			// P5           float64 `uadmin:"progress_bar:1.0:#f00"`
			{
				Name:              "P5",
				DisplayName:       "P 5",
				Type:              cPROGRESSBAR,
				TypeName:          "float64",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string{1.0: "#f00"},
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			// P6           float64 `uadmin:"progress_bar:0.3:red,0.7:yellow,1.0:lime"`
			{
				Name:              "P6",
				DisplayName:       "P 6",
				Type:              cPROGRESSBAR,
				TypeName:          "float64",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string{0.3: "red", 0.7: "yellow", 1.0: "lime"},
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			// Price        float64 `uadmin:"money"`
			{
				Name:              "Price",
				DisplayName:       "Price",
				Type:              cMONEY,
				TypeName:          "float64",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			// List         testList
			{
				Name:              "List",
				DisplayName:       "List",
				Type:              cLIST,
				TypeName:          "testList",
				Value:             nil,
				Help:              "",
				Max:               "",
				Min:               "",
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          "",
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice{{K: 0, V: " - "}, {K: 1, V: "A"}},
				IsMethod:          false,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
			/*
			   func (TestModelB) Method__List__Form() string {
			   	return "Value"
			   }
			*/
			{
				Name:              "Method__List__Form",
				DisplayName:       "Method",
				Type:              cSTRING,
				TypeName:          "",
				Value:             nil,
				Help:              "",
				Max:               nil,
				Min:               nil,
				Format:            "",
				DefaultValue:      "",
				Required:          false,
				Pattern:           "",
				PatternMsg:        "",
				Hidden:            false,
				ReadOnly:          cTRUE,
				Searchable:        false,
				Filter:            false,
				ListDisplay:       true,
				FormDisplay:       true,
				CategoricalFilter: false,
				Translations:      []translation(nil),
				Choices:           []Choice(nil),
				IsMethod:          true,
				ErrMsg:            "",
				ProgressBar:       map[float64]string(nil),
				LimitChoicesTo:    nil,
				UploadTo:          "",
				Encrypt:           false,
			},
		},
	}
	compareSchema("TestModelB", schema, expectedSchema, t)
}

func compareSchema(modelName string, got, expected ModelSchema, t *testing.T) {
	if got.Name != expected.Name {
		t.Errorf("getSchema invalid schema.Name: (%s) expected (%s) in %s", got.Name, expected.Name, modelName)
	}
	if got.ModelName != expected.ModelName {
		t.Errorf("getSchema invalid schema.ModelName: (%s) expected (%s) in %s", got.ModelName, expected.ModelName, modelName)
	}
	if got.DisplayName != expected.DisplayName {
		t.Errorf("getSchema invalid schema.DisplayName: (%s) expected (%s) in %s", got.DisplayName, expected.DisplayName, modelName)
	}
	if len(got.Fields) != len(expected.Fields) {
		t.Errorf("getSchema invalid number of fields: (%d) expected (%d) in %s", len(got.Fields), len(expected.Fields), modelName)
	}
	for i := range got.Fields {
		if got.Fields[i].Type != expected.Fields[i].Type {
			t.Errorf("getSchema F.Type: (%s) expected (%s) in %s.%s", got.Fields[i].Type, expected.Fields[i].Type, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].Name != expected.Fields[i].Name {
			t.Errorf("getSchema F.Name: (%s) expected (%s) in %s.%s", got.Fields[i].Name, expected.Fields[i].Name, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].DisplayName != expected.Fields[i].DisplayName {
			t.Errorf("getSchema F.DisplayName: (%s) expected (%s) in %s.%s", got.Fields[i].DisplayName, expected.Fields[i].DisplayName, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].TypeName != expected.Fields[i].TypeName {
			t.Errorf("getSchema F.TypeName: (%s) expected (%s) in %s.%s", got.Fields[i].TypeName, expected.Fields[i].TypeName, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].Value != expected.Fields[i].Value {
			t.Errorf("getSchema F.Value: (%#v) expected (%#v) in %s.%s", got.Fields[i].Value, expected.Fields[i].Value, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].Help != expected.Fields[i].Help {
			t.Errorf("getSchema F.Help: (%s) expected (%s) in %s.%s", got.Fields[i].Help, expected.Fields[i].Help, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].Max != expected.Fields[i].Max {
			t.Errorf("getSchema F.Max: (%#v) expected (%#v) in %s.%s", got.Fields[i].Max, expected.Fields[i].Max, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].Min != expected.Fields[i].Min {
			t.Errorf("getSchema F.Min: (%#v) expected (%#v) in %s.%s", got.Fields[i].Min, expected.Fields[i].Min, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].Format != expected.Fields[i].Format {
			t.Errorf("getSchema F.Format: (%s) expected (%s) in %s.%s", got.Fields[i].Format, expected.Fields[i].Format, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].DefaultValue != expected.Fields[i].DefaultValue {
			t.Errorf("getSchema F.DefaultValue: (%s) expected (%s) in %s.%s", got.Fields[i].DefaultValue, expected.Fields[i].DefaultValue, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].Required != expected.Fields[i].Required {
			t.Errorf("getSchema F.Required: (%v) expected (%v) in %s.%s", got.Fields[i].Required, expected.Fields[i].Required, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].Pattern != expected.Fields[i].Pattern {
			t.Errorf("getSchema F.Pattern: (%s) expected (%s) in %s.%s", got.Fields[i].Pattern, expected.Fields[i].Pattern, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].PatternMsg != expected.Fields[i].PatternMsg {
			t.Errorf("getSchema F.PatternMsg: (%s) expected (%s) in %s.%s", got.Fields[i].PatternMsg, expected.Fields[i].PatternMsg, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].Hidden != expected.Fields[i].Hidden {
			t.Errorf("getSchema F.Hidden: (%v) expected (%v) in %s.%s", got.Fields[i].Hidden, expected.Fields[i].Hidden, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].ReadOnly != expected.Fields[i].ReadOnly {
			t.Errorf("getSchema F.ReadOnly: (%s) expected (%s) in %s.%s", got.Fields[i].ReadOnly, expected.Fields[i].ReadOnly, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].Searchable != expected.Fields[i].Searchable {
			t.Errorf("getSchema F.Searchable: (%v) expected (%v) in %s.%s", got.Fields[i].Searchable, expected.Fields[i].Searchable, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].Filter != expected.Fields[i].Filter {
			t.Errorf("getSchema F.Filter: (%v) expected (%v) in %s.%s", got.Fields[i].Filter, expected.Fields[i].Filter, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].ListDisplay != expected.Fields[i].ListDisplay {
			t.Errorf("getSchema F.ListDisplay: (%v) expected (%v) in %s.%s", got.Fields[i].ListDisplay, expected.Fields[i].ListDisplay, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].FormDisplay != expected.Fields[i].FormDisplay {
			t.Errorf("getSchema F.FormDisplay: (%v) expected (%v) in %s.%s", got.Fields[i].FormDisplay, expected.Fields[i].FormDisplay, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].CategoricalFilter != expected.Fields[i].CategoricalFilter {
			t.Errorf("getSchema F.CategoricalFilter: (%v) expected (%v) in %s.%s", got.Fields[i].CategoricalFilter, expected.Fields[i].CategoricalFilter, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].Translations != nil && expected.Fields[i].Translations != nil {
			if len(got.Fields[i].Translations) == len(expected.Fields[i].Translations) {
				for index := range expected.Fields[i].Translations {
					if got.Fields[i].Translations[index] != expected.Fields[i].Translations[index] {
						t.Errorf("getSchema F.Translations: (%#v) expected (%#v) in %s.%s[%d]", got.Fields[i].Translations[index], expected.Fields[i].Translations[index], modelName, expected.Fields[i].Name, index)
					}
				}
			} else {
				t.Errorf("getSchema F.Translations: (%#v) expected (%#v) in %s.%s", got.Fields[i].Translations, expected.Fields[i].Translations, modelName, expected.Fields[i].Name)
			}
		} else {
			if (got.Fields[i].Translations != nil) != (expected.Fields[i].Translations != nil) {
				t.Errorf("getSchema F.Translations: (%#v) expected (%#v) in %s.%s", got.Fields[i].Translations, expected.Fields[i].Translations, modelName, expected.Fields[i].Name)
			}
		}
		if got.Fields[i].Choices != nil && expected.Fields[i].Choices != nil {
			if len(got.Fields[i].Choices) == len(expected.Fields[i].Choices) {
				for index := range expected.Fields[i].Choices {
					if got.Fields[i].Choices[index] != expected.Fields[i].Choices[index] {
						t.Errorf("getSchema F.Choices: (%#v) expected (%#v) in %s.%s[%d]", got.Fields[i].Choices[index], expected.Fields[i].Choices[index], modelName, expected.Fields[i].Name, index)
					}
				}
			} else {

				t.Errorf("getSchema F.Choices: (%#v) expected (%#v) in %s.%s", got.Fields[i].Choices, expected.Fields[i].Choices, modelName, expected.Fields[i].Name)
			}
		} else {
			if (got.Fields[i].Choices != nil) != (expected.Fields[i].Choices != nil) {
				t.Errorf("getSchema F.Choices: (%#v) expected (%#v) in %s.%s", got.Fields[i].Choices, expected.Fields[i].Choices, modelName, expected.Fields[i].Name)
			}
		}
		if got.Fields[i].IsMethod != expected.Fields[i].IsMethod {
			t.Errorf("getSchema F.IsMethod: (%v) expected (%v) in %s.%s", got.Fields[i].IsMethod, expected.Fields[i].IsMethod, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].ErrMsg != expected.Fields[i].ErrMsg {
			t.Errorf("getSchema F.ErrMsg: (%s) expected (%s) in %s.%s", got.Fields[i].ErrMsg, expected.Fields[i].ErrMsg, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].ProgressBar != nil && expected.Fields[i].ProgressBar != nil {
			if len(got.Fields[i].ProgressBar) == len(expected.Fields[i].ProgressBar) {
				for index := range expected.Fields[i].ProgressBar {
					if got.Fields[i].ProgressBar[index] != expected.Fields[i].ProgressBar[index] {
						t.Errorf("getSchema F.ProgressBar: (%#v) expected (%#v) in %s.%s[%f]", got.Fields[i].ProgressBar[index], expected.Fields[i].ProgressBar[index], modelName, expected.Fields[i].Name, index)
					}
				}
			} else {
				t.Errorf("getSchema F.ProgressBar: (%#v) expected (%#v) in %s.%s", got.Fields[i].ProgressBar, expected.Fields[i].ProgressBar, modelName, expected.Fields[i].Name)
			}
		} else {
			if (got.Fields[i].ProgressBar != nil) != (expected.Fields[i].ProgressBar != nil) {
				t.Errorf("getSchema F.ProgressBar: (%#v) expected (%#v) in %s.%s", got.Fields[i].ProgressBar, expected.Fields[i].ProgressBar, modelName, expected.Fields[i].Name)
			}
		}
		if got.Fields[i].LimitChoicesTo != nil && expected.Fields[i].LimitChoicesTo != nil {
			t.Errorf("getSchema F.LimitChoicesTo: (%v) expected (%v) in %s.%s", got.Fields[i].LimitChoicesTo == nil, expected.Fields[i].LimitChoicesTo == nil, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].UploadTo != expected.Fields[i].UploadTo {
			t.Errorf("getSchema F.UploadTo: (%s) expected (%s) in %s.%s", got.Fields[i].UploadTo, expected.Fields[i].UploadTo, modelName, expected.Fields[i].Name)
		}
		if got.Fields[i].Encrypt != expected.Fields[i].Encrypt {
			t.Errorf("getSchema F.Encrypt: (%v) expected (%v) in %s.%s", got.Fields[i].Encrypt, expected.Fields[i].Encrypt, modelName, expected.Fields[i].Name)
		}
	}
}
