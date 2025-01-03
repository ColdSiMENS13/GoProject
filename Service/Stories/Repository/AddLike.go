package StoriesRepository

import (
	config "GoApp/databaseConf/database"
	"log"
)

func AddLikeToStory(id int) string {
	db := config.OpenConnection()

	_, err := db.Query("UPDATE stories SET likes = likes + 1 WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	return "Successfully added like!"
}
