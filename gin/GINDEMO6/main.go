package main

import (
	"gin_tutorial/GINDEMO6/routers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte("serect"))
	r.Use(sessions.Sessions("mysession", store))
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
