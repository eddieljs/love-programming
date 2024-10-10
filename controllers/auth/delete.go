package auth

import (
	"net/http"
	"untitled/models"
	"untitled/tools"

	"github.com/gin-gonic/gin"
)

type DeleteCon struct{}

func (DeleteCon) DeleteVideo(ctx *gin.Context) {
	videoInfo := struct {
		Id int `json:"id" form:"id"`
	}{}
	if err := ctx.ShouldBind(&videoInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	video := models.Video{}
	models.DB.Where("id = ?", videoInfo.Id).Delete(&video)
	tools.Success(ctx, gin.H{}, "删除成功")
}

func (DeleteCon) DeleteResource(ctx *gin.Context) {
	resourceInfo := struct {
		Id int `json:"id" form:"id"`
	}{}
	if err := ctx.ShouldBind(&resourceInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	resource := models.Resource{}
	models.DB.Where("id = ?", resourceInfo.Id).Delete(&resource)
	tools.Success(ctx, gin.H{}, "删除成功")
}

func (DeleteCon) DeleteChoice(ctx *gin.Context) {
	choiceInfo := struct {
		Id int `json:"id" form:"id"`
	}{}
	if err := ctx.ShouldBind(&choiceInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	choice := models.Choice{}
	models.DB.Where("id = ?", choiceInfo.Id).Delete(&choice)
	tools.Success(ctx, gin.H{}, "删除成功")
}

func (DeleteCon) DeleteCode(ctx *gin.Context) {
	codeInfo := struct {
		Id int `json:"id" form:"id"`
	}{}
	if err := ctx.ShouldBind(&codeInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	code := models.Code{}
	models.DB.Where("id = ?", codeInfo.Id).Delete(&code)
	tools.Success(ctx, gin.H{}, "删除成功")
}

func (DeleteCon) DeletePoint(ctx *gin.Context) {
	pointInfo := struct {
		Id int `json:"id" form:"id"`
	}{}
	if err := ctx.ShouldBind(&pointInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	point := models.ExercisePoint{}
	models.DB.Where("id = ?", pointInfo.Id).Delete(&point)
	models.DB.Where("point_id = ?", pointInfo.Id).Delete(&models.Choice{})
	models.DB.Where("point_id = ?", pointInfo.Id).Delete(&models.Code{})
	tools.Success(ctx, gin.H{}, "删除成功")
}
