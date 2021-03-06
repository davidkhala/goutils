package goutils

import (
	"fmt"
	"testing"
)

type URL struct {
	Url string `json:"url"`
}

type TITLE struct {
	Title string `json:"title"`
}

func TestFromJson(t *testing.T) {

	type compound struct {
		URL
		TITLE
	}
	var composite compound
	FromJson([]byte(`{  "url": "www.google.com",  "title": "Google"}`), &composite)
	fmt.Println(composite)
	fmt.Println(string(ToJson(composite)))
}
func TestFromJson2(t *testing.T) {

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

}
func TestToJson(t *testing.T) {
	var urls = []URL{{"www.google.com"}, {"facebook"}}
	var jsonBytes = ToJson(urls)
	fmt.Println(string(jsonBytes))
}
