package handler

import (
	"net/http"

	"exammple.com/event-booking-api/internal/repository"
	"github.com/gin-gonic/gin"
)

func RegisterUserForEvent(context *gin.Context) {
	eventId := context.Param("id")
	userId := context.GetInt64("userId")

	event, err := repository.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = repository.RegisterUserForEvent(*event, userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "You are registered for the event."})
}

func GetRegistrationsForEvent(context *gin.Context) {
	eventId := context.Param("id")

	event, err := repository.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	userIds, err := repository.GetRegistrationsForEvent(*event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch registrations."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"registrations": userIds})
}

func CancelUserEventRegistration(context *gin.Context) {
	eventId := context.Param("id")
	userId := context.GetInt64("userId")

	event, err := repository.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = repository.CancelUserRegistrationForEvent(*event, userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "You cancelled you registration for the event."})
}
