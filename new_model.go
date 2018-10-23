package uadmin

import (
	"reflect"
)

// newModel creates a new model from a model name
func newModel(modelName string, pointer bool) (reflect.Value, bool) {
	model := models[modelName]
	if model == nil {
		return reflect.ValueOf(nil), false
	}
	m := reflect.New(reflect.TypeOf(model)).Elem()
	if pointer {
		m = m.Addr()
	}
	return m, true
}

// newModelArray creates a new model from a model name
func newModelArray(modelName string, pointer bool) (reflect.Value, bool) {
	model := models[modelName]
	if model == nil {
		return reflect.ValueOf(nil), false
	}
	modelType := reflect.TypeOf(model)
	m := reflect.New(reflect.SliceOf(modelType)).Elem()
	if pointer {
		m = m.Addr()
	}
	return m, true
}

// deepCopy for ModelSchema
func deepCopy(a interface{}) interface{} {
	b := a
	valueOfA := reflect.ValueOf(a)
	valueOfB := reflect.New(reflect.TypeOf(a)).Elem()
	valueOfB.Set(reflect.ValueOf(b))
	t := reflect.TypeOf(a)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type.Kind() == reflect.Slice {
			dst := reflect.New(reflect.SliceOf(t.Field(i).Type.Elem())).Elem()
			for index := 0; index < valueOfA.Field(i).Len(); index++ {
				temp := valueOfA.Field(i).Index(index)
				dst = reflect.Append(dst, temp)

			}
			valueOfB.Addr().Elem().Field(i).Set(dst)
		}
	}
	return valueOfB.Interface()
}
