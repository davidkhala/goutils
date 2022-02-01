package main

import (
	"github.com/davidkhala/goutils"
	"testing"
)

func TestSetup(t *testing.T) {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	goutils.PanicError(r.Run(":8080"))
}
