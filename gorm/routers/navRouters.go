package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/leebrouse/gorm/controllers"
)

func NavRouterInit(r *gin.Engine) {
	// create nav interface
	navController := controllers.NewNavControllers()
	//router group
	navRouter := r.Group("/nav")
	{
		navRouter.GET("/index", navController.Index)
		navRouter.GET("/add", navController.Add)
		navRouter.GET("/update", navController.Update)
		navRouter.GET("/delete", navController.Delete)
	}
}
