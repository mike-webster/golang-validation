package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mike-webster/golang-validation/models"
)

// leadHandler will handle POST requests to /lead
func leadHandler(c *gin.Context) {
	var lead models.LeadSourceExample
	err := c.Bind(&lead)
	if err != nil {
		c.Set("controllerError", true)
		return
	}

	c.Status(200)
}
