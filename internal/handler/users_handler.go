package handler

import (
	"net/http"

	domain "exammple.com/event-booking-api/internal/domain"
	"exammple.com/event-booking-api/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo domain.UserRepository
}

func NewUserHandler(repo domain.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) SignupUser(context *gin.Context) {
	var user domain.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = h.repo.Signup(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Signup successfully."})
}

func (h *UserHandler) GetUsers(context *gin.Context) {
	users, err := h.repo.GetAll()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch users"})
	}

	context.JSON(http.StatusOK, gin.H{"data": users})
}

func (h *UserHandler) Login(context *gin.Context) {
	var user domain.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = h.repo.Login(&user)

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
