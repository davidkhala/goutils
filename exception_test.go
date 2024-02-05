package goutils

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"runtime/debug"
	"testing"
)

func TestPanicError(t *testing.T) {
	t.Run("panic nil", func(t *testing.T) {
		defer Deferred(func(err error, params ...interface{}) (success bool) {
			t.Fatal("panic(nil) should not trigger handler")
			return false
		})
		panic(nil)

	})
	t.Run("panic error", func(t *testing.T) {
		var err = errors.New("any")
		defer Deferred(func(err error, params ...interface{}) (success bool) {
			debug.PrintStack()
			return true
		})
		PanicError(err)
	})
	t.Run("panic string", func(t *testing.T) {

		defer Deferred(func(err error, params ...interface{}) (success bool) {
			assert.Equal(t, "str", err.Error())
			return true
		})
		panic("str")
	})
	t.Run("panic number: will not be handled ", func(t *testing.T) {
		var panic1 = func() {
			defer Deferred(func(err error, params ...interface{}) (success bool) {
				t.Fatal("panic(1) should not trigger handler")
				return false
			})
			panic(1)
		}
		assert.PanicsWithValue(t, 1, panic1)

	})
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
