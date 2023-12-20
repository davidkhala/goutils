package goutils

import (
	"fmt"
	"github.com/kortschak/utter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtter(t *testing.T) {
	_map := map[string]string{}
	assert.Equal(t, "map[]", fmt.Sprint(_map))
	assert.Equal(t, "map[string]string{\n}\n", utter.Sdump(_map))

}
