package http

import (
	"fmt"
	"net/http"
	"testing"
)

// This should be the first, otherwise originalGlobal will not be nil
func TestBeginTLSConfig(t *testing.T) {
	var originalGlobal = http.DefaultTransport.(*http.Transport).TLSClientConfig
	if originalGlobal != nil {
		t.Fatal("original global is not clean", originalGlobal)
	}
	var config = GetTLSConfigGlobal()
	var changedGlobal = http.DefaultTransport.(*http.Transport).TLSClientConfig
	if changedGlobal == nil {
		t.Fatalf("changed global should not be nil")
	}
	if config == nil {
		t.Fatalf("reponse config should not be nil")
	}
}

func TestGet(t *testing.T) {
	resp := Get("http://www.google.com", nil)
	fmt.Println(resp.Status)
}
