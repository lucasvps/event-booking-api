package routes

import (
	"exammple.com/event-booking-api/internal/handler"
	"exammple.com/event-booking-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// MIDDLEWARE GROUP
	authenticatedRoute := server.Group("/")
	authenticatedRoute.Use(middlewares.Authenticate)

	// EVENTS ROUTES
	server.GET("/events", handler.GetEvents)
	server.GET("/events/:id", handler.GetEventById)

	authenticatedRoute.POST("/events", handler.CreateEvent)
	authenticatedRoute.DELETE("/events/:id", handler.DeleteEvent)
	authenticatedRoute.PUT("/events/:id", handler.UpdateEvent)

	// REGISTRATION ROUTES
	server.GET("/events/:id/register", handler.GetRegistrationsForEvent)
	authenticatedRoute.POST("/events/:id/register", handler.RegisterUserForEvent)
	authenticatedRoute.DELETE("/events/:id/register", handler.CancelUserEventRegistration)

	// USER ROUTES
	server.POST("/signup", handler.SignupUser)
	server.GET("/users", handler.GetUsers)

	// AUTH ROUTES
	server.POST("/login", handler.Login)

}
