package routers

import (
	"gin_tutorial/GINDEMO5/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.Engine) {

	adminController := admin.NewAdminController()

	adminRouter := r.Group("/admin")
	{
		adminRouter.GET("/add", adminController.Add)
		adminRouter.POST("/doupload", adminController.DoUpload)

		adminRouter.GET("/edit", adminController.Edit)
		adminRouter.POST("/editUpload", adminController.EditUpload)
	}
}
