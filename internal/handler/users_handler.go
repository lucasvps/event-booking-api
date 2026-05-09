package handler

import (
	"net/http"

	models "exammple.com/event-booking-api/internal/domain"
	"exammple.com/event-booking-api/internal/repository"
	"exammple.com/event-booking-api/pkg/utils"
	"github.com/gin-gonic/gin"
)

func SignupUser(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = repository.SignupUser(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Signup successfully."})
}

func GetUsers(context *gin.Context) {
	users, err := repository.GetAllUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch users"})
	}

	context.JSON(http.StatusOK, gin.H{"data": users})
}

func Login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = repository.Login(&user)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"accessToken": token})
}
