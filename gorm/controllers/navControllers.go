package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leebrouse/gorm/models"
)

type INavControllers interface {
	IBaseControllers
	Index(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type NavControllers struct {
	BaseControllers
}

// function
func NewNavControllers() INavControllers {
	return &NavControllers{}
}

func (nav *NavControllers) Index(c *gin.Context) {
	navs := []models.Navs{}
	models.DB.Find(&navs)

	c.JSON(http.StatusOK, gin.H{
		"Navs":    navs,
		"success": true,
	})
}

func (nav *NavControllers) Add(c *gin.Context) {

}

func (nav *NavControllers) Update(c *gin.Context) {

}

func (nav *NavControllers) Delete(c *gin.Context) {

}
