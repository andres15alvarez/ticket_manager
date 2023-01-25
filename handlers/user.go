package handlers

import (
	"net/http"
	"strconv"

	"github.com/andres15alvarez/ticket_manager/config"
	"github.com/andres15alvarez/ticket_manager/models"
	"github.com/gin-gonic/gin"
)

func ListUserHandler(c *gin.Context) {
	db := config.GetDB()
	users, err := models.GetAllUsers(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Users not found"})
		return
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func GetUserHandler(c *gin.Context) {
	db := config.GetDB()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	user, err := models.GetUserByID(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else {
		c.JSON(http.StatusOK, user)
	}

}
