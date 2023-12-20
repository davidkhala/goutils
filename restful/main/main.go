package main

import (
	"github.com/davidkhala/goutils"
	"github.com/davidkhala/goutils/restful"
	_ "github.com/davidkhala/goutils/restful/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title go-swagger
// @version 1.16
// @contact.email david-khala@hotmail.com
func main() {
	app := restful.App(true)
	app.GET("/ping", restful.Ping)
	app.GET("/", restful.Ping)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // refers to /swagger/*any

	goutils.PanicError(app.Run(":8080"))
}
