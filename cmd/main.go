package main

import (
	StoriesController "GoApp/Controller/Stories"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

func init() {
	if err := godotenv.Load("prod.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/stories", StoriesController.GetStories)
	router.GET("/storyTree", StoriesController.GetStoryTree)

	router.Run("localhost:8087")
}
