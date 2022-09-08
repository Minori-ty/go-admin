package service

import (
	"context"
	"fmt"
	"main/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewUserService(collection *mongo.Collection, ctx context.Context) UserService {
	return UserService{
		ctx:        ctx,
		collection: collection,
	}
}

func (u UserService) GetUser(name string) model.User {
	var user *model.User
	query := bson.D{bson.E{Key: "username", Value: name}}
	err := u.collection.FindOne(u.ctx, query).Decode(&user)
	if err != nil {
		fmt.Println(name)
		fmt.Printf("%#v", user)
		fmt.Println("查询失败")
	}
	return *user
}
