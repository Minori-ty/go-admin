package main

import (
	"context"
	"main/database"
	"main/middleware"
	"main/server"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection  *mongo.Collection
	ctx         context.Context
	r           *gin.Engine
	MongoClient *mongo.Client
)

func init() {
	ctx = context.TODO()
	MongoClient = database.NewMongoDB()
	r = gin.Default()
	r.Use(middleware.Cors())
	collection = MongoClient.Database("web").Collection("users")

	server.RegisterRoutes(r, collection, ctx)
}

func main() {
	defer MongoClient.Disconnect(ctx)
	r.Run()
}
