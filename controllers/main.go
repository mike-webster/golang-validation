package controllers

import "github.com/gin-gonic/gin"

var router *gin.Engine

// GetRouter will return a configured router
func GetRouter() *gin.Engine {
	if router != nil {
		// I'm only cacheing this to avoid recreating the router
		// a bunch of times for tests
		return router
	}
	r := gin.Default()
	r.Use(mwLogBody())
	r.Use(mwParseValidation())
	r.POST("/car", carHandler)
	r.POST("/album", albumHandler)
	r.POST("/password", passwordHandler)
	r.POST("/lead", leadHandler)
	router = r
	return router
}
