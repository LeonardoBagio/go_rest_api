package main

import (
	"example.com/rest-api/database"
	"example.com/rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	server := gin.Default()

	routes.RegisterPublicRoutes(server)
	routes.RegisterAuthenticatedRoutes(server)

	server.Run(":8080")
}
