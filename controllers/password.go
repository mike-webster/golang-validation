package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mike-webster/golang-validation/models"
)

func passwordHandler(c *gin.Context) {
	var password models.PasswordExample
	err := c.Bind(&password)
	if err != nil {
		c.Set("controllerError", true)
		return
	}

	c.Status(200)
}
