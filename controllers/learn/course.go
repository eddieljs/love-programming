package learn

import (
	"fmt"
	"net/http"
	"untitled/models"
	"untitled/tools"

	"github.com/gin-gonic/gin"
)

type CourseCon struct{}

//func (CourseCon) SlctCourse(ctx *gin.Context) {
//	userInfo, _ := ctx.Get("user")
//	user := userInfo.(models.User)
//	// courseId,key
//	cuInfo := struct {
//		CourseId int    `json:"courseId" form:"courseId"`
//		Key      string `json:"key" form:"key"`
//	}{}
//	if err := ctx.ShouldBind(&cuInfo); err != nil {
//		ctx.JSON(http.StatusOK, gin.H{
//			"err": err.Error(),
//		})
//		return
//	}
//	cu := models.CourseUser{}
//	models.DB.Where("course_id = ? AND user_id = ?", cuInfo.CourseId, user.Id).First(&cu)
//	if cu.Id == 0 {
//		// 没有选。
//		course := models.Course{}
//		models.DB.Where("id = ?", cuInfo.CourseId).First(&course)
//		if cuInfo.Key == course.Key {
//			//  加课码正确
//			cu.CourseId = cuInfo.CourseId
//			cu.UserId = user.Id
//			models.DB.Create(&cu)
//			tools.Success(ctx, gin.H{}, "选课成功")
//		} else {
//			// 加课码错误
//			tools.Fail(ctx, gin.H{}, "加课码错误")
//		}
//	} else {
//		tools.Fail(ctx, gin.H{}, "已经选过了")
//	}
//}

func (CourseCon) SlctCourse(ctx *gin.Context) {
	userInfo, _ := ctx.Get("user")
	user := userInfo.(models.User)

	// 课程ID和加课码
	cuInfo := struct {
		CourseId int    `json:"courseId" form:"courseId"`
		Key      string `json:"key" form:"key"`
	}{}

	// 打印请求体
	body, _ := ctx.GetRawData()
	fmt.Println("Request Body:", string(body))

	// 绑定 JSON 数据
	if err := ctx.ShouldBindJSON(&cuInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}

	fmt.Println("Received courseId:", cuInfo.CourseId, "key:", cuInfo.Key) // 打印接收到的 courseId 和 key

	// 查找加课码对应的教师
	courseTeacher := models.CourseTeacher{}
	models.DB.Where("course_id = ? AND key = ?", cuInfo.CourseId, cuInfo.Key).First(&courseTeacher)

	if courseTeacher.Id == 0 {
		// 加课码不正确
		tools.Fail(ctx, gin.H{}, "加课码错误")
		return
	}

	// 查询学生是否已经选过该课程与教师
	cu := models.CourseUser{}
	models.DB.Where("course_id = ? AND user_id = ? AND teacher_id = ?", cuInfo.CourseId, user.Id, courseTeacher.TeacherId).First(&cu)

	if cu.Id == 0 {
		// 学生没有选过此课程
		cu.CourseId = cuInfo.CourseId
		cu.UserId = user.Id
		cu.TeacherId = courseTeacher.TeacherId
		models.DB.Create(&cu)

		tools.Success(ctx, gin.H{}, "选课成功")
	} else {
		// 已经选过该课程
		tools.Fail(ctx, gin.H{}, "已经选过该课程")
	}
}

