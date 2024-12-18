package main

import (
	"github.com/mkashifaslam/golang/rest-api/db"
	"github.com/mkashifaslam/golang/rest-api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var logger = log.Default()

func main() {
	db.InitDB()

	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)

	err := server.Run(":8080")
	if err != nil {
		logger.Println(err)
	}
}

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
	logger.Println("Call CreateEvent Route")
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		logger.Println("[Event] RequestError:", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	event.UserID = 1
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
