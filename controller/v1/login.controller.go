package v1

import (
	"fmt"
	"main/model"
	"main/service"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginService service.LoginService
}

func NewLoginController(LoginService service.LoginService) LoginController {
	return LoginController{}
}

func (u LoginController) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	loginform := &model.LoginForm{
		Username: username,
		Password: password,
	}
	fmt.Printf("%#v", loginform)
	c.JSON(200, u.LoginService.Login(loginform))
}

func (u *LoginController) RegisterRoutes(r *gin.RouterGroup) {
	userroute := r.Group("/")
	userroute.POST("/login", u.Login)
}
