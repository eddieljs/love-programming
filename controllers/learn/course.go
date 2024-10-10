package learn

import (
	"net/http"
	"wcapp/models"
	"wcapp/tools"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CourseCon struct{}

func (CourseCon) SlctCourse(ctx *gin.Context) {
	userInfo, _ := ctx.Get("user")
	user := userInfo.(models.User)
	// courseId,key
	cuInfo := struct {
		CourseId int    `json:"courseId" form:"courseId"`
		Key      string `json:"key" form:"key"`
	}{}
	if err := ctx.ShouldBind(&cuInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	cu := models.CourseUser{}
	models.DB.Where("course_id = ? AND user_id = ?", cuInfo.CourseId, user.Id).First(&cu)
	if cu.Id == 0 {
		// 没有选。
		course := models.Course{}
		models.DB.Where("id = ?", cuInfo.CourseId).First(&course)
		if cuInfo.Key == course.Key {
			//  加课码正确
			cu.CourseId = cuInfo.CourseId
			cu.UserId = user.Id
			models.DB.Create(&cu)
			tools.Success(ctx, gin.H{}, "选课成功")
		} else {
			// 加课码错误
			tools.Fail(ctx, gin.H{}, "加课码错误")
		}
	} else {
		tools.Fail(ctx, gin.H{}, "已经选过了")
	}
}

func (CourseCon) CourseList(ctx *gin.Context) {
	userInfo, _ := ctx.Get("user")
	user := userInfo.(models.User)
	courseList := models.User{}
	models.DB.Where("id = ?", user.Id).Preload("Courses").Find(&courseList)
	tools.Success(ctx, gin.H{
		"courses": courseList.Courses,
	}, "已选课程")
}

func (CourseCon) AllCourse(ctx *gin.Context) {
	courseList := []models.Course{}
	models.DB.Find(&courseList)
	tools.Success(ctx, gin.H{
		"courses": courseList,
	}, "所有课程")
}

// 查看一个教师的课程
func (CourseCon) TeacherCourse(ctx *gin.Context) {
	teacherInfo := struct {
		Id int `json:"id" form:"id"`
	}{}
	if err := ctx.ShouldBind(&teacherInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	courseList := models.User{}
	models.DB.Where("id = ?", teacherInfo.Id).Preload("Courses", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Teacher")
	}).First(&courseList)
	tools.Success(ctx, gin.H{
		"courses": courseList.Courses,
	}, "开课列表")
}
