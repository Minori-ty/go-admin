package controller

import (
	"main/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var username = "saff"
	user := services.GetUser(username)
	c.JSON(http.StatusOK, user)
}

func AddUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "添加成功",
	})
}

func RegisterUserRoutes(r *gin.RouterGroup) {
	router := r.Group("/")
	router.GET("/get", GetUser)
	router.POST("/add", AddUser)
}
