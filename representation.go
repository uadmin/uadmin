package uadmin

import (
	"fmt"
	"reflect"
)

// getID !
func getID(m reflect.Value) uint {
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
