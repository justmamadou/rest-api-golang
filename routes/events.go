package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/justmamadou/rest-api-golang/models"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve events"})
		return
	}
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	userId := c.GetInt64("userId")

	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
		return
	}

	event.UserID = userId
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Message": "Successfully created event", "event": event})
}

func getEvent(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve event"})
		return
	}
	c.JSON(http.StatusOK, event)
}

func updateEvent(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	userId := c.GetInt64("userId")
	event, err := models.GetEventByID(eventID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve event"})
		return
	}

	if event.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this event"})
		return
	}

	var updateEvent models.Event
	err = c.ShouldBindJSON(&updateEvent)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
		return
	}

	updateEvent.ID = eventID
	err = updateEvent.UpdateEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Successfully updated event", "event": updateEvent})

}

func deleteEvent(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	userId := c.GetInt64("userId")
	event, err := models.GetEventByID(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve event"})
		return
	}

	if event.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete this event"})
		return
	}

	err = models.DeleteEvent(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Successfully deleted event"})
}
