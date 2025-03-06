package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/leebrouse/gorm/controllers"
)

func StudentRouterInit(r *gin.Engine) {
	studentController := controllers.NewStudentControllers()

	studentRouter := r.Group("/student")
	{
		studentRouter.GET("/index", studentController.Index)
	}
}
