package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Hello": "World",
		})

		// ctx.String(http.StatusOK, "Value:%v", "\nHello World")
	})

	r.POST("/post", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Method": "POST",
		})
	})

	r.PUT("/put", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Method": "PUT",
		})
	})

	r.DELETE("/delete", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Method": "DELETE",
		})
	})

	r.Run()
}
