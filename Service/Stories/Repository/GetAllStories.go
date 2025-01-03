package StoriesRepository

import (
	config "GoApp/databaseConf/database"
	"log"
)

type Stories struct {
	Id          int
	Title       string
	Description string
	Image       *string
	Likes       int
	Views       int
}

func GetAllStories() []Stories {
	db := config.OpenConnection()

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
	result.Close()

	return storyRows
}
