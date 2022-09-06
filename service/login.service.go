package service

import (
	"context"
	"fmt"
	"main/model"
)

type LoginService struct {
	ctx context.Context
}

func NewLoginService() LoginService {
	return LoginService{}
}

func (u LoginService) Login(loginform *model.LoginForm) model.Login {

	fmt.Println(loginform)
	login := model.Login{
		Code:     "200",
		Message:  "登录成功",
		Username: "守夜人",
		Token:    "bearToken abcdefg",
	}
	return login
}
