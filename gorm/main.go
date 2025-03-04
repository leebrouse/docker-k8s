package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leebrouse/gorm/routers"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//add routers
	routers.UserRoutersInit(r)

	return r
}
func main() {
	r := NewRouter()

	r.Run()
}
