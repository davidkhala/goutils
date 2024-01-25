package testutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func AssertSuccess(t *testing.T, err error, msgAndArgs ...interface{}) {
	assert.NoError(t, err, msgAndArgs...)
}
func AssertError(t *testing.T, err error, expectedErrString string, msgAndArgs ...interface{}) {
	assert.EqualError(t, err, expectedErrString, msgAndArgs...)
}
func AssertOK(t *testing.T, value bool, msgAndArgs ...interface{}) {
	assert.True(t, value, msgAndArgs...)
}
