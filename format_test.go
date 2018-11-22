package goutils

import (
	"fmt"
	"testing"
	"time"
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

func TestInt2Float(t *testing.T) {
	var amount = "1"
	var amountFloat float64
	amountFloat = float64(Atoi(amount))
	fmt.Println(amountFloat)
}
func TestIntToByte(t *testing.T) {
	var i int64 = 1
	fmt.Println(byte(i))
	i = 256
	fmt.Println(byte(i))
	i = -1
	fmt.Println(byte(i))
}
func TestItoRunes(t *testing.T) {
	const charSpace = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var runes = []rune(charSpace)
	var result = ItoRunes(50, runes)
	fmt.Println(result)
}

func TestTimeFormat(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format("20060102")) //correct
}

func TestPaddingLeft(t *testing.T) {
	var str = "123"
	var result = PaddingLeft(str, 7, '0')
	fmt.Println(result)
}
