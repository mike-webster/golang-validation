package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mike-webster/golang-validation/models"
)

// albumHandler will handle POST requests to /album
func albumHandler(c *gin.Context) {
	var album models.AlbumExample
	err := c.Bind(&album)
	if err != nil {
		c.Set("controllerError", true)
		return
	}

	c.Status(200)
}
