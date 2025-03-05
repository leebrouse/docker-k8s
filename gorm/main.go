package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leebrouse/gorm/routers"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//add user routers
	routers.UserRoutersInit(r)
	//add nav routers
	routers.NavRouterInit(r)
	//bank routers
	routers.BankRouterInit(r)

	return r
}
func main() {
	r := NewRouter()

	r.Run()
}
