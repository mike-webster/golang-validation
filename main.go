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
	r.Use(mwParseValidation())
	r.POST("/car", carHandler)
	return r
}

// mwParseValidation will parse the gross default error messages into
// readable, nice messages we can display.
func mwParseValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		_, exists := c.Get("controllerError")
		if exists {
			ret := map[string]string{}
			for _, e := range c.Errors {
				log.Println("error: ", e)
				switch e.Type {
				case gin.ErrorTypeBind:
					helpful := e.Err.(validator.ValidationErrors)
					for _, err := range helpful {
						ret[err.Field] = ValidationErrorToText(err)
					}
				case gin.ErrorTypePrivate:
					ret["msg"] = e.Error()
				default:
					log.Println("what is this error? ", e.Error())
				}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, ret)
			return
		}
	}
}

// carHandler will handle POST requests to /car
func carHandler(c *gin.Context) {
	var ret CarExample
	err := c.Bind(&ret)
	if err != nil {
		c.Set("controllerError", true)
		return
	}

	c.Status(200)
}
}
