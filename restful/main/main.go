package main

import (
	"github.com/davidkhala/goutils"
	"github.com/davidkhala/goutils/restful"
	_ "github.com/davidkhala/goutils/restful/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

// @title go-swagger
// @version v0.0.1
// @contact.email david-khala@hotmail.com
func main() {
	app := restful.App(true)
	app.StaticFile("/favicon.ico", "./favicon.ico")
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8080"
	}
	restful.SetContext(map[string]any{
		port: port,
	}, app)
	app.GET("/ping", restful.Ping)
	app.GET("/", restful.Ping)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // refers to /swagger/*any

	goutils.PanicError(app.Run(":" + port))
}
