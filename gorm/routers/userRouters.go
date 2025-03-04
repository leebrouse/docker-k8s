package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/leebrouse/gorm/controllers"
)

func UserRoutersInit(r *gin.Engine) {

	//new UserController
	userController := controllers.NewUserController()

	userRouter := r.Group("/user")
	{
		userRouter.GET("/index", userController.Index)
		userRouter.GET("/add", userController.Add)
		userRouter.GET("/update", userController.Update)
		userRouter.GET("/delete", userController.Delete)
	}

}
