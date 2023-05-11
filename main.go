package main

import (
	"duducharapa/lyrics/api/handlers"
	"duducharapa/lyrics/database"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	database.Connect()

	route.GET("/lyrics", handlers.GetLyrics)
	route.POST("/lyrics", handlers.CreateLyric)

	route.GET("/users", handlers.ListUsers)
	route.POST("/users", handlers.RegisterUser)

	route.Run()
}
