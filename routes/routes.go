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

	authenticatedRegistrationRoutes := server.Group("events/:id/registrations", middleware.Authenticate)

	authenticatedRegistrationRoutes.POST("/", createRegistration)
	authenticatedRegistrationRoutes.GET("/", getAllRegistrations)
	
	server.POST("/signup", signup)
	server.POST("/login", login)
}