package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event."})
		return
	}

	result, err := models.Register(event, userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register for event."})
		return
	}

	registrationId, err := result.LastInsertId()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch registration id."})
		return
	}

	registration, err := models.GetRegistrationById(registrationId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch registration."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully registered for event.", "registration": registration})
}

func cancelRegistration(context *gin.Context) {

}
