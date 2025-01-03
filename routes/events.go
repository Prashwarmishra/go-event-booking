package routes

import (
	"go-event-booking/models"
	"go-event-booking/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": err })
		return
	} 

	context.JSON(http.StatusOK, gin.H{"message": "request successful", "events": events})
}

func getEventById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "message": "bad request - invalid event id" })
		return
	}

	data, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": "internal server error" })
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "request successful", "event": data})
}

func createEvent(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{ "message": "unauthorized" })
		return
	}

	err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{ "message": "invalid token" })
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	
	event.UserID = 1

	data, err := event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": err })
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "request successful", "data": data})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad request - invalid event id"})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	updatedEvent.ID = event.ID
	updatedEvent.UserID = event.UserID

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad request - unable to parse body"})
		return
	}

	err = updatedEvent.UpdateEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{ "message": "request successful", "data": updatedEvent })
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "message": "bad request - invalid event id" })
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": "internal server error" })
		return
	}

	err = event.DeleteEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": "internal server error" })
		return
	}

	context.JSON(http.StatusOK, gin.H{ "message": "request successful", "data": event })

}