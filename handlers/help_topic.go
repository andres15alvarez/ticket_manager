package handlers

import (
	"net/http"

	"github.com/andres15alvarez/ticket_manager/config"
	"github.com/andres15alvarez/ticket_manager/models"
	"github.com/gin-gonic/gin"
)

func ListHelpTopicHandler(c *gin.Context) {
	db := config.GetDB()
	helpTopics, err := models.GetAllHelpTopics(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	if len(helpTopics) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No help topics found"})
	} else {
		c.JSON(http.StatusOK, helpTopics)
	}
}
