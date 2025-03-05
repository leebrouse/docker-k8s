package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// interface
type IBaseControllers interface{
	Success(c *gin.Context)
	Fail(c *gin.Context)
}

// struct
type BaseControllers struct{}

// function
func (base *BaseControllers) Success(c *gin.Context) {
	c.String(http.StatusOK, "Operate Success")
}

func (base *BaseControllers) Fail(c *gin.Context) {
	c.String(http.StatusBadRequest, "Operate Failed")
}
