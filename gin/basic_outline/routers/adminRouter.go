package routers

import (
	"gin_tutorial/basic_outline/controller"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.Engine) {

	adminController := controller.NewAdminController()

	adminRouter := r.Group("/admin")
	{
		adminRouter.GET("/", adminController.AController)
	}
}
