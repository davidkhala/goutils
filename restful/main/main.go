//go:generate swag init -g main/main.go -o ../docs -d ../
package main

import (
	"github.com/davidkhala/goutils/restful"
	_ "github.com/davidkhala/goutils/restful/docs"
	"github.com/gin-gonic/gin"
)

// @title go-swagger
// @version v0.0.1
// @contact.email david-khala@hotmail.com
func main() {
	app, run := restful.SampleApp(8080)

	app.GET("/context/:key", Context)
	app.GET("/panic/:error", Panic)
	run()

}

// Context
// @Router /context/{key} [get]
// @Param key	path string true "context key"
// @Success 200 {string} string
// @Failure 404 {string} string
// @Produce text/plain
func Context(c *gin.Context) {
	restful.Context(c)
}

// Panic
// @Router /panic/{error} [get]
// @Param error path string true "the error message to be replied back in response"
// @Failure 500 {string} string
func Panic(c *gin.Context) {
	restful.Panic(c)
}
