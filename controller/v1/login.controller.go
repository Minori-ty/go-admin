package v1

import (
	"main/model"
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginService service.LoginServiceInt
}

func NewLoginController(LoginService service.LoginServiceInt) LoginController {
	return LoginController{
		LoginService: LoginService,
	}
}

func (u *LoginController) Login(ctx *gin.Context) {
	var form model.LoginForm
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := u.LoginService.Login(&form)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "添加失败",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":    2000,
		"message": "添加成功",
	})
}

func (u *LoginController) RegisterRoutes(r *gin.RouterGroup) {
	userroute := r.Group("/")
	userroute.POST("/login", u.Login)
}
