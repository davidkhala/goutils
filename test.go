package goutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func AssertSuccess(t *testing.T, err error) {
	assert.Empty(t, err)
}
