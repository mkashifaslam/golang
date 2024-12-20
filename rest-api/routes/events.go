package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mkashifaslam/golang/rest-api/models"
	"github.com/mkashifaslam/golang/rest-api/utils"
	"log"
	"net/http"
	"strconv"
)

var logger = log.Default()

func getEvents(context *gin.Context) {
	logger.Println("Call GetEvents Route")
	events, err := models.GetAllEvents()
	if err != nil {
		logger.Println("[Event] RequestError:", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch events. try again later",
		})
		return
	}

	logger.Println("Events:", events)
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	logger.Println("Call GetEvent Route")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	logger.Println("[Event] RequestID:", eventId)
	if err != nil {
		logger.Println("[Event] RequestError:", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse event id",
		})
		return
	}
	var event *models.Event
	event, err = models.GetEventByID(eventId)
	if err != nil {
		logger.Println("[Event] GetError:", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not get event",
		})
		return
	}

	logger.Println("Event:", event)
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	logger.Println("Call CreateEvent Route")
	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		logger.Println("[Event] RequestError:", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	event.UserID = userId
	err = event.Save()
	if err != nil {
		logger.Println("[Event] SaveError:", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save event. try again later",
		})
		return
	}

	logger.Println("New Event", event)
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}

func updateEvent(context *gin.Context) {
	logger.Println("Call UpdateEvent Route")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	logger.Println("[Event] RequestID:", eventId)
	if err != nil {
		logger.Println("[Event] RequestError:", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse event id",
		})
		return
	}
	_, err = models.GetEventByID(eventId)
	if err != nil {
		logger.Println("[Event] GetError:", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event",
		})
		return
	}

	var updatedEvent *models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		logger.Println("[Event] RequestError:", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		logger.Println("[Event] UpdateError:", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update event",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
	})

}

func deleteEvent(context *gin.Context) {
	logger.Println("Call DeleteEvent Route")
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
		logger.Println("[Event] GetError:", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event",
		})
		return
	}
	err = event.Delete()
	if err != nil {
		logger.Println("[Event] DeleteError:", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete event",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}
