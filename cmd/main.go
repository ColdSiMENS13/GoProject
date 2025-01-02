package main

import (
	StoriesController "GoApp/Controller/Stories"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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
	router.GET("/stroies/:id", StoriesController.GetStoryTree)
	router.GET("/stories/:id/like", StoriesController.AddLike)

	router.Run("localhost:8087")
}
