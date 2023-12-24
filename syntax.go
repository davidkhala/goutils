package goutils

import (
	"context"
	"reflect"
)

// GetGoContext used by the initialization
func GetGoContext() context.Context {
	return context.Background()
}

func GetType(object interface{}) string {
	if t := reflect.TypeOf(object); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
func Map[T, V any](array []T, fn func(T) V) []V {
	result := make([]V, len(array))
	for i, t := range array {
		result[i] = fn(t)
	}
	return result
}
