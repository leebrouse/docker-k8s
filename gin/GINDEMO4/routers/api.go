package routers

import (
	controllers "gin_tutorial/GINDEMO4/Controllers"
	"gin_tutorial/GINDEMO4/middleware"

	"github.com/gin-gonic/gin"
)

func ApiRouterInit(r *gin.Engine) {
	apiController := controllers.NewApiController()

	//global middleware
	//  r.Use(middleware.InitMiddlewareOne, middleware.InitMiddlewareTwo)

	// apiRouter := r.Group("/", middleware.InitMiddlewareOne, middleware.InitMiddlewareTwo)
	apiRouter := r.Group("/")
	{
		apiRouter.GET("/", apiController.Default)
		apiRouter.GET("/news", apiController.GetNews)

		//add SharedDataMiddle
		apiRouter.GET("/user", middleware.SharedDataMiddle, apiController.GetUserName)
	}
}

//Predict the output of this code snippet:

//Middleware 1 executed before handler
//Middleware 2 executed before handler
//default interface
//Middleware 2 executed after handler
//Middleware 1 executed after handler
