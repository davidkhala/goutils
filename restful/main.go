package main

import (
	"github.com/davidkhala/goutils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title go-swagger
// @version 1.16
// @contact.email david-khala@hotmail.com
func main() {
	app := Run(true)
	app.GET("/ping", Ping)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // refers to /swagger/*any

	goutils.PanicError(app.Run(":8080"))
}
