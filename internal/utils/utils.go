package utils

import (
	"log/slog"
	"reflect"
)

// ConvertToInterfaceSlice converts a slice of any type to a slice of interface{}
func ConvertToInterfaceSlice(slice interface{}) []interface{} {
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		slog.Error("ConvertToInterfaceSlice: input is not a slice")
	}

	s := make([]interface{}, val.Len())
	for i := 0; i < val.Len(); i++ {
		s[i] = val.Index(i).Interface()
	}

	return s
}
