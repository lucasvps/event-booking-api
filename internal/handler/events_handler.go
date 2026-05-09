package handler

import (
	"fmt"
	"net/http"

	"exammple.com/event-booking-api/internal/domain"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	repo domain.EventRepository
}

func NewEventHandler(repo domain.EventRepository) *EventHandler {
	return &EventHandler{repo: repo}
}

func (h *EventHandler) GetEvents(context *gin.Context) {
	events, err := h.repo.GetAll()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events."})
		return
	}

	context.JSON(http.StatusOK, events)
}

func (h *EventHandler) GetEventById(context *gin.Context) {
	id := context.Param("id")

	event, err := h.repo.GetById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": event})
}

func (h *EventHandler) CreateEvent(context *gin.Context) {
	var event domain.Event

	err := context.ShouldBindJSON(&event)

	userId := context.GetInt64("userId")

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.UserID = userId

	err = h.repo.Save(&event)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create the event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func (h *EventHandler) DeleteEvent(context *gin.Context) {
	id := context.Param("id")
	userId := context.GetInt64("userId")

	eventById, err := h.repo.GetById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	if eventById.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to delete this event."})
		return
	}

	err = h.repo.Delete(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "An error ocurred. The event could not be deleted."})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"message": "The event was deleted successfully."})
}

func (h *EventHandler) UpdateEvent(context *gin.Context) {

	userId := context.GetInt64("userId")
	id := context.Param("id")

	eventById, err := h.repo.GetById(id)

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

	err = h.repo.Update(id, event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "An error ocurred. The event could not be updated."})
	}

	context.JSON(http.StatusOK, gin.H{"event": event})
}
