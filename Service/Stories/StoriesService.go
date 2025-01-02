package StoriesService

import (
	config "GoApp/databaseConf/database"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Stories struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       *string `json:"image,omitempty"`
	Likes       int     `json:"likes"`
	Views       int     `json:"views"`
}

type StoryTree struct {
	Id          int          `json:"id"`
	StoryId     int          `json:"story_id"`
	ParentId    *int         `json:"parent_id,omitempty"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Image       *string      `json:"image,omitempty"`
	Children    []*StoryTree `json:"children,omitempty"`
}

type LikeResponse struct {
	Message string `json:"message"`
}

func GetAllStories() []Stories {
	db := openConnection()

	result, err := db.Query("SELECT * FROM stories")
	if err != nil {
		log.Fatal(err)
	}

	var storyRows []Stories
	for result.Next() {
		var r Stories
		if err := result.Scan(&r.Id, &r.Title, &r.Description, &r.Image, &r.Likes, &r.Views); err != nil {
			log.Fatal(err)
		}

		storyRows = append(storyRows, r)
	}

	return storyRows
}

func AddLikeToStory(id int) LikeResponse {
	db := openConnection()

	_, err := db.Query("UPDATE stories SET likes = likes + 1 WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	return LikeResponse{
		Message: "Thanks for like",
	}
}

func GetStoryTreeByStoryId(id int) *StoryTree {
	db := openConnection()

	_, err := db.Query("UPDATE stories SET views = views + 1 WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	result, err := db.Query("SELECT * FROM stories_data WHERE story_id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	var StoryTreeRows []StoryTree
	for result.Next() {
		var r StoryTree
		if err := result.Scan(&r.Id, &r.StoryId, &r.ParentId, &r.Title, &r.Description, &r.Image); err != nil {
			log.Fatal(err)
		}

		StoryTreeRows = append(StoryTreeRows, r)
	}

	var resultTree *StoryTree
	temp := make(map[int]*StoryTree)
	for _, item := range StoryTreeRows {
		story := &StoryTree{
			Id:          item.Id,
			StoryId:     item.StoryId,
			ParentId:    item.ParentId,
			Title:       item.Title,
			Description: item.Description,
			Image:       item.Image,
			Children:    []*StoryTree{},
		}

		temp[item.Id] = story

		if item.ParentId == nil {
			resultTree = story
			continue
		}

		temp[*item.ParentId].addChild(story)
	}

	return resultTree
}

func (c *StoryTree) addChild(child *StoryTree) {
	c.Children = append(c.Children, child)
}

func openConnection() *sql.DB {
	conString := config.GetConnectionString()
	db, err := sql.Open("postgres", conString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return db
}
