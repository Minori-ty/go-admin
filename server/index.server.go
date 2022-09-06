package server

import (
	v1 "main/controller/v1"
	v2 "main/controller/v2"
	"main/service"

	"github.com/gin-gonic/gin"
)

func registerV1Routes(r *gin.Engine) {
	route := r.Group("/v1")

	LoginService := service.NewLoginService()
	LoginController := v1.NewLoginController(LoginService)
	LoginController.RegisterRoutes(route)
}

func registerV2Routes(r *gin.Engine) {
	route := r.Group("/v2")

	UserService := service.NewUserService()
	UserController := v2.NewUserController(UserService)
	UserController.RegisterRoutes(route)

	AdminService := service.NewAdminService()
	AdminController := v2.NewAdminController(AdminService)
	AdminController.RegisterRoutes(route)
}

func RegisterRoutes(r *gin.Engine) {
	registerV1Routes(r)
	registerV2Routes(r)
}
