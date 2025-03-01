package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseControllor struct{}

func (bc BaseControllor) Success(ctx *gin.Context) {
	ctx.String(http.StatusOK, "success")
}

func (bc BaseControllor) Error(ctx *gin.Context) {
	ctx.String(http.StatusOK, "error")
}
