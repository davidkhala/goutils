package goutils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

type URL struct {
	Url string `json:"url"`
}

func TestFromJson(t *testing.T) {

	type TITLE struct {
		Title string `json:"title"`
	}
	type compound struct {
		URL
		TITLE
	}
	var composite compound
	FromJson([]byte(`{  "url": "www.google.com",  "title": "Google"}`), &composite)
	fmt.Println(string(ToJson(composite)))
	assert.Panics(t,func(){
		FromJson(nil, &composite)
	},"should panic from nil json bytes")

}
func TestToJson(t *testing.T) {
	var urls = []URL{{"www.google.com"}, {"facebook"}}
	var jsonBytes = ToJson(urls)
	fmt.Println(string(jsonBytes))
}
