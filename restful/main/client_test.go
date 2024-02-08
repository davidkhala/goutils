package main

import (
	"context"
	http2 "github.com/davidkhala/goutils/http"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestMain(m *testing.M) {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		main()
	}(ctx)
	m.Run()
	cancel()
}
func TestGet(t *testing.T) {
	t.Run("", func(t *testing.T) {
		response := http2.Get("http://localhost:8080", nil)
		assert.Equal(t, response.Trim().Body, "pong")
	})
	t.Run("ping", func(t *testing.T) {
		response := http2.Get("http://localhost:8080/ping", nil)
		assert.Equal(t, response.Trim().Body, "pong")
	})
	t.Run("panic", func(t *testing.T) {
		var errString = "errorMessage"
		var response = http2.Get("http://localhost:8080/panic/"+errString, nil)
		var trimmed = response.Trim()
		assert.Equal(t, errString, trimmed.Body)
		assert.Equal(t, http.StatusInternalServerError, trimmed.StatusCode)
	})
	t.Run("context getter", func(t *testing.T) {
		var response = http2.Get("http://localhost:8080/context/PORT", nil)
		var trimmed = response.Trim()
		assert.Equal(t, "8080", trimmed.Body)
		assert.Equal(t, http.StatusOK, trimmed.StatusCode)
		response = http2.Get("http://localhost:8080/context/any", nil)
		trimmed = response.Trim()
		assert.Equal(t, "any", trimmed.Body)
		assert.Equal(t, http.StatusNotFound, trimmed.StatusCode)
	})

}
