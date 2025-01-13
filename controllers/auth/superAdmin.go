// package auth
//
// import (
//
//	"net/http"
//	"time"
//	"untitled/models"
//	"untitled/tools"
//
//	"github.com/gin-gonic/gin"
//
// )
//
// type SuperAdminCon struct{}
//
//	func (SuperAdminCon) Exam(ctx *gin.Context) {
//		// 接收id,f。
//		ceInfo := struct {
//			Id int    `json:"id" form:"id"`
//			F  string `json:"f" form:"f"`
//		}{}
//		if err := ctx.ShouldBind(&ceInfo); err != nil {
//			ctx.JSON(http.StatusOK, gin.H{
//				"err": err.Error(),
//			})
//			return
//		}
//		ce := models.CourseExam{}
//		models.DB.Where("id = ?", ceInfo.Id).First(&ce)
//		cu := models.CourseUser{
//			UserId: ce.UserId,
//		}
//		course := models.Course{
//			Title:     ce.CourseTitle,
//			Key:       ce.CourseKey,
//			Time:      time.Now().Format("2006-01-02 15:04:05"),
//			TeacherId: ce.UserId,
//		}
//		if ceInfo.F == "y" {
//			tx := models.DB.Begin()
//			tx.Where("id = ?", ce.Id).Delete(&ce)
//			tx.Create(&course)
//			cu.CourseId = course.Id
//			tx.Create(&cu)
//			tx.Commit()
//			tools.Success(ctx, gin.H{
//				"course": course,
//			}, "审核通过")
//		} else if ceInfo.F == "n" {
//			models.DB.Where("id = ?", ce.Id).Delete(&ce)
//			tools.Success(ctx, gin.H{
//				"course": course,
//			}, "审核不通过")
//		}
//	}
//
//	func (SuperAdminCon) CourseList(ctx *gin.Context) {
//		ceList := []models.CourseExam{}
//		models.DB.Preload("User").Find(&ceList)
//		tools.Success(ctx, gin.H{
//			"list": ceList,
//		}, "申请列表")
//	}
package auth

import (
	"net/http"
	"untitled/models"
	"untitled/tools"

	"github.com/gin-gonic/gin"
)

type SuperAdminCon struct{}

//func (SuperAdminCon) Exam(ctx *gin.Context) {
//	// 接收id,f。
//	ceInfo := struct {
//		Id int    `json:"id" form:"id"`
//		F  string `json:"f" form:"f"`
//	}{}
//	if err := ctx.ShouldBind(&ceInfo); err != nil {
//		ctx.JSON(http.StatusOK, gin.H{
//			"err": err.Error(),
//		})
//		return
//	}
//
//	// 查询课程申请信息
//	ce := models.CourseExam{}
//	if err := models.DB.Where("id = ?", ceInfo.Id).First(&ce).Error; err != nil {
//		ctx.JSON(http.StatusOK, gin.H{
//			"err": "未找到该课程申请",
//		})
//		return
//	}
//
//	// 根据审核结果处理
//	if ceInfo.F == "y" {
//		// 审核通过：创建课程并关联教师
//		course := models.Course{
//			Title:     ce.CourseTitle,
//			Key:       ce.CourseKey,
//			Time:      ce.CourseTime,
//			TeacherId: ce.UserId,
//		}
//
//		tx := models.DB.Begin()
//		if err := tx.Create(&course).Error; err != nil {
//			tx.Rollback()
//			ctx.JSON(http.StatusOK, gin.H{
//				"err": "创建课程失败",
//			})
//			return
//		}
//
//		// 删除课程申请记录
//		if err := tx.Where("id = ?", ce.Id).Delete(&models.CourseExam{}).Error; err != nil {
//			tx.Rollback()
//			ctx.JSON(http.StatusOK, gin.H{
//				"err": "删除课程申请失败",
//			})
//			return
//		}
//
//		tx.Commit()
//		tools.Success(ctx, gin.H{
//			"course": course,
//		}, "审核通过")
//	} else if ceInfo.F == "n" {
//		// 审核不通过：删除课程申请记录
//		if err := models.DB.Where("id = ?", ce.Id).Delete(&models.CourseExam{}).Error; err != nil {
//			ctx.JSON(http.StatusOK, gin.H{
//				"err": "删除课程申请失败",
//			})
//			return
//		}
//		tools.Success(ctx, gin.H{}, "审核不通过")
//	} else {
//		ctx.JSON(http.StatusOK, gin.H{
//			"err": "无效的审核参数",
//		})
//	}
//}

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

	// 查询课程申请信息
	ce := models.CourseExam{}
	if err := models.DB.Where("id = ?", ceInfo.Id).First(&ce).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": "未找到该课程申请",
		})
		return
	}

	// 根据审核结果处理
	if ceInfo.F == "y" {
		// 检查是否已存在相同 key 的课程
		var existingCourse models.Course
		if err := models.DB.Where("`key` = ?", ce.CourseKey).First(&existingCourse).Error; err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"err": "已存在相同 key 的课程",
			})
			return
		}

		// 审核通过：创建课程并关联教师
		course := models.Course{
			Title:     ce.CourseTitle,
			Key:       ce.CourseKey,
			Time:      ce.CourseTime,
			TeacherId: ce.UserId,
		}

		tx := models.DB.Begin()
		if err := tx.Create(&course).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusOK, gin.H{
				"err": "创建课程失败",
			})
			return
		}

		// 删除课程申请记录
		if err := tx.Where("id = ?", ce.Id).Delete(&models.CourseExam{}).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusOK, gin.H{
				"err": "删除课程申请失败",
			})
			return
		}

		tx.Commit()
		tools.Success(ctx, gin.H{
			"course": course,
		}, "审核通过")
	} else if ceInfo.F == "n" {
		// 审核不通过：删除课程申请记录
		if err := models.DB.Where("id = ?", ce.Id).Delete(&models.CourseExam{}).Error; err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"err": "删除课程申请失败",
			})
			return
		}
		tools.Success(ctx, gin.H{}, "审核不通过")
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"err": "无效的审核参数",
		})
	}
}

func (SuperAdminCon) CourseList(ctx *gin.Context) {
	// 定义查询结果的结构体
	type CourseExamInfo struct {
		Id          int    `json:"id"`
		CourseTitle string `json:"course_title"`
		CourseKey   string `json:"course_key"`
		CourseTime  string `json:"course_time"`
		TeacherName string `json:"teacher_name"`
	}

	// 查询课程申请列表
	var ceList []CourseExamInfo
	models.DB.Table("course_exam").
		Select("course_exam.id, course_exam.course_title, course_exam.course_key, course_exam.course_time AS course_time, user.name AS teacher_name").
		Joins("LEFT JOIN user ON course_exam.user_id = user.id").
		Find(&ceList)

	// 返回课程申请列表
	tools.Success(ctx, gin.H{
		"list": ceList,
	}, "申请列表")
}
