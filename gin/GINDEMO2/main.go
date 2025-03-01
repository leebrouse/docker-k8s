package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// struct members should use uppercase
type Articals struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./template/**/*")

	//front
	r.GET("/", func(c *gin.Context) {
		username := c.Query("username")
		age := c.Query("age")
		page := c.DefaultQuery("page", "1")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})

	// GET 请求，渲染用户注册页面
	r.GET("/user", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/user.html", gin.H{})
	})

	// POST 请求，处理用户注册数据并返回JSON
	r.POST("/doAdduser", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	r.POST("/doAdduser2", func(ctx *gin.Context) {
		user := &UserInfo{}
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, user)
	})

	r.GET("/getuser", func(ctx *gin.Context) {
		user := &UserInfo{}

		fmt.Printf("user value%#v", user)
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, user)
		}

	})

	r.GET("/news", func(c *gin.Context) {
		news := &Articals{
			Title:   "title",
			Content: "content",
		}

		c.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "Newspaper Title",
			"news":  news,
		})
	})

	//back
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台首页",
		})
	})

	r.GET("/admin/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/news.html", gin.H{
			"message": "后台新闻",
		})
	})
	r.Run()
}
