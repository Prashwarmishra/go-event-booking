package routes

import (
	"go-event-booking/models"
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
	var event models.Event
	err := context.ShouldBindJSON(&event)
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	
	event.UserID = context.GetInt64("userId")

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": err })
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "request successful", "data": event})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad request - invalid event id"})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	userId := context.GetInt64("userId")

	if userId != event.UserID {
		context.JSON(http.StatusBadRequest, gin.H{ "message": "unauthorized" })
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
		context.JSON(http.StatusInternalServerError, gin.H{ "message": err })
		return
	}

	userId := context.GetInt64("userId")

	if userId != event.UserID {
		context.JSON(http.StatusBadRequest, gin.H{ "message": "unauthorized" })
		return
	}

	err = event.DeleteEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": "internal server error" })
		return
	}

	context.JSON(http.StatusOK, gin.H{ "message": "request successful", "data": event })

}