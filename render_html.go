package uadmin

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

// RenderHTML creates a new template and applies a parsed template to the specified
// data object. For function, Tf is available by default and if you want to add functions
//to your template, just add them to funcs which will add them to the template with their
// original function names. If you added anonymous functions, they will be available in your
// templates as func1, func2 ...etc.
func RenderHTML(w http.ResponseWriter, path string, data interface{}, funcs ...interface{}) {
	var err error
	var funcVal reflect.Value
	var funcName string

	funcMap := template.FuncMap{
		"Tf": Tf,
	}

	for i := range funcs {
		funcVal = reflect.ValueOf(funcs[i])
		if funcVal.Type().Kind() != reflect.Func {
			Trail(WARNING, "Interface passed to RenderHTML in funcs parameter should only be a function. Got (%s) in position %d", funcVal.Type().Kind(), i)
			continue
		}

		funcName = runtime.FuncForPC(funcVal.Pointer()).Name()
		funcName = funcName[strings.LastIndex(funcName, ".")+1:]
		funcMap[funcName] = funcs[i]
	}

	t := template.New("").Funcs(funcMap)
	t, err = t.ParseFiles(path)
	if err != nil {
		fmt.Fprint(w, err.Error())
		Trail(ERROR, "RenderHTML unable to parse %s. %s", path, err)
		return
	}

	path = filepath.Base(path)
	err = t.ExecuteTemplate(w, path, data)
	if err != nil {
		ignoredErrors := []string{
			"write tcp",
		}
		for i := range ignoredErrors {
			if strings.HasPrefix(err.Error(), ignoredErrors[i]) {
				return
			}
		}
		Trail(ERROR, "Unable to render html template file (%s). %s", path, err)
	}
}
