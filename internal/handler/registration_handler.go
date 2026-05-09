package handler

import (
	"net/http"

	"exammple.com/event-booking-api/internal/domain"
	"github.com/gin-gonic/gin"
)

type RegistrationHandler struct {
	repo domain.EventRepository
}

func NewRegistrationHandler(repo domain.EventRepository) *RegistrationHandler {
	return &RegistrationHandler{repo: repo}
}

func (h *RegistrationHandler) RegisterUserForEvent(context *gin.Context) {
	eventId := context.Param("id")
	userId := context.GetInt64("userId")

	event, err := h.repo.GetById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = h.repo.RegisterUserForEvent(*event, userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "You are registered for the event."})
}

func (h *RegistrationHandler) GetRegistrationsForEvent(context *gin.Context) {
	eventId := context.Param("id")

	event, err := h.repo.GetById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	userIds, err := h.repo.GetRegistrationsForEvent(*event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch registrations."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"registrations": userIds})
}

func (h *RegistrationHandler) CancelUserEventRegistration(context *gin.Context) {
	eventId := context.Param("id")
	userId := context.GetInt64("userId")

	event, err := h.repo.GetById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = h.repo.CancelUserRegistrationForEvent(*event, userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "You cancelled you registration for the event."})
}
