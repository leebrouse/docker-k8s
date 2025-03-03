package routers

import (
	"gin_tutorial/GINDEMO5/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.Engine) {

	adminController := admin.NewAdminController()

	adminRouter := r.Group("/admin")
	{
		// 单个文件上传
		adminRouter.GET("/add", adminController.Add)
		adminRouter.POST("/doupload", adminController.DoUpload)

		//多个不同名字的文件上传
		adminRouter.GET("/edit", adminController.Edit)
		adminRouter.POST("/editUpload", adminController.EditUpload)

		//多个相同名字的文件上传
		adminRouter.GET("/multipart", adminController.Multipart)
		adminRouter.POST("/multipartUpload", adminController.MultipartUpload)
	}
}
