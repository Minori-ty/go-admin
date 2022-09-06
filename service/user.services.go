package service

import (
	"context"
	"fmt"
	"main/model"
)

type UserService struct {
	ctx context.Context
}

func NewUserService() UserService {
	return UserService{}
}

func (u UserService) GetUser(name string) model.User {
	username := u.ctx
	fmt.Println(username)
	user := model.User{
		Id:       1,
		Username: "123",
		Age:      27,
	}
	return user
}
