package main

import (
	"github.com/davidkhala/goutils/restful"
	_ "github.com/davidkhala/goutils/restful/docs"
)

// @title go-swagger
// @version v0.0.1
// @contact.email david-khala@hotmail.com
func main() {
	app, run := restful.SampleApp(8080)

	app.GET("/", restful.Ping)
	app.GET("/context/:key", restful.Context)
	app.GET("/panic/:error", restful.Panic)
	run()

}
