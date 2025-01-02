package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mkashifaslam/golang/rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	// events routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// auth middleware
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	// auth routes
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	// users routes
	server.POST("/signup", signup)
	server.POST("/login", login)
}
