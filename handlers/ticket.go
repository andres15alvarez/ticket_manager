package handlers

import (
	"net/http"
	"strconv"

	"github.com/andres15alvarez/ticket_manager/config"
	"github.com/andres15alvarez/ticket_manager/models"
	"github.com/gin-gonic/gin"
)

func GetTicketHandler(c *gin.Context) {
	db := config.GetDB()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	ticket, err := models.GetTicketByID(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if ticket == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	} else {
		c.JSON(http.StatusOK, ticket)
	}

}

func CreateTicketHandler(c *gin.Context) {
	db := config.GetDB()
	var ticket models.Ticket
	if err := c.BindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ticketCreated, err := models.CreateTicket(db, &ticket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ticketCreated)
}

func ListTicketHandler(c *gin.Context) {
	db := config.GetDB()
	var tickets []*models.Ticket
	var err error
	stateID, err := strconv.ParseInt(c.Query("state"), 10, 64)
	if stateID == 0 {
		tickets, err = models.GetAllTickets(db)
	} else {
		tickets, err = models.GetTicketsByState(db, stateID)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tickets)
}
