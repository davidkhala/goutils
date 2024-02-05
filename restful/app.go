package restful

import (
	"github.com/davidkhala/goutils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// Ping
// @Router /ping [get]
// @Success 200 {string} string "pong"
// @Produce text/plain
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func Context(c *gin.Context) {
	var key = c.Param("key")
	value, exists := c.Get(key)
	if exists {
		c.String(http.StatusOK, string(goutils.ToJson(value)))
	} else {
		c.String(http.StatusNotFound, key)
	}
}
func Panic(c *gin.Context) {
	defer goutils.Deferred(func(errString string, params ...interface{}) (success bool) {
		c.String(http.StatusInternalServerError, errString)
		return true
	})
	panic("error")
}

func App(coloredConsole bool) *gin.Engine {
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

func SetContext(data map[string]any, app *gin.Engine) gin.HandlerFunc {
	var interceptor = func(context *gin.Context) {
		for key, value := range data {
			context.Set(key, value)
		}
	}
	if app != nil {
		app.Use(interceptor)
	}
	return interceptor

}
