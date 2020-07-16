package http

import (
	"fmt"
	"github.com/davidkhala/goutils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	resp := Get("http://www.google.com", nil)
	fmt.Println(string(resp.BodyBytes()))
}
func TestTimeout(t *testing.T) {
	var invalidUrl = "http://abc.google.com"
	var handler = func(errString string, params ...interface{}) (success bool) {
		return assert.Regexp(t, "^Get "+invalidUrl+": dial tcp: lookup ", errString)
	}
	defer goutils.Deferred(handler)
	Get(invalidUrl, nil)

}
func TestBeginTLSConfig(t *testing.T) {
	var originalGlobal = http.DefaultTransport.(*http.Transport).TLSClientConfig
	if originalGlobal != nil {
		t.Fatal("original global is not clean", originalGlobal)
	}
	var config = BeginTLSConfig()
	var changedGlobal = http.DefaultTransport.(*http.Transport).TLSClientConfig
	if changedGlobal == nil {
		t.Fatalf("changed global should not be nil")
	}
	if config == nil {
		t.Fatalf("reponse config should not be nil")
	}
}
