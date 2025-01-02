package StoriesController

import (
	StoriesService "GoApp/Service/Stories"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetStories(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, StoriesService.GetAllStories())
}

func AddLike(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, StoriesService.AddLikeToStory(id))
}

func GetStoryTree(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, StoriesService.GetStoryTreeByStoryId(id))
}
