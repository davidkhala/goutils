package restful

import (
	"github.com/davidkhala/goutils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"strconv"
)

func SampleApp(defaultPort int) (*gin.Engine, func()) {
	app := App(true)
	app.StaticFile("/favicon.ico", "./favicon.ico")
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = strconv.Itoa(defaultPort)
	}
	SetContext(map[string]any{
		"PORT": port,
	}, app)
	app.GET("/ping", Ping)
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // refers to /swagger/*any
	return app, func() {
		goutils.PanicError(app.Run(":" + port))
	}
}
