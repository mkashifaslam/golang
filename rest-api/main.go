package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mkashifaslam/golang/rest-api/db"
	"github.com/mkashifaslam/golang/rest-api/models"
	"log"
	"net/http"
)

var logger = log.Default()

func main() {
	db.InitDB()

	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	err := server.Run(":8080")
	if err != nil {
		logger.Println(err)
	}
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch events. try again later",
		})
		return
	}

	logger.Println("Events:", events)
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		logger.Println("RequestError:", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	event.UserID = 1
	logger.Println("New Event Data", event)
	err = event.Save()
	if err != nil {
		logger.Println("SaveError:", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save event. try again later",
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}
