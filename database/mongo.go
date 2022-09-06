package database

import (
	"context"
	"fmt"
	"log"
	"main/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx context.Context
)

func NewMongoDB() {
	ctx = context.TODO()
	MongoConnection := options.Client().SetAuth(options.Credential{
		Username: "root",
		Password: "root",
	}).ApplyURI("mongodb://" + config.Mongo)

	MongoClient, err := mongo.Connect(ctx, MongoConnection)

	if err != nil {
		log.Fatal(err)
	}
	err = MongoClient.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("连接上mongoDB")
	UserCollection := MongoClient.Database("userdb").Collection("users")
	fmt.Println(UserCollection)
}
