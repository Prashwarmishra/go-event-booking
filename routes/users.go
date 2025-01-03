package routes

import (
	"go-event-booking/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "message": "bad request - invalid payload" })
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": "server error - failed to create user" })
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}