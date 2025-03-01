package main

import (
	"gin_tutorial/GINDEMO3/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//default
	router.DefaultRouter(r)
	//api
	router.ApiRouter(r)
	//admin
	router.AdminRouter(r)
	r.Run()
}
