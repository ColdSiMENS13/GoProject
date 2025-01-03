package StoriesService

import (
	StoriesRepository "GoApp/Service/Stories/Repository"
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
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Image       *string      `json:"image,omitempty"`
	Children    []*StoryTree `json:"children,omitempty"`
}

type LikeResponse struct {
	Message string `json:"message"`
}

func GetAllStories() []Stories {
	result := StoriesRepository.GetAllStories()

	var stories []Stories
	for _, item := range result {
		stories = append(stories, Stories{
			Id:          item.Id,
			Title:       item.Title,
			Description: item.Description,
			Image:       item.Image,
			Likes:       item.Likes,
			Views:       item.Views,
		})
	}

	return stories
}

func AddLikeToStory(id int) LikeResponse {
	message := StoriesRepository.AddLikeToStory(id)

	return LikeResponse{
		Message: message,
	}
}

func GetStoryTreeByStoryId(id int) *StoryTree {
	result := StoriesRepository.GetStoryTreeByStoryId(id)

	var resultTree *StoryTree
	temp := make(map[int]*StoryTree)
	for _, item := range result {
		story := &StoryTree{
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
