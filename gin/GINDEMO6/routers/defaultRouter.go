package routers

import (
	"gin_tutorial/GINDEMO6/controller"

	"github.com/gin-gonic/gin"
)

func DefaultRouterInit(r *gin.Engine) {

	defaultController := controller.NewDefaultController()

	defaultRouter := r.Group("/")
	{
		defaultRouter.GET("/", defaultController.Default)
		defaultRouter.GET("/news", defaultController.News)
		defaultRouter.GET("/shop", defaultController.Shop)
		defaultRouter.GET("/delete", defaultController.Delete)
	}
}
