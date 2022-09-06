package services

import (
	"main/models"
)

func GetUser(name string) models.User {
	user := models.User{
		Id:       1,
		Username: "123",
		Age:      27,
	}
	return user
}
