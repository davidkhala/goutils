package restful

import (
	"github.com/gin-gonic/gin"
	"os"
)

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
