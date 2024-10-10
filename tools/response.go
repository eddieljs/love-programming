package tools

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, message string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  message,
	})
}

func Success(ctx *gin.Context, data gin.H, message string) {
	Response(ctx, http.StatusOK, 200, data, message)
}
func Fail(ctx *gin.Context, data gin.H, message string) {
	Response(ctx, http.StatusOK, 400, data, message)
}
