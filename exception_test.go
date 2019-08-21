package goutils

import (
	"fmt"
	"github.com/pkg/errors"
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

func TestErrorsWrap(t *testing.T) {
	// wrap will add this message to stack
	// This should be used directly, not within a wrapper of goutils
	var wrapped = errors.Wrap(err, "errors.Wrap")
	fmt.Printf("%+v\n\n", wrapped)
	wrapped = errors.Wrap(err, "")
	fmt.Printf("%+v\n\n", wrapped)
}

func TestErrorsWithMessage(t *testing.T) {
	// withMessage will not add this message to stack
	var withMessages = errors.WithMessage(err, "errors.WithMessage")
	fmt.Printf("%+v\n\n", withMessages)
	withMessages = errors.WithMessage(err, "")
	fmt.Printf("%+v\n\n", withMessages)

	withMessages = errors.WithMessage(withMessages, "abc")
	fmt.Printf("%+v\n\n", withMessages)
}
