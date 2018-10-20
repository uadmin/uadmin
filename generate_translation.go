package uadmin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

/*
// Field appears in JSON as key "myName".
Field int `json:"myName"`

// Field appears in JSON as key "myName" and
// the field is omitted from the object if its value is empty,
// as defined above.
Field int `json:"myName,omitempty"`

// Field appears in JSON as key "Field" (the default), but
// the field is skipped if empty.
// Note the leading comma.
Field int `json:",omitempty"`

// Field is ignored by this package.
Field int `json:"-"`

// Field appears in JSON as key "-".
Field int `json:"-,"`
*/

type fieldLanguage struct {
	DisplayName string            `json:"display_name"`
	Help        string            `json:"help"`
	PatternMsg  string            `json:"pattern_msg"`
	ErrMsg      map[string]string `json:"err_msg"`
}

type structLanguage struct {
	DisplayName string                   `json:"display_name"`
	Fields      map[string]fieldLanguage `json:"fields"`
}

func generateTranslation(m ModelSchema) {
	structLang := structLanguage{}
	structLang.DisplayName = m.DisplayName
	structLang.Fields = map[string]fieldLanguage{}
	for _, f := range m.Fields {
		structLang.Fields[f.Name] = fieldLanguage{
			DisplayName: f.DisplayName,
			Help:        f.Help,
			PatternMsg:  f.PatternMsg,
			ErrMsg:      map[string]string{},
		}
	}
	pkgName := fmt.Sprint(reflect.TypeOf(models[strings.ToLower(m.Name)]))
	l, _ := json.MarshalIndent(structLang, "", "  ")
	os.MkdirAll("./static/i18n/"+pkgName+"/", 0744)
	translated := []byte(l)
	err := ioutil.WriteFile("./static/i18n/"+pkgName+"/"+m.Name+"_en.json", translated, 0644)
	if err != nil {
		Trail(ERROR, "generateTranslation error writing a file %v", err)
	}
}