//	func (CourseCon) CourseList(ctx *gin.Context) {
//		userInfo, _ := ctx.Get("user")
//		user := userInfo.(models.User)
//		courseList := models.User{}
//		models.DB.Where("id = ?", user.Id).Preload("Courses").Find(&courseList)
//		tools.Success(ctx, gin.H{
//			"courses": courseList.Courses,
//		}, "已选课程")
//	}
func (CourseCon) CourseList(ctx *gin.Context) {
	userInfo, _ := ctx.Get("user")
	user := userInfo.(models.User)

	// 查询已选课程及其详细信息
	var courses []struct {
		CourseId    int    `json:"course_id"`
		CourseTitle string `json:"course_title"`
		CourseTime  string `json:"course_time"`
		TeacherName string `json:"teacher_name"`
	}

	models.DB.Table("course_user").
		Joins("JOIN course ON course_user.course_id = course.id").
		Joins("JOIN teacher ON course.teacher_id = teacher.id").
		Where("course_user.user_id = ?", user.Id).
		Select("course_user.course_id, course.title AS course_title, course.time AS course_time, teacher.name AS teacher_name").
		Scan(&courses)

	tools.Success(ctx, gin.H{
		"courses": courses,
	}, "已选课程")
}

func (CourseCon) AllCourse(ctx *gin.Context) {
	courseList := []models.Course{}
	models.DB.Find(&courseList)
	tools.Success(ctx, gin.H{
		"courses": courseList,
	}, "所有课程")
}

// // 查看一个教师的课程
//
//	func (CourseCon) TeacherCourse(ctx *gin.Context) {
//		teacherInfo := struct {
//			Id int `json:"id" form:"id"`
//		}{}
//		if err := ctx.ShouldBind(&teacherInfo); err != nil {
//			ctx.JSON(http.StatusOK, gin.H{
//				"err": err.Error(),
//			})
//			return
//		}
//		courseList := models.User{}
//		models.DB.Where("id = ?", teacherInfo.Id).Preload("Courses", func(db *gorm.DB) *gorm.DB {
//			return db.Preload("Teacher")
//		}).First(&courseList)
//		tools.Success(ctx, gin.H{
//			"courses": courseList.Courses,
//		}, "开课列表")
//	}
//
//	func (CourseCon) TeacherCourse(ctx *gin.Context) {
//		// 绑定教师 ID
//		teacherInfo := struct {
//			Id int `json:"id" form:"id"`
//		}{}
//		if err := ctx.ShouldBind(&teacherInfo); err != nil {
//			tools.Fail(ctx, gin.H{}, "参数绑定失败: "+err.Error())
//			return
//		}
//
//		// 检查教师 ID 是否有效
//		if teacherInfo.Id == 0 {
//			tools.Fail(ctx, gin.H{}, "教师 ID 不能为空")
//			return
//		}
//
//		// 查询教师的所有课程
//		var courses []struct {
//			Title string `json:"title"` // 课程名
//			Key   string `json:"key"`   // 选课码
//			Time  string `json:"time"`  // 上课时间
//		}
//		err := models.DB.Table("course").
//			Where("teacher_id = ?", teacherInfo.Id).
//			Select("title, `key`, time"). // 使用反引号包裹 key
//			Find(&courses).Error
//
//		if err != nil {
//			tools.Fail(ctx, gin.H{}, "查询失败: "+err.Error())
//			return
//		}
//
//		// 检查是否查询到课程
//		if len(courses) == 0 {
//			tools.Fail(ctx, gin.H{}, "该教师暂无课程")
//			return
//		}
//
//		// 返回课程列表
//		tools.Success(ctx, gin.H{
//			"courses": courses,
//		}, "开课列表")
//	}
func (CourseCon) TeacherCourse(ctx *gin.Context) {
	// 绑定教师 ID
	teacherInfo := struct {
		Id int `json:"id" form:"id"`
	}{}
	if err := ctx.ShouldBind(&teacherInfo); err != nil {
		tools.Fail(ctx, gin.H{}, "参数绑定失败: "+err.Error())
		return
	}

	// 检查教师 ID 是否有效
	if teacherInfo.Id == 0 {
		tools.Fail(ctx, gin.H{}, "教师 ID 不能为空")
		return
	}

	// 查询教师的所有课程
	var courses []struct {
		ID    int    `json:"id"`    // 课程 ID
		Title string `json:"title"` // 课程名
		Key   string `json:"key"`   // 选课码
		Time  string `json:"time"`  // 上课时间
	}
	err := models.DB.Table("course").
		Where("teacher_id = ?", teacherInfo.Id).
		Select("id, title, `key`, time"). // 添加 id 字段
		Find(&courses).Error

	if err != nil {
		tools.Fail(ctx, gin.H{}, "查询失败: "+err.Error())
		return
	}

	// 检查是否查询到课程
	if len(courses) == 0 {
		tools.Fail(ctx, gin.H{}, "该教师暂无课程")
		return
	}

	// 返回课程列表
	tools.Success(ctx, gin.H{
		"courses": courses,
	}, "开课列表")
}

