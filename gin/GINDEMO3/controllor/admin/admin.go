package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type UserController interface {
// 	Userindex(ctx *gin.Context)
// 	UserAdd(ctx *gin.Context)
// }

type UserControllerImpl struct {
	BaseControllor //inheritance
}

func (c UserControllerImpl) Userindex(ctx *gin.Context) {
	ctx.String(http.StatusOK, "用户列表")
}

func (c UserControllerImpl) UserAdd(ctx *gin.Context) {
	ctx.String(http.StatusOK, "添加用户")

}
