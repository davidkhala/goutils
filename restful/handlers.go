package restful

import (
	"github.com/davidkhala/goutils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping
// @Router /ping [get]
// @Success 200 {string} string "pong"
// @Produce text/plain
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

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

func Panic(c *gin.Context) {
	defer goutils.Deferred(func(err error, params ...interface{}) (success bool) {
		c.String(http.StatusInternalServerError, err.Error())
		return true
	})
	var errString = c.Param("error")
	panic(errString)
}

func OnError(handler func(error, int)) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Next()
		for index, e := range c.Errors {
			err := e.Err
			handler(err, index)
		}
	}

}
