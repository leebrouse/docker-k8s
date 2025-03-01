package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddlewareOne(ctx *gin.Context) {
	// start := time.Now()
	fmt.Println("Middleware 1 executed before handler")
	// fmt.Println("Current time:", start)
	ctx.Next() // 继续执行后续的 handler

	fmt.Println("Middleware 1 executed after handler")
	// fmt.Println("Elapsed time:", time.Since(start))
}

func InitMiddlewareTwo(ctx *gin.Context) {
	fmt.Println("Middleware 2 executed before handler")

	ctx.Next()

	fmt.Println("Middleware 2 executed after handler")
}

func SharedDataMiddle(ctx *gin.Context) {
	fmt.Println(time.Now())
	fmt.Println(ctx.Request.URL)
	ctx.Set("username", "leebrouse")

	cCp := ctx.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done! in path", cCp.Request.URL.Path)
	}()
}
