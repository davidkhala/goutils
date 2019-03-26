package goutils

import (
	"fmt"
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
func TestMap2JSON(t *testing.T) {
	Map["abc"] = "2"
	fmt.Println(string(ToJson(Map)))
}
func TestInf(t *testing.T) {
	var positiveInf = math.Inf(1)
	fmt.Println(positiveInf)
	fmt.Println(1.00-positiveInf < 0)

}
func TestSlice(t *testing.T) {
	var array = []int{1, 2, 3}
	fmt.Println(append([]int{}, array[1:]...))
}
