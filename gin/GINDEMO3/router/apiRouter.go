package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	apiRouter := r.Group("/api")
	{
		apiRouter.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "api interface")
		})

		apiRouter.GET("/userlist", func(c *gin.Context) {
			c.String(http.StatusOK, "user list")
		})

		apiRouter.GET("/plist", func(c *gin.Context) {
			c.String(http.StatusOK, "plist")
		})
	}
}
