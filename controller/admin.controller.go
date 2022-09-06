package controller

import (
	"main/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArticle(c *gin.Context) {
	var id = "1"
	article := services.GetArticle(id)
	c.JSON(http.StatusOK, article)
}

func RegisterArticleRoutes(r *gin.RouterGroup) {
	userroute := r.Group("/")
	userroute.GET("/article", GetArticle)
}
