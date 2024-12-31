package main

import (
	"go-event-booking/db"
	"go-event-booking/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)
	
	server.Run(":8080")
}

