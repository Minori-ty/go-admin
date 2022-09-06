package v2

import (
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	AdminService service.AdminService
}

func NewAdminController(AdminService service.AdminService) AdminController {
	return AdminController{
		AdminService: AdminService,
	}
}

func (u *AdminController) GetArticle(c *gin.Context) {
	var id = "1"
	article := u.AdminService.GetArticle(id)
	c.JSON(http.StatusOK, article)
}

func (u *AdminController) RegisterRoutes(r *gin.RouterGroup) {
	userroute := r.Group("/")
	userroute.GET("/article", u.GetArticle)
}
