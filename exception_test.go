package goutils

import (
	"errors"
	"runtime/debug"
	"testing"
)

var err = errors.New("any")

func TestPanicError(t *testing.T) {
	var handler = func(errString string, params ...interface{}) (success bool) {
		debug.PrintStack()

		return true
	}
	defer Deferred(handler)

	PanicError(err)

}
