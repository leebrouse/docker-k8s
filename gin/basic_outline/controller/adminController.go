package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// interface
type IAdminController interface {
	AController(c *gin.Context)
}

// struct
type AdminController struct{}

// function

func NewAdminController() IAdminController {
	return &AdminController{}
}

func (admin *AdminController) AController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "This is admin ",
	})
}
