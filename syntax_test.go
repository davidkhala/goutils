package goutils

import (
	"fmt"
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
