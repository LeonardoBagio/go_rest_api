package main

import (
	"example.com/rest-api/database"
	"example.com/rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	server := gin.Default()

	routes.RegisterEventRoutes(server)
	routes.RegisterUserRoutes(server)

	server.Run(":8080")
}
