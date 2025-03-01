package router

import (
	"gin_tutorial/GINDEMO3/controllor/admin"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.Engine) {

	// userController := &admin.UserControllerImpl{}

	adminRouter := r.Group("/admin")
	{
		adminRouter.GET("/", admin.UserControllerImpl{}.Userindex)

		adminRouter.GET("/success", admin.UserControllerImpl{}.Success)
		adminRouter.GET("/error", admin.UserControllerImpl{}.Error)

		adminRouter.GET("/userlist", admin.UserControllerImpl{}.UserAdd)

		adminRouter.GET("/plist", func(c *gin.Context) {
			c.String(http.StatusOK, "backend plist")
		})
	}
}
