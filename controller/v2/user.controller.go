package v2

import (
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(UserService service.UserService) UserController {
	return UserController{
		UserService: UserService,
	}
}

func (u *UserController) GetUser(c *gin.Context) {
	username := c.Query("name")
	user := u.UserService.GetUser(username)
	c.JSON(http.StatusOK, user)
}

func (u *UserController) AddUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "添加成功",
	})
}

func (u *UserController) RegisterRoutes(r *gin.RouterGroup) {
	router := r.Group("/")
	router.GET("/get", u.GetUser)
	router.POST("/add", u.AddUser)
}
