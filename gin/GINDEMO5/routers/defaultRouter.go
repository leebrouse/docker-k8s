package routers

import (
	"gin_tutorial/GINDEMO5/controllers/itying"

	"github.com/gin-gonic/gin"
)

func DefaultRouterInit(r *gin.Engine) {

	defaultController := itying.NewDefaultController()

	r.GET("/", defaultController.DefaultTime)
}
