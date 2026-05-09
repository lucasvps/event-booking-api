package main

import (
	"exammple.com/event-booking-api/db"
	"exammple.com/event-booking-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
