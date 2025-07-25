package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/justmamadou/rest-api-golang/middlerwares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/:id", getEvent)

	authenticatedGroup := server.Group("/events").Use(middlerwares.Authentication)

	authenticatedGroup.POST("/", createEvent)
	authenticatedGroup.PUT("/:id", updateEvent)
	authenticatedGroup.DELETE("/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
