package learn

import (
	"github.com/gin-gonic/gin"
	"untitled/tools"
)

type LearnController struct{}

func (con LearnController) Index(ctx *gin.Context) {
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"title": "标题",
	// 	"tags":  "标签",
	// 	"url":   "链接",
	// })
	var title = []string{"教学视频", "学习资源", "选择题", "讨论区", "代码补全题"}
	tools.Success(ctx, gin.H{
		"title": title,
	}, "")
}
