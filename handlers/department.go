package handlers

import (
	"net/http"

	"github.com/andres15alvarez/ticket_manager/config"
	"github.com/andres15alvarez/ticket_manager/models"
	"github.com/gin-gonic/gin"
)

func ListDepartmentHandler(c *gin.Context) {
	db := config.GetDB()
	departments, err := models.GetAllDepartments(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	if len(departments) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No departments found"})
	} else {
		c.JSON(http.StatusOK, departments)
	}
}
