package goutils

import (
	"errors"
	"fmt"
	"reflect"
)

func PanicString(err string) {
	if err != "" {
		panic(errors.New(err))
	}
}
func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}
func PrintError(err error) {
	var _, printErr = fmt.Printf("%+v\n\n", err)
	PanicError(printErr)
}
func AssertNil(i any, message string) {
	//refer: https://glucn.medium.com/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	var ok = false
	if reflect.ValueOf(i).Kind() == reflect.Ptr {
		ok = reflect.ValueOf(i).IsNil()
	} else if i == nil {
		ok = true
	}
	if !ok {
		panic(message)
	}
}
