package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mkashifaslam/golang/rest-api/models"
	"net/http"
)

func signup(context *gin.Context) {
	logger.Println("Call Signup Route")

	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		logger.Println("[User] RequestError:", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})

		return
	}

	err = user.Save()
	if err != nil {
		logger.Println("[User] SaveError:", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save user",
		})
		return
	}

	logger.Println("New User", user)
	context.JSON(http.StatusCreated, gin.H{
		"message": "User successfully created",
	})
}

func login(context *gin.Context) {
	logger.Println("Call login Route")
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		logger.Println("[User] RequestError:", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		logger.Println("[User] LoginError:", err)
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful!",
	})
}
