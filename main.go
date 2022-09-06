package main

import (
	"main/database"
	"main/middleware"
	"main/server"

	"github.com/gin-gonic/gin"
)

func init() {
	database.NewMongoDB()
}
func main() {
	r := gin.Default()
	r.Use(middleware.Cors())
	server.RegisterRoutes(r)

	r.Run()
}
