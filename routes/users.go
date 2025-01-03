package routes

import (
	"go-event-booking/models"
	"go-event-booking/utils"
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

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "message": "bad request - email and password are mandatory fields" })
		return
	}

	err = user.ValidateUser()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "message": "bad request - invalid email or password" })
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": "internal server error - failed to generate auth token" })
		return
	}

	context.JSON(http.StatusOK, gin.H{ "message": "request successful", "token": token })
}