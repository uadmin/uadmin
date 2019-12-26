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
	if t.Kind() != reflect.Ptr && t.Kind() == reflect.Struct {
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
	} else if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct {
		// Check if nil
		if v.IsNil() {
			return ""
		}

		if _, ok := t.Elem().FieldByName("Name"); ok {
			return v.Elem().FieldByName("Name").String()
		}
	} else if t.Kind() == reflect.Int && t.Name() != "int" {
		val := v.Int()
		// This is a static list type
		for i := 0; i < v.NumMethod(); i++ {
			ret := v.Method(i).Call([]reflect.Value{})
			if len(ret) > 0 {
				if ret[0].Int() == val {
					return t.Method(i).Name
				}
			}
		}
	}
	return fmt.Sprint(a)
}

// getChoices return a list of choices
func getChoices(ModelName string) []Choice {
	choices := []Choice{
		{" - ", 0, false},
	}

	m, ok := NewModelArray(strings.ToLower(ModelName), false)

	// If no model exists, return an empty choices list
	if !ok {
		return choices
	}
	//TODO: implement limit choices to
	// Get all choices
	All(m.Addr().Interface())
	for i := 0; i < m.Len(); i++ {
		id := GetID(m.Index(i))
		choices = append(choices, Choice{GetString(m.Index(i).Interface()), uint(id), false})
	}
	return choices
}

// getModelName returns the name of a model
func getModelName(a interface{}) string {
	if val, ok := a.(reflect.Value); ok {
		return getModelName(val.Interface())
	}
	if val, ok := a.(*reflect.Value); ok {
		return getModelName(val.Elem().Interface())
	}
	if reflect.TypeOf(a).Kind() == reflect.Ptr {
		return getModelName(reflect.ValueOf(a).Elem().Interface())
	}
	if reflect.TypeOf(a).Kind() == reflect.Slice {
		return getModelName(reflect.New(reflect.TypeOf(a).Elem()))
	}
	//if val, ok := a.(reflect.Type); ok {
	//	return strings.ToLower(val.Name())
	//}
	return strings.ToLower(reflect.TypeOf(a).Name())
}
