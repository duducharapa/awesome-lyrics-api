package repositories

import (
	"duducharapa/lyrics/api/dtos"
	"duducharapa/lyrics/api/models"
	"duducharapa/lyrics/database"
)

func ConvertLyric(createDTO *dtos.CreateLyricDTO, lyric *models.Lyric) {
	lyric.Name = createDTO.Name
	lyric.Author = createDTO.Author
}

func CreateLyric(createDTO *dtos.CreateLyricDTO, lyric *models.Lyric) error {
	ConvertLyric(createDTO, lyric)

	err := database.DB.Create(lyric).Error

	return err
}

func FindLyrics(lyrics *[]models.Lyric) error {
	err := database.DB.Find(lyrics).Error

	return err
}
