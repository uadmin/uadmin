package uadmin

import (
	"fmt"
	"reflect"
	"strings"
)

// GetID !
func GetID(m reflect.Value) uint {
	if m.Kind() == reflect.Ptr {
		return uint(m.Elem().FieldByName("ID").Uint())
	}
	return uint(m.FieldByName("ID").Uint())
}

// GetString returns string representation on an instance of
// a model
func GetString(a interface{}) string {
	str, ok := a.(fmt.Stringer)
	if ok {
		return str.String()
	}
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	if t.Kind() != reflect.Ptr {
		v = reflect.Indirect(reflect.New(t))
		v.Set(reflect.ValueOf(a))
		sp := v.Addr().Interface()
		str, ok := sp.(fmt.Stringer)
		if ok {
			return str.String()
		}
		if _, ok := t.FieldByName("Name"); ok {
			return v.FieldByName("Name").String()
		}
	} else {
		if _, ok := t.Elem().FieldByName("Name"); ok {
			return v.Elem().FieldByName("Name").String()
		}
	}
	return fmt.Sprint(a)
}

// getChoices return a list of choices
func getChoices(ModelName string) []Choice {
	choices := []Choice{
		Choice{" - ", 0, false},
	}

	m, ok := NewModelArray(strings.ToLower(ModelName), true)

	// If no model exists, return an empty choices list
	if !ok {
		return choices
	}

	// Get all choices
	All(m.Addr().Interface())
	for i := 0; i < m.Len(); i++ {
		id := GetID(m.Index(i))
		choices = append(choices, Choice{fmt.Sprint(m.Index(i).Interface()), uint(id), false})
	}
	return choices
}
