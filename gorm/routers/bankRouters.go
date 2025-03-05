package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/leebrouse/gorm/controllers"
)

func BankRouterInit(r *gin.Engine) {
	// create nav interface
	bankControllers := controllers.NewBankControllers()
	//router group
	navRouter := r.Group("/bank")
	{
		navRouter.GET("/index", bankControllers.Increase)
	}
}
