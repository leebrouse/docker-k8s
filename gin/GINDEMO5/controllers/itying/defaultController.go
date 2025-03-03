package itying

import (
	"gin_tutorial/GINDEMO5/Utiles"

	"github.com/gin-gonic/gin"
)

type IDefaultController interface {
	DefaultTime(c *gin.Context)
}

type DefaultController struct{}

func (dc *DefaultController) DefaultTime(c *gin.Context) {

	time := Utiles.UnixToTime(1600000000)

	c.JSON(200, gin.H{
		"message": "Hello, World!",
		"time":    time,
	})
}

func NewDefaultController() IDefaultController {
	return &DefaultController{}
}
