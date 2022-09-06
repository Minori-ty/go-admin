package service

import (
	"context"
	"fmt"
	"main/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type LoginServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewLoginService(collection *mongo.Collection, ctx context.Context) LoginServiceInt {
	return &LoginServiceImpl{
		ctx:        ctx,
		collection: collection,
	}
}

func (u *LoginServiceImpl) Login(loginform *model.LoginForm) error {
	// var user *model.User
	fmt.Println(loginform)
	_, err := u.collection.InsertOne(u.ctx, loginform)
	// query := bson.D{bson.E{Key: "user_name", Value: "user"}}
	// err := u.collection.FindOne(u.ctx, query).Decode(&user)
	// fmt.Println(u.collection)
	// _, err := u.collection.InsertOne(u.ctx, model.User{
	// 	Id:       1,
	// 	Username: "v2",
	// 	Age:      24,
	// })

	return err
}
