package main

import (
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func main() {
	r := getRouter()
	r.Run("0.0.0.0:3001")
}

func getRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", pingHandler)
	return r
}

func pingHandler(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}
