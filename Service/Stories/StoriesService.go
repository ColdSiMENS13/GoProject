package StoriesService

import (
	config "GoApp/databaseConf/database"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Stories struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       *string `json:"image,omitempty"`
}

type StoryTree struct {
	Id          int         `json:"id"`
	StoryId     int         `json:"story_id"`
	ParentId    *int        `json:"parent_id,omitempty"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Image       *string     `json:"image,omitempty"`
	Children    []StoryTree `json:"children,omitempty"`
}

func GetAllStories() []Stories {
	conString := config.GetConnectionString()

	db, err := sql.Open("postgres", conString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Query("SELECT * FROM stories")
	if err != nil {
		log.Fatal(err)
	}

	var storyRows []Stories
	for result.Next() {
		var r Stories
		if err := result.Scan(&r.Id, &r.Title, &r.Description, &r.Image); err != nil {
			log.Fatal(err)
		}

		storyRows = append(storyRows, r)
	}

	return storyRows
}

func GetStoryTreeByStoryId(id int) []StoryTree {
	conString := config.GetConnectionString()
	db, err := sql.Open("postgres", conString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

	var resultTree []StoryTree
	temp := make(map[int]*StoryTree)
	for _, item := range StoryTreeRows {
		story := &StoryTree{
			Id:          item.Id,
			StoryId:     item.StoryId,
			ParentId:    item.ParentId,
			Title:       item.Title,
			Description: item.Description,
			Image:       item.Image,
			Children:    []StoryTree{},
		}

		temp[item.Id] = story

		if item.ParentId == nil {
			resultTree = append(resultTree, *story)
			continue
		}

		temp[*item.ParentId].addChild(*story)
	}

	return resultTree
}

func (c *StoryTree) addChild(child StoryTree) {
	c.Children = append(c.Children, child)
}
