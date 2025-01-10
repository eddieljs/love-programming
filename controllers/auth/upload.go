package auth

import (
	"log"
	"net/http"
	"time"
	"untitled/models"
	"untitled/tools"

	"github.com/gin-gonic/gin"
)

type UploadCon struct{}

func (UploadCon) UploadResource(ctx *gin.Context) {
	// 接收title,url,courseId
	resource := models.Resource{}
	if err := ctx.ShouldBind(&resource); err != nil {
		tools.Fail(ctx, gin.H{
			"ERROR": err.Error(),
		}, "获取信息失败")
		return
	}
	// 添加资源
	models.DB.Create(&resource)
	tools.Success(ctx, gin.H{
		"file": resource,
	}, "上传成功")
	// if fileInfo.Type == "p" {
	// 	file = file.(models.Resource)
	// } else if fileInfo.Type == "v" {
	// 	file = file.(models.Video)
	// }
}

func (UploadCon) UploadVideo(ctx *gin.Context) {
	// 接收title,url,courseId
	video := models.Video{}
	if err := ctx.ShouldBind(&video); err != nil {
		tools.Fail(ctx, gin.H{
			"ERROR": err.Error(),
		}, "获取信息失败")
		return
	}
	// 添加资源
	models.DB.Create(&video)
	tools.Success(ctx, gin.H{
		"file": video,
	}, "上传成功")
}

// 上传选择题
func (UploadCon) UploadChoice(ctx *gin.Context) {
	// 接收title,url,pointId,ans,analy
	choice := models.Choice{}
	if err := ctx.ShouldBind(&choice); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	models.DB.Create(&choice)
	tools.Success(ctx, gin.H{
		"choice": choice,
	}, "上传成功")
}

// 上传代码题
func (UploadCon) UploadCode(ctx *gin.Context) {
	// 接收title,code,pointId
	code := models.Code{}
	if err := ctx.ShouldBind(&code); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	models.DB.Create(&code)
	tools.Success(ctx, gin.H{
		"code": code,
	}, "上传成功")
}

// 上传语言
// func (UploadCon) UploadLang(ctx *gin.Context) {
// 	cate := models.ExerciseCate{}
// 	if err := ctx.ShouldBind(&cate); err != nil {
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"err": err.Error(),
// 		})
// 		return
// 	}
// 	models.DB.Create(&cate)
// 	tools.Success(ctx, gin.H{
// 		"cate": cate,
// 	}, "创建成功")
// }

// 上传某个知识点
func (UploadCon) UploadPoint(ctx *gin.Context) {
	// 接收title,courseId
	point := models.ExercisePoint{}
	if err := ctx.ShouldBind(&point); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	models.DB.Create(&point)
	tools.Success(ctx, gin.H{
		"point": point,
	}, "上传成功")
}

func (UploadCon) AddCourse(ctx *gin.Context) {
	// 接收title,key,message
	userInfo, _ := ctx.Get("user")
	user := userInfo.(models.User)
	courseInfo := struct {
		Title   string `json:"title" form:"title"`
		Key     string `json:"key" form:"key"`
		Message string `json:"message" form:"message"`
	}{}
	if err := ctx.ShouldBind(&courseInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	ce := models.CourseExam{
		CourseTitle: courseInfo.Title,
		CourseKey:   courseInfo.Key,
		User:        user,
		Message:     courseInfo.Message,
		Time:        time.Now().Format("2006-01-02 15:04:05"),
	}
	log.Println(ce.Time)
	models.DB.Create(&ce)
	tools.Success(ctx, gin.H{
		"req": ce,
	}, "请求成功，审核中。")
}
