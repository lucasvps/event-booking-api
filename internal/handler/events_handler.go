package handler

import (
	"fmt"
	"net/http"

	"exammple.com/event-booking-api/internal/domain"
	repository "exammple.com/event-booking-api/internal/repository"
	"github.com/gin-gonic/gin"
)

func GetEvents(context *gin.Context) {
	events, err := repository.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events."})
		return
	}

	context.JSON(http.StatusOK, events)
}

func GetEventById(context *gin.Context) {
	id := context.Param("id")

	event, err := repository.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": event})
}

func CreateEvent(context *gin.Context) {
	var event domain.Event

	err := context.ShouldBindJSON(&event)

	userId := context.GetInt64("userId")

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.UserID = userId

	err = repository.Save(&event)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create the event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func DeleteEvent(context *gin.Context) {
	id := context.Param("id")
	userId := context.GetInt64("userId")

	eventById, err := repository.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	if eventById.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to delete this event."})
		return
	}

	err = repository.DeleteEvent(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "An error ocurred. The event could not be deleted."})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"message": "The event was deleted successfully."})
}

func UpdateEvent(context *gin.Context) {

	userId := context.GetInt64("userId")
	id := context.Param("id")

	eventById, err := repository.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	if eventById.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to update this event."})
		return
	}

	var event domain.Event
	err = context.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = eventById.ID
	event.UserID = eventById.UserID

	err = repository.UpdateEvent(id, event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "An error ocurred. The event could not be updated."})
	}

	context.JSON(http.StatusOK, gin.H{"event": event})
}
