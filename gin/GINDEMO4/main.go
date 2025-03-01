package main

import (
	"gin_tutorial/GINDEMO4/routers"

	"github.com/gin-gonic/gin"
)

// Factory Pattern
func NewRouter() *gin.Engine {
	r := gin.Default()

	//You can add routers,middlerwares,handlers and so on here
	// TODO:
	//routers
	routers.ApiRouterInit(r)

	return r
}

func main() {
	r := NewRouter()

	r.Run()
}
