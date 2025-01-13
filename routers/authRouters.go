package routers

import (
	"github.com/gin-gonic/gin"
	"untitled/controllers/auth"
	"untitled/tools"
)

func AuthRouters(r *gin.Engine) {
	Routers := r.Group("/user")
	{
		Routers.POST("/login", auth.LoginCon{}.OpenIdLogin)
		// Routers.POST("/register", auth.RegisterCon{}.Register)
		Routers.GET("/space", tools.AuthMiddleware, auth.SpaceCon{}.Space)
		Routers.POST("/changeInfo", tools.AuthMiddleware, auth.SpaceCon{}.ChangeInfo)
		Routers.POST("/active", tools.AuthMiddleware, auth.RegisterCon{}.VIPRegister)
		Routers.POST("/getUserInfo", tools.AuthMiddleware, auth.RegisterCon{}.GetUserInfo) // 获取用户信息（姓名、学号、班级等）
	}
}

func AdminRouters(r *gin.Engine) {
	Routers := r.Group("/admin")
	{
		// Routers.POST("/key", tools.AuthMiddleware, tools.AdminAuth, auth.GetLicenseCon{}.GetLicense)
		// Routers.GET("/keyList", tools.AuthMiddleware, tools.AdminAuth, auth.LicenseListCon{}.LicenseList)
		Routers.POST("/active/t", tools.AuthMiddleware, auth.RegisterCon{}.TeacherRegister)
		Routers.POST("/upload/p", tools.AuthMiddleware, tools.AdminAuth, auth.UploadCon{}.UploadResource)
		Routers.POST("/upload/v", tools.AuthMiddleware, tools.AdminAuth, auth.UploadCon{}.UploadVideo)
		Routers.POST("/delete/p", tools.AuthMiddleware, tools.AdminAuth, auth.DeleteCon{}.DeleteResource)
		Routers.POST("/delete/v", tools.AuthMiddleware, tools.AdminAuth, auth.DeleteCon{}.DeleteVideo)
		Routers.POST("/upload/choice", tools.AuthMiddleware, tools.AdminAuth, auth.UploadCon{}.UploadChoice)
		Routers.POST("/upload/code", tools.AuthMiddleware, tools.AdminAuth, auth.UploadCon{}.UploadCode)
		Routers.POST("/delete/choice", tools.AuthMiddleware, tools.AdminAuth, auth.DeleteCon{}.DeleteChoice)
		Routers.POST("/delete/code", tools.AuthMiddleware, tools.AdminAuth, auth.DeleteCon{}.DeleteCode)
		Routers.POST("/upload/point", tools.AuthMiddleware, tools.AdminAuth, auth.UploadCon{}.UploadPoint)
		Routers.POST("/delete/point", tools.AuthMiddleware, tools.AdminAuth, auth.DeleteCon{}.DeletePoint)
		Routers.POST("/add/course", tools.AuthMiddleware, tools.AdminAuth, auth.UploadCon{}.AddCourse)
		// Routers.POST("/upload/lang", tools.AuthMiddleware, tools.AdminAuth, auth.UploadCon{}.UploadLang)
		Routers.POST("/super/login", auth.LoginCon{}.AdminLogin)
		Routers.POST("/super/exam", tools.AuthMiddleware, tools.SuperAdminAuth, auth.SuperAdminCon{}.Exam)
		Routers.GET("/super/courseList", tools.AuthMiddleware, tools.SuperAdminAuth, auth.SuperAdminCon{}.CourseList)
	}
}
