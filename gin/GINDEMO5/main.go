package main

import (
	"gin_tutorial/GINDEMO5/routers"

	"github.com/gin-gonic/gin"
)

// factory pattern
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/**/*")

	//You can add routers,middlerwares,handlers and so on here

	routers.DefaultRouterInit(r) //default router
	routers.AdminRouterInit(r)   //admin router

	return r
}

func main() {
	r := NewRouter()
	
	r.Run()
}
