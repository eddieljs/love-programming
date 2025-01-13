package auth

import (
	"errors"
	"gorm.io/gorm"
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

// 新增课程
//
//	func (UploadCon) AddCourse(ctx *gin.Context) {
//		// 接收title,key,message
//		userInfo, _ := ctx.Get("user")
//		user := userInfo.(models.User)
//		courseInfo := struct {
//			Title   string `json:"title" form:"title"`
//			Key     string `json:"key" form:"key"`
//			Message string `json:"message" form:"message"`
//		}{}
//		if err := ctx.ShouldBind(&courseInfo); err != nil {
//			ctx.JSON(http.StatusOK, gin.H{
//				"err": err.Error(),
//			})
//			return
//		}
//		// 设置时区为中国时区 (Asia/Shanghai)
//		location, err := time.LoadLocation("Asia/Shanghai")
//		if err != nil {
//			// 如果无法加载时区，返回错误
//			ctx.JSON(http.StatusOK, gin.H{
//				"err": "无法加载时区",
//			})
//			return
//		}
//		// 获取当前时间并设置为中国时区
//		currentTime := time.Now().In(location).Format("2006-01-02 15:04:05")
//		ce := models.CourseExam{
//			CourseTitle: courseInfo.Title,
//			CourseKey:   courseInfo.Key,
//			User:        user,
//			Message:     courseInfo.Message,
//			Time:        currentTime,
//		}
//
//		models.DB.Create(&ce)
//		tools.Success(ctx, gin.H{
//			"req": ce,
//		}, "请求成功，审核中。")
//	}
func (UploadCon) AddCourse(ctx *gin.Context) {
	// 接收 title, key, course_time
	userInfo, _ := ctx.Get("user")
	user := userInfo.(models.User)
	courseInfo := struct {
		Title      string `json:"title" form:"title"`
		Key        string `json:"key" form:"key"`
		CourseTime string `json:"course_time" form:"course_time"` // 上课时间
	}{}
	if err := ctx.ShouldBind(&courseInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	// 验证 key是否已经存在
	var existingCourse models.CourseExam
	if err := models.DB.Where("course_key = ? ", courseInfo.Key).First(&existingCourse).Error; err == nil {
		// 如果 key 已经存在，返回错误
		tools.Fail(ctx, gin.H{}, "选课码已存在，请使用其他选课码")
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果查询出错，返回错误
		tools.Fail(ctx, gin.H{}, "查询选课码失败: "+err.Error())
		return
	}
	// 设置时区为中国时区 (Asia/Shanghai)
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		// 如果无法加载时区，返回错误
		tools.Fail(ctx, gin.H{}, "无法加载时区")
		return
	}
	// 获取当前时间并设置为中国时区
	currentTime := time.Now().In(location).Format("2006-01-02 15:04:05")

	// 创建课程审核信息
	ce := models.CourseExam{
		CourseTitle: courseInfo.Title,
		CourseKey:   courseInfo.Key,
		CourseTime:  courseInfo.CourseTime, // 上课时间
		User:        user,
		Time:        currentTime,
	}

	// 保存到数据库
	if err := models.DB.Create(&ce).Error; err != nil {
		tools.Fail(ctx, gin.H{}, "创建课程失败: "+err.Error())
		return
	}

	// 返回成功响应
	tools.Success(ctx, gin.H{
		"req": ce,
	}, "请求成功，审核中。")
}
