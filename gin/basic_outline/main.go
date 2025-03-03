package main

import (
	"gin_tutorial/basic_outline/routers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	//adminRouter
	routers.AdminRouterInit(r)
	//defaultRouter
	routers.DefaultRouterInit(r)

	return r
}

func main() {
	r := NewRouter()

	r.Run()
}
