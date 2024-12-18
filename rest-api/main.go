package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mkashifaslam/golang/rest-api/db"
	"github.com/mkashifaslam/golang/rest-api/routes"
	"log"
)

var logger = log.Default()

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		logger.Println(err)
	}
}
