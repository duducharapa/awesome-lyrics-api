package database

import (
	"duducharapa/lyrics/api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect with database")
	}

	database.AutoMigrate(&models.Lyric{})
	database.AutoMigrate(&models.User{})

	DB = database
}
