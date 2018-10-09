package goutils

import (
	"fmt"
	"testing"
)

func TestFloat(t *testing.T) {
	var amountStr = "123456789.12"
	var amountFloat = ParseFloat(amountStr)
	fmt.Println(amountFloat, FormatFloat(amountFloat, 2))

}

var randomBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func TestRandString(t *testing.T) {
	var result = RandString(12, randomBytes)
	fmt.Println(result, len(result))
}
