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

// Context
// @Router /context/{key} [get]
// @Param key	path string true "context key"
// @Success 200 {string} string
// @Failure 404 {string} string
// @Produce text/plain
func Context(c *gin.Context) {
	var key = c.Param("key")
	value := c.GetString(key)

	if value == "" {
		c.String(http.StatusNotFound, key)
	} else {
		c.String(http.StatusOK, value)
	}
}
func AnyContext(c *gin.Context) {
	var key = c.Param("key")
	value, exists := c.Get(key)
	if exists {
		c.String(http.StatusOK, string(goutils.ToJson(value)))
	} else {
		c.String(http.StatusNotFound, key)
	}
}

// Panic
// @Router /panic/{error} [get]
// @Param error path string true "the error message to be replied back in response"
// @Failure 500 {string} string
func Panic(c *gin.Context) {
	defer goutils.Deferred(func(err error, params ...interface{}) (success bool) {
		c.String(http.StatusInternalServerError, err.Error())
		return true
	})
	var errString = c.Param("error")
	panic(errString)
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

	if coloredConsole {
		gin.ForceConsoleColor()
	} else {
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
