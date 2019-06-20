package controllers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

// mwParseValidation will parse the gross default error messages into
// readable, nice messages we can display.
func mwParseValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		_, exists := c.Get("controllerError")
		if exists {
			ret := map[string]string{}
			for _, e := range c.Errors {
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

// mwLogBody just prints out the posted body for each request
// this helps troubleshoot failing tests - we wouldn't need this
// in a production env.
func mwLogBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		bs, _ := c.Request.GetBody()
		bytes, _ := ioutil.ReadAll(bs)
		log.Println("BODY: ", string(bytes))
		c.Next()
	}
}
