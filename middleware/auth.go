package middleware

import (
	"go-event-booking/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	
	if token == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "message": "unauthorized" })
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{ "message": err })
		return
	}

	context.Set("userId", userId)
	context.Next()
}