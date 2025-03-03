package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// interface
type IDefaultController interface {
	Default(c *gin.Context)
	News(c *gin.Context)
	Shop(c *gin.Context)
	Delete(c *gin.Context)
}

// struct
type DefaultController struct{}

// function

func NewDefaultController() IDefaultController {
	return &DefaultController{}
}

func (admin *DefaultController) Default(c *gin.Context) {

	//set cookie
	c.SetCookie("username", "leebrouse", 3600, "/", "localhost", false, false)
	//add cookie hobby

	c.JSON(http.StatusOK, gin.H{
		"message":  "This is default ",
		"function": "Set cookie",
	})
}

func (admin *DefaultController) News(c *gin.Context) {

	username, err := c.Cookie("username")
	if err != nil {
		c.String(http.StatusBadRequest, "Don`t have the cookie key")
		return
	}

	//get cookie hobby

	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"message":  "This is default ",
		"function": "Get cookie value",
	})
}

func (admin *DefaultController) Shop(c *gin.Context) {
	username, err := c.Cookie("username")
	if err != nil {
		c.String(http.StatusBadRequest, "Don`t have the cookie key")
		return
	}

	//get cookie hobby

	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"message":  "This is default ",
		"function": "Get cookie value",
	})
}

func (admin *DefaultController) Delete(c *gin.Context) {

	c.SetCookie("username", "leebrouse", -1, "/", "localhost", false, false)

	c.JSON(http.StatusOK, gin.H{
		"message":  "This is default ",
		"function": "Delete cookie",
		"success":  true,
	})
}
