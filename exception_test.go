package goutils

import (
	"errors"
	"github.com/stretchr/testify/assert"
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
func TestAssert(t *testing.T) {
	type localType struct {
	}
	var nPointer *localType
	var transientMap map[string][]byte
	var intSlice []int
	var str string
	AssertNil(str, "string should be nil")
	AssertNil("", "empty string should be nil")
	AssertNil(transientMap, "transientMap should be nil")
	AssertNil(intSlice, "[]int should be nil")

	AssertNil(nPointer, "should be OK")
	AssertNil(nil, "should be OK")
	var msg = "should be nil"
	assert.PanicsWithValue(t, msg, func() {
		AssertNil("non-nil", msg)
	})
}
