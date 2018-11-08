package goutils

import "errors"

func AssertEmpty(rest []byte, message string) {
	if rest != nil && len(rest) > 0 {
		PanicString(message)
	}
}

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