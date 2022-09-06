package server

import (
	"context"
	v1 "main/controller/v1"
	v2 "main/controller/v2"
	"main/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func registerV1Routes(r *gin.Engine, collection *mongo.Collection, ctx context.Context) {
	route := r.Group("/v1")

	LoginService := service.NewLoginService(collection, ctx)
	LoginController := v1.NewLoginController(LoginService)
	LoginController.RegisterRoutes(route)
}

func registerV2Routes(r *gin.Engine, collection *mongo.Collection, ctx context.Context) {
	route := r.Group("/v2")

	UserService := service.NewUserService(collection, ctx)
	UserController := v2.NewUserController(UserService)
	UserController.RegisterRoutes(route)

	AdminService := service.NewAdminService()
	AdminController := v2.NewAdminController(AdminService)
	AdminController.RegisterRoutes(route)
}

func RegisterRoutes(r *gin.Engine, collection *mongo.Collection, ctx context.Context) {
	registerV1Routes(r, collection, ctx)

	registerV2Routes(r, collection, ctx)
}
