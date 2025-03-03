package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// interface
type IDefaultController interface {
	DController(c *gin.Context)
}

// struct
type DefaultController struct{}

// function

func NewDefaultController() IDefaultController {
	return &DefaultController{}
}

func (admin *DefaultController) DController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "This is default ",
	})
}
