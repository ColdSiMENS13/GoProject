package StoriesRepository

import (
	config "GoApp/databaseConf/database"
	"log"
)

type StoryTree struct {
	Id          int
	StoryId     int
	ParentId    *int
	Title       string
	Description string
	Image       *string
}

func GetStoryTreeByStoryId(id int) []StoryTree {
	db := config.OpenConnection()

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
	result.Close()

	return StoryTreeRows
}
