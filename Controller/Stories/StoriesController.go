package StoriesController

import (
	StoriesService "GoApp/Service/Stories"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetStories(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, StoriesService.GetAllStories())
}

func GetStoryTree(c *gin.Context) {
	id, err := strconv.Atoi(c.Request.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, StoriesService.GetStoryTreeByStoryId(id))
}
