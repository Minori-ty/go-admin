package service

import (
	"context"
	"fmt"
	"main/model"
)

type AdminService struct {
	ctx context.Context
}

func NewAdminService() AdminService {
	return AdminService{}
}

func (u *AdminService) GetArticle(id string) model.Article {
	fmt.Println(u.ctx)
	article := model.Article{
		Id:      1,
		Content: "内容",
	}
	return article
}
