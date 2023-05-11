package handlers

import (
	"duducharapa/lyrics/api/dtos"
	"duducharapa/lyrics/api/models"
	"duducharapa/lyrics/api/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var createDTO dtos.CreateUserDTO
	var user models.User

	if err := context.ShouldBindJSON(&createDTO); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Bad request data"})
		return
	}

	if err := repositories.CreateUser(&createDTO, &user); err != nil {
		context.JSON(http.StatusInternalServerError, nil)
		return
	}

	context.JSON(http.StatusCreated, user)
}

func ListUsers(context *gin.Context) {
	var users []models.User

	if err := repositories.FindUsers(&users); err != nil {
		context.JSON(http.StatusInternalServerError, nil)
		return
	}

	context.JSON(http.StatusOK, users)
}
