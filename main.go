package main

import (
	"main/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controller.RegisterUserRoutes(&r.RouterGroup)
	controller.RegisterArticleRoutes(&r.RouterGroup)
	r.Run()
}