// 根据选课码查询课程信息
//func (CourseCon) SearchCourseByKey(ctx *gin.Context) {
//	// 获取选课码
//	key := ctx.Query("key")
//	if key == "" {
//		tools.Fail(ctx, gin.H{}, "选课码不能为空")
//		return
//	}
//
//	// 查询课程信息
//	var course models.Course
//	//var teacher models.Teacher
//	err := models.DB.Table("course").
//		Joins("JOIN teacher ON course.teacher_id = teacher.id").
//		Where("course.key = ?", key).
//		Select("course.id, course.title, course.time, teacher.name AS teacher_name").
//		Scan(&course).Error
//
//	if err != nil {
//		tools.Fail(ctx, gin.H{}, "查询失败")
//		return
//	}
//
//	if course.Id == 0 {
//		tools.Fail(ctx, gin.H{}, "未找到相关课程")
//		return
//	}
//
//	tools.Success(ctx, gin.H{
//		"course": course,
//	}, "查询成功")
//}

// 根据选课码查询课程信息（包含教师名称）
func (CourseCon) SearchCourseByKey(ctx *gin.Context) {
	// 获取选课码
	key := ctx.Query("key")
	if key == "" {
		tools.Fail(ctx, gin.H{}, "选课码不能为空")
		return
	}

	// 查询课程信息和教师名称
	var result struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Time        string `json:"time"`
		TeacherId   int    `json:"teacher_id"`
		TeacherName string `json:"teacher_name"` // 教师名称
	}
	err := models.DB.Table("course").
		Joins("JOIN teacher ON course.teacher_id = teacher.id").
		Where("course.key = ?", key).
		Select("course.id, course.title, course.time, course.teacher_id, teacher.name AS teacher_name").
		Scan(&result).Error

	if err != nil {
		tools.Fail(ctx, gin.H{}, "查询失败")
		return
	}

	if result.Id == 0 {
		tools.Fail(ctx, gin.H{}, "未找到相关课程")
		return
	}

	// 返回课程信息（包含教师名称）
	tools.Success(ctx, gin.H{
		"course": result,
	}, "查询成功")
}

// 学生加入课程
func (CourseCon) JoinCourse(ctx *gin.Context) {
	userInfo, _ := ctx.Get("user")
	user := userInfo.(models.User)

	// 获取课程ID
	var req struct {
		CourseId int `json:"courseId"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		tools.Fail(ctx, gin.H{}, "参数错误")
		return
	}

	// 查询课程信息
	var course models.Course
	if err := models.DB.Where("id = ?", req.CourseId).First(&course).Error; err != nil {
		tools.Fail(ctx, gin.H{}, "课程不存在")
		return
	}

	// 检查是否已经加入课程
	var courseUser models.CourseUser
	err := models.DB.Where("course_id = ? AND user_id = ?", req.CourseId, user.Id).First(&courseUser).Error
	if err == nil {
		tools.Fail(ctx, gin.H{}, "已经加入该课程")
		return
	}

	// 插入选课记录
	courseUser = models.CourseUser{
		CourseId:   req.CourseId,
		UserId:     user.Id,
		CourseTime: course.Time,      // 课程时间
		TeacherId:  course.TeacherId, // 教师 ID
	}
	if err := models.DB.Create(&courseUser).Error; err != nil {
		tools.Fail(ctx, gin.H{}, "加入失败")
		return
	}

	tools.Success(ctx, gin.H{}, "加入成功")
}
