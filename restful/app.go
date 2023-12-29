package restful

import (
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
