package goutils

import (
	"errors"
	"fmt"
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
