package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "首页",
		})
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
