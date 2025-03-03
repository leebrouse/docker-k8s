package routers

import (
	"gin_tutorial/GINDEMO6/controller"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.Engine) {

	adminController := controller.NewAdminController()

	adminRouter := r.Group("/admin")
	{
		adminRouter.GET("/", adminController.Default)
		adminRouter.GET("/news", adminController.News)
	}
}
