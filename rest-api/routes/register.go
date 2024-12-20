package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mkashifaslam/golang/rest-api/models"
	"net/http"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	logger.Println("[Event] RequestID:", eventId)
	if err != nil {
		logger.Println("[Event] RequestError:", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse event id",
		})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch event",
		})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not register user for event",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}

func cancelRegistration(context *gin.Context) {}
