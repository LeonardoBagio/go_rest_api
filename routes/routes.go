package routes

import "github.com/gin-gonic/gin"

func RegisterEventRoutes(server *gin.Engine) {
	/// GET
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	/// POST
	server.POST("/events", createEvent)
	/// PUT
	server.PUT("/events/:id", updateEvent)
	/// DELETE
	server.DELETE("/events/:id", deleteEvent)
}

func RegisterUserRoutes(server *gin.Engine) {
	/// GET

	/// POST
	server.POST("/singup", singUp)
}
