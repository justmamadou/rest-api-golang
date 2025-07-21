package main

import (
	"github.com/gin-gonic/gin"
	"github.com/justmamadou/rest-api-golang/db"
	"github.com/justmamadou/rest-api-golang/models"
	"net/http"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(c *gin.Context) {
	var events = models.GetAllEvents()
	c.JSONP(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()
	c.JSON(http.StatusCreated, gin.H{"Message": "Successfully created event", "event": event})
}
