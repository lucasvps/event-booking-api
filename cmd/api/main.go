package main

import (
	"exammple.com/event-booking-api/db"
	"exammple.com/event-booking-api/internal/handler"
	"exammple.com/event-booking-api/internal/repository"
	"exammple.com/event-booking-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database := db.InitDB()

	eventRepository := repository.NewEventRepository(database)
	userRepository := repository.NewUserRepository(database)

	userHandler := handler.NewUserHandler(userRepository)
	eventHandler := handler.NewEventHandler(eventRepository)
	registrationHandler := handler.NewRegistrationHandler(eventRepository)

	server := gin.Default()

	routes.RegisterRoutes(server, userHandler, eventHandler, registrationHandler)

	server.Run(":8080")
}
