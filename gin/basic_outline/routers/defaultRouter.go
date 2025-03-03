package routers

import (
	"gin_tutorial/basic_outline/controller"

	"github.com/gin-gonic/gin"
)

func DefaultRouterInit(r *gin.Engine) {

	defaultController := controller.NewDefaultController()

	defaultRouter := r.Group("/")
	{
		defaultRouter.GET("/", defaultController.DController)
	}
}
