package main

import (
	"context"
	http2 "github.com/davidkhala/goutils/http"
	"github.com/kortschak/utter"
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
		// TODO WIP
		var response = http2.Get("http://localhost:8080/panic", nil)
		utter.Dump(response.Trim())
		var trimmed = response.Trim()
		//assert.Equal(t, trimmed.Body, "error")
		assert.Equal(t, trimmed.StatusCode, http.StatusInternalServerError)
	})

}
