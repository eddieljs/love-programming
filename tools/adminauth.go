package tools

import (
	"net/http"
	"wcapp/models"

	"github.com/gin-gonic/gin"
)

func AdminAuth(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	if user.(models.User).Auth < 7 {
		Response(ctx, http.StatusOK, 401, gin.H{}, "需要管理员权限")
		ctx.Abort()
		return
	}
}

func SuperAdminAuth(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	if user.(models.User).Auth < 9 {
		Response(ctx, http.StatusOK, 401, gin.H{}, "需要超管权限")
		ctx.Abort()
		return
	}
}
