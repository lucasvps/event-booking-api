package routes

import (
	"exammple.com/event-booking-api/internal/handler"
	"exammple.com/event-booking-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, userHandler *handler.UserHandler, eventHandler *handler.EventHandler, registrationHandler *handler.RegistrationHandler) {
	// MIDDLEWARE GROUP
	authenticatedRoute := server.Group("/")
	authenticatedRoute.Use(middlewares.Authenticate)

	// AUTH ROUTES
	server.POST("/login", userHandler.Login)

	// USER ROUTES
	server.POST("/signup", userHandler.SignupUser)
	server.GET("/users", userHandler.GetUsers)

	// REGISTRATION ROUTES
	server.GET("/events/:id/register", registrationHandler.GetRegistrationsForEvent)
	authenticatedRoute.POST("/events/:id/register", registrationHandler.RegisterUserForEvent)
	authenticatedRoute.DELETE("/events/:id/register", registrationHandler.CancelUserEventRegistration)

	// EVENTS ROUTES
	server.GET("/events", eventHandler.GetEvents)
	server.GET("/events/:id", eventHandler.GetEventById)

	authenticatedRoute.POST("/events", eventHandler.CreateEvent)
	authenticatedRoute.DELETE("/events/:id", eventHandler.DeleteEvent)
	authenticatedRoute.PUT("/events/:id", eventHandler.UpdateEvent)
}
