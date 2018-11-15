package http

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	resp := Get("http://www.google.com", nil)
	fmt.Println(string(resp.BodyBytes()))
}

