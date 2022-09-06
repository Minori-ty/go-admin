package services

import (
	"main/models"
)

func GetArticle(id string) models.Article {
	article := models.Article{
		Id:      1,
		Content: "内容",
	}
	return article
}
