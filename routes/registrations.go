package routes

import (
	"go-event-booking/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAllRegistrations(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "message": "bad request - failed to parse event id" })
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "message": "bad request - event not found" })
		return
	}

	registrations, err := event.GetAllRegistrations()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": "internal server error - failed to fetch registrations" })
		return
	}

	context.JSON(http.StatusOK, gin.H{ "message": "request successful", "data": registrations })
}

func createRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventsId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "message": "bad request - failed to parse event id" })
		return
	}

	event, err := models.GetEventById(eventsId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "message": "bad request - event not found" })
		return
	}

	err = event.CreateRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": "failed to create registation" })
		return
	}

	context.JSON(http.StatusOK, gin.H{ "message": "user registration successful" })
}