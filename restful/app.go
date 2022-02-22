package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// Ping
// @Success 200 {string} pong
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func Run(coloredConsole bool) *gin.Engine {
	var ci string
	ci = os.Getenv("CI")
	switch ci {
	case "":
		gin.SetMode(gin.DebugMode)
	case "true":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	if !coloredConsole {
		// Disable Console Color
		gin.DisableConsoleColor()
	}

	r := gin.Default()

	return r
}
