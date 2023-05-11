package handlers

import (
	"duducharapa/lyrics/api/dtos"
	"duducharapa/lyrics/api/models"
	"duducharapa/lyrics/api/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLyrics(context *gin.Context) {
	var lyrics []models.Lyric

	if err := repositories.FindLyrics(&lyrics); err != nil {
		context.JSON(http.StatusInternalServerError, nil)
		return
	}

	context.JSON(http.StatusOK, lyrics)
}

func CreateLyric(context *gin.Context) {
	var createDTO dtos.CreateLyricDTO
	var lyric models.Lyric

	if err := context.ShouldBindJSON(&createDTO); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Bad request data"})
		return
	}

	if err := repositories.CreateLyric(&createDTO, &lyric); err != nil {
		context.JSON(http.StatusInternalServerError, nil)
		return
	}

	context.JSON(http.StatusCreated, lyric)
}
