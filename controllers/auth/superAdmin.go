package auth

import (
	"net/http"
	"time"
	"untitled/models"
	"untitled/tools"

	"github.com/gin-gonic/gin"
)

type SuperAdminCon struct{}

func (SuperAdminCon) Exam(ctx *gin.Context) {
	// 接收id,f。
	ceInfo := struct {
		Id int    `json:"id" form:"id"`
		F  string `json:"f" form:"f"`
	}{}
	if err := ctx.ShouldBind(&ceInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	ce := models.CourseExam{}
	models.DB.Where("id = ?", ceInfo.Id).First(&ce)
	cu := models.CourseUser{
		UserId: ce.UserId,
	}
	course := models.Course{
		Title:     ce.CourseTitle,
		Key:       ce.CourseKey,
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		TeacherId: ce.UserId,
	}
	if ceInfo.F == "y" {
		tx := models.DB.Begin()
		tx.Where("id = ?", ce.Id).Delete(&ce)
		tx.Create(&course)
		cu.CourseId = course.Id
		tx.Create(&cu)
		tx.Commit()
		tools.Success(ctx, gin.H{
			"course": course,
		}, "审核通过")
	} else if ceInfo.F == "n" {
		models.DB.Where("id = ?", ce.Id).Delete(&ce)
		tools.Success(ctx, gin.H{
			"course": course,
		}, "审核不通过")
	}
}

func (SuperAdminCon) CourseList(ctx *gin.Context) {
	ceList := []models.CourseExam{}
	models.DB.Preload("User").Find(&ceList)
	tools.Success(ctx, gin.H{
		"list": ceList,
	}, "申请列表")
}
