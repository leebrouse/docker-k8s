package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRouter(r *gin.Engine) {
	defaultRouter := r.Group("/")
	{
		defaultRouter.GET("/", )

		defaultRouter.GET("/news", func(c *gin.Context) {
			c.String(http.StatusOK, "新闻")
		})
	}
}
