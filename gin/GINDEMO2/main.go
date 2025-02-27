package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// struct members should use uppercase
type Articals struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./template/*")

	r.GET("/ping", func(c *gin.Context) {

		a := &Articals{
			Title:       "title",
			Description: "description",
			Name:        "name",
		}

		c.JSON(http.StatusOK, a)
	})

	// http://localhost:8080/ping2?callback=hello
	// hello({"title":"title","description":"description","name":"name"});
	r.GET("/ping2", func(c *gin.Context) {

		a := &Articals{
			Title:       "title",
			Description: "description",
			Name:        "name",
		}

		c.JSONP(http.StatusOK, a)
	})

	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"success": true,
			"message": "Hello, world!",
		})
	})

	r.GET("/html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "new.html", gin.H{
			"title": "Main website",
		})
	})

	r.GET("/string", func(c *gin.Context) {
		name := c.Query("name")

		c.String(http.StatusOK, "Hello %s", name)
	})

	r.Run()
}
