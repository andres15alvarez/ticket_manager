package handlers

import (
	"github.com/andres15alvarez/ticket_manager/config"
	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"environment": config.InitializeConfig().Environment,
	})
}
