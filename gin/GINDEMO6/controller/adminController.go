package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// interface
type IAdminController interface {
	Default(c *gin.Context)
	News(c *gin.Context)
}

// struct
type AdminController struct{}

// function

func NewAdminController() IAdminController {
	return &AdminController{}
}

func (admin *AdminController) Default(c *gin.Context) {

	//set session
	session := sessions.Default(c)
	session.Set("username", "leebrouse")
	session.Save() //It is need to save after set session key and value

	c.JSON(http.StatusOK, gin.H{
		"message": "This is admin ",
	})
}

func (admin *AdminController) News(c *gin.Context) {

	session := sessions.Default(c)
	username := session.Get("username").(string)

	c.JSON(http.StatusOK, gin.H{
		"message":  "This is admin ",
		"username": username,
	})
}
