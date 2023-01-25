package handlers

import (
	"net/http"

	"github.com/andres15alvarez/ticket_manager/config"
	"github.com/andres15alvarez/ticket_manager/models"
	"github.com/gin-gonic/gin"
)

func ListStateHandler(c *gin.Context) {
	db := config.GetDB()
	states, err := models.GetAllStates(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if len(states) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No states found"})
		return
	} else {
		c.JSON(http.StatusOK, states)
	}
}
