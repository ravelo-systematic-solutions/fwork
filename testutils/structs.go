package testutils

import (
	"reflect"
	"runtime"
)

func GetType(elem interface{}) string {
	t := reflect.TypeOf(elem)
	if t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else if t.Kind() == reflect.Func {
		return runtime.FuncForPC(reflect.ValueOf(elem).Pointer()).Name()
	} else {
		return t.Name()
	}
}
