package main

import (
	"log"

	gin "github.com/gin-gonic/gin"
)

func main() {
	r := getRouter()
	r.Run("0.0.0.0:3001")
}

// getRouter creates the gin router and sets up the routes
func getRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/car", carHandler)
	return r
}

// mwParseValidation will parse the gross default error messages into
// readable, nice messages we can display.
func mwParseValidation(c gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO, use borrowed validation to set nicer errors
		c.Next()
	}
}

// carHandler will handle POST requests to /car
func carHandler(c *gin.Context) {
	var ret CarExample
	err := c.Bind(&ret)
	if err != nil {
		c.JSON(400, gin.H{"err": err})
		return
	}

	log.Println("car: ", ret)
}
