package middlewares

import (
	"net/http"

	"exammple.com/event-booking-api/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	authHeaderToken := context.Request.Header.Get("Authorization")

	if authHeaderToken == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized Request"})
		return
	}

	userId, err := utils.VerifyToken(authHeaderToken)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
