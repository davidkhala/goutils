package main

import (
	"github.com/davidkhala/goutils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @title go-swagger
// @version 1.16
// @contact.email david-khala@hotmail.com
func main() {
	app := Run(true)

	// Ping test
	app.GET("/ping", Ping)

	// Get user value
	app.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := app.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := app.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	goutils.PanicError(app.Run(":8080"))
}
