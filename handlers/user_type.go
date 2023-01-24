package handlers

import (
	"net/http"

	"github.com/andres15alvarez/ticket_manager/config"
	"github.com/andres15alvarez/ticket_manager/models"
	"github.com/gin-gonic/gin"
)

func ListUserTypeHandler(c *gin.Context) {
	db := config.GetDB()
	userTypes, err := models.GetAllUserTypes(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	if len(userTypes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No user types found"})
	} else {
		c.JSON(http.StatusOK, userTypes)
	}
}
