package service

import "main/model"

type LoginServiceInt interface {
	Login(loginform *model.LoginForm) error
}
