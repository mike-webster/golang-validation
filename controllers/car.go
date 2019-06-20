package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mike-webster/golang-validation/models"
)

// carHandler will handle POST requests to /car
func carHandler(c *gin.Context) {
	var ret models.CarExample
	err := c.Bind(&ret)
	if err != nil {
		c.Set("controllerError", true)
		return
	}

	c.Status(200)
}
