package routes

import (
	"go-event-booking/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticatedEventRoutes := server.Group("/events", middleware.Authenticate)

	authenticatedEventRoutes.GET("/",  getEvents)
	authenticatedEventRoutes.GET("/:id", getEventById)
	authenticatedEventRoutes.POST("/", createEvent)
	authenticatedEventRoutes.PUT("/:id", updateEvent)
	authenticatedEventRoutes.DELETE("/:id", deleteEvent)
	
	server.POST("/signup", signup)
	server.POST("/login", login)
}