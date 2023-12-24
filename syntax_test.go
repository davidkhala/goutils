package goutils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

var Map = map[string]string{}

func TestRange(t *testing.T) {
	Map["abc"] = "1"
	for key, value := range Map {
		fmt.Println(key, value)
	}
}
func TestReflect(t *testing.T) {
	t.Run("get name of type", func(t *testing.T) {

		type Start struct {
		}
		var name = GetType(Start{})
		assert.Equal(t, "Start", name)
	})
}

func TestInf(t *testing.T) {
	var positiveInf = math.Inf(1)
	assert.Equal(t, "+Inf", fmt.Sprint(positiveInf))
	assert.True(t, 1.00-positiveInf < 0)

}
func TestSlice(t *testing.T) {
	var array = []int{1, 2, 3}
	assert.Equal(t, "[2 3]", fmt.Sprint(append([]int{}, array[1:]...)))
}

func TestNil(t *testing.T) {
	var arrays []string
	arrays = nil
	assert.Equal(t, 0, len(arrays))
}
func TestPath(t *testing.T) {
	fmt.Println(HomeResolve("delphi-fabric", "config"))
}
