package dto

import (
	"fmt"
	"reflect"
)

type URLDTO struct {
	Identificador string `json:"identificador" required:"true"`
	Url           string `json:"dirurl" required:"true"`
	Etiquetas     string `json:"etiquetas"`
}

func ValidateStruct(s interface{}) error {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		requiredTag := field.Tag.Get("required")
		if requiredTag == "true" && isEmptyValue(value) {
			return fmt.Errorf("el campo %s es obligatorio", field.Name)
		}
	}
	return nil
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
