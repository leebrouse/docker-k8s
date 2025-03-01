package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 1. 定义接口，解耦控制器
type IApiController interface {
	Default(c *gin.Context)
	GetNews(c *gin.Context)
	GetUserName(c *gin.Context)
}

// 2. 具体实现 ApiController
type ApiController struct{}

// 3. 工厂方法，返回接口类型
func NewApiController() IApiController {
	return &ApiController{}
}

// 4. 实现 IApiController 接口
func (api *ApiController) Default(c *gin.Context) {
	c.String(http.StatusOK, "default interface")
}

func (api *ApiController) GetNews(c *gin.Context) {
	c.String(http.StatusOK, "get news interface")
}

func (api *ApiController) GetUserName(c *gin.Context) {
	username, _ := c.Get("username")

	v, _ := username.(string)
	c.String(http.StatusOK, "Hello username:"+v) //类型断言
}
