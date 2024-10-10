package routers

import (
	"wcapp/controllers/learn"
	"wcapp/tools"

	"github.com/gin-gonic/gin"
)

// 学习界面首页
func LearnRouters(r *gin.Engine) {
	Routers := r.Group("/learn")
	{
		Routers.GET("/", learn.LearnController{}.Index)
	}
}

// 课件界面
func ResourceRouters(r *gin.Engine) {
	Routers := r.Group("/learn/resource")
	{
		// 获取资料下的分类
		// /learn/resource/
		// Routers.GET("/", learn.ResourceCon{}.Cate)
		// 获取每个分类的资料
		// /learn/resource/cate?id=
		Routers.GET("/", learn.ResourceCon{}.Resource)
	}
}

// 视频界面
func VideoRouters(r *gin.Engine) {
	Routers := r.Group("/learn/video")
	{
		// 获取视频的分类
		// /learn/video/
		// Routers.GET("/", learn.VideoCon{}.Cate)
		// 获取每个分类的视频
		// /learn/video/cate?id=
		Routers.GET("/", learn.VideoCon{}.Video)
	}
}

// 习题界面
func ExerciseRouters(r *gin.Engine) {
	Routers := r.Group("/learn/exercise")
	{
		// 获取习题分类
		// /learn/exercise/
		// Routers.GET("/", learn.ExerciseCon{}.Cate)
		// 获取每个分类的习题
		// /learn/exercise/cate?id=
		Routers.GET("/", learn.ExerciseCon{}.Point)
		Routers.GET("/point", learn.ExerciseCon{}.Exercise)
	}
}

// 讨论界面
func DiscussionRouters(r *gin.Engine) {
	Routers := r.Group("/learn/discussion")
	{
		// 获取讨论列表
		// /learn/discussion
		Routers.GET("/", learn.DiscussionCon{}.Index)
		Routers.GET("/detail", learn.DiscussionCon{}.Detail)
		Routers.POST("/publish/d", tools.AuthMiddleware, learn.DiscussionCon{}.PulishDiscussion)
		Routers.POST("/publish/c", tools.AuthMiddleware, learn.DiscussionCon{}.PulishComment)
	}
}

// 编译器
func CompileRouters(r *gin.Engine) {
	Routers := r.Group("learn/compile")
	{
		Routers.POST("/", learn.CompilerCon{}.Compile)
	}
}

// 代码补全练习题
// func CodeRouters(r *gin.Engine) {
// 	Routers := r.Group("learn/code")
// 	{
// 		Routers.GET("/", learn.CodeCon{}.Cate)
// 		Routers.GET("/cate", learn.CodeCon{}.Code)
// 	}
// }

// 学习记录
func RecordRouters(r *gin.Engine) {
	Routers := r.Group("/learn/record")
	{
		Routers.POST("/add", tools.AuthMiddleware, learn.RecordCon{}.AddRecord)
		Routers.GET("/list", tools.AuthMiddleware, learn.RecordCon{}.Record)
		Routers.POST("/addChoice", tools.AuthMiddleware, learn.RecordCon{}.AddExerRecord)
		Routers.GET("/choiceList", tools.AuthMiddleware, learn.RecordCon{}.ChoiceRecord)
	}
}

// 课程
func CourseRouters(r *gin.Engine) {
	Routers := r.Group("/learn/course")
	{
		Routers.POST("/select", tools.AuthMiddleware, learn.CourseCon{}.SlctCourse)
		Routers.GET("/list", tools.AuthMiddleware, learn.CourseCon{}.CourseList)
		Routers.GET("/all", tools.AuthMiddleware, learn.CourseCon{}.AllCourse)
		Routers.GET("/teacher", tools.AuthMiddleware, learn.CourseCon{}.TeacherCourse)
	}
}
