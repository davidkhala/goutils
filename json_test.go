package goutils

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/kortschak/utter"
	"github.com/stretchr/testify/assert"
	"testing"
)

type URL struct {
	Url string `json:"url"`
}

type TITLE struct {
	Title string `json:"title"`
}

func TestFromJson(t *testing.T) {

	t.Run("compond structure: case 1", func(t *testing.T) {
		type compound struct {
			URL
			TITLE
		}
		var composite compound
		FromJson([]byte(`{  "url": "www.google.com",  "title": "Google"}`), &composite)
		fmt.Println(composite)
		fmt.Println(string(ToJson(composite)))
		assert.Panics(t, func() {
			FromJson(nil, &composite)
		}, "should panic from nil json bytes")
	})

	t.Run("compound structure", func(t *testing.T) {
		type TITLE struct {
			Title string
		}
		type compound struct {
			URL
			TITLE
		}
		var composite = compound{URL{"www.google.com"}, TITLE{"Google"}}

		fmt.Println(composite)
		fmt.Println(string(ToJson(composite)))
		assert.Panics(t, func() {
			FromJson(nil, &composite)
		}, "should panic from nil json bytes")
	})
}
func TestFromJson2(t *testing.T) {

	t.Run("json array", func(t *testing.T) {
		var arr []string
		FromJson([]byte(`["whoami"]`), &arr)
		assert.Equal(t, `[whoami]`, spew.Sprint(arr))
	})
}

func TestToJson(t *testing.T) {
	t.Run("url array", func(t *testing.T) {
		var urls = []URL{{"www.google.com"}, {"facebook"}}
		var jsonBytes = ToJson(urls)
		fmt.Println(string(jsonBytes))
	})
	t.Run("string array", func(t *testing.T) {
		var stringArray = []string{
			"whoami",
		}
		var jsonBytes = ToJson(stringArray)
		fmt.Println(string(jsonBytes))
	})
	t.Run("without structure", func(t *testing.T) {
		var obj = map[string]interface{}{
			"a": 1,
			"b": "buffer",
		}
		var jsonBytes = ToJson(obj)
		fmt.Println(string(jsonBytes))
	})
	t.Run("auto loading", func(t *testing.T) {
		var _map = map[string]string{
			"a": Base64Encode([]byte("b1234")),
			"c": Base64Encode([]byte("d1234")),
		}
		var byteMap = map[string][]byte{}
		var _json = ToJson(_map)
		FromJson(_json, &byteMap)
		utter.Dump(byteMap)
		for _, value := range byteMap {
			println(string(value))
		}
	})

}
