package auth

import (
	"github.com/gin-gonic/gin"
	"untitled/models"
	"untitled/tools"
)

type RegisterCon struct{}

// func (RegisterCon) Register(ctx *gin.Context) {
// 	userInfo := struct {
// 		Username  string `json:"username" form:"username"`
// 		Password  string `json:"password" form:"password"`
// 		SecretKey string `json:"secretKey" form:"secretKey"`
// 	}{}
// 	if err := ctx.ShouldBind(&userInfo); err != nil {
// 		// ctx.JSON(http.StatusBadRequest, gin.H{
// 		// 	"ERROR": err.Error(),
// 		// })
// 		tools.Fail(ctx, gin.H{
// 			"ERROR": err.Error(),
// 		}, "获取信息失败")
// 	}
// 	fmt.Printf("userInfo: %v\n", userInfo)
// 	user := models.User{}
// 	// 查找数据库中的信息
// 	models.DB.Where("username = ?", userInfo.Username).First(&user)
// 	if user.Id != 0 {
// 		// ctx.JSON(http.StatusOK, gin.H{
// 		// 	"msg": "用户名已存在",
// 		// })
// 		tools.Success(ctx, gin.H{}, "用户名已存在")
// 		return
// 	}
// 	// 检验注册码是不是真的
// 	license := models.License{}
// 	models.DB.Where("secret_key=?", userInfo.SecretKey).First(&license)
// 	if license.Id == 0 {
// 		// 假的
// 		// ctx.JSON(http.StatusOK, gin.H{
// 		// 	"msg": "注册码错误",
// 		// })
// 		tools.Response(ctx, http.StatusOK, 401, gin.H{}, "注册码错误")
// 		return
// 	} else {
// 		// 真的，并且删除。
// 		models.DB.Delete(&license)
// 	}
// 	// 都没问题，可以注册。
// 	user.Username = userInfo.Username
// 	user.Password = userInfo.Password
// 	user.Auth = 1
// 	if err := models.DB.Create(&user).Error; err != nil {
// 		fmt.Println("数据添加失败！")
// 		// fmt.Printf("err: %v\n", err)
// 		tools.Fail(ctx, gin.H{
// 			"username": userInfo.Username,
// 			"password": userInfo.Password,
// 			"error":    err,
// 		}, "异常")
// 	} else {
// 		fmt.Println("数据添加成功！")
// 	}
// 	// ctx.JSON(http.StatusOK, gin.H{
// 	// 	"msg": "注册成功",
// 	// })
// 	tools.Success(ctx, gin.H{
// 		"username": userInfo.Username,
// 		"password": userInfo.Password,
// 	}, "注册成功")
// }

func (RegisterCon) TeacherRegister(ctx *gin.Context) {
	userInfo, _ := ctx.Get("user")
	user := userInfo.(models.User)
	UserInfo := struct {
		// Id  int    `json:"id" form:"id"`
		Key string `json:"key" form:"key"`
	}{}
	if err := ctx.ShouldBind(&UserInfo); err != nil {
		tools.Fail(ctx, gin.H{
			"ERROR": err.Error(),
		}, "获取信息失败")
		return
	}
	// user := models.User{}
	// models.DB.Where("id = ?", UserInfo.Id).First(&user)
	if user.Auth >= 7 {
		// 已经是管理员了
		tools.Success(ctx, gin.H{
			"user": user,
		}, "已经是老师了")
		return
	}
	// 不是管理员，验证一下激活码。
	license := models.License{}
	models.DB.Where("secret_key = ?", UserInfo.Key).First(&license)
	if license.Id == 0 {
		// 假的激活码
		tools.Fail(ctx, gin.H{
			"user": user,
			"key":  UserInfo.Key,
		}, "激活码错误")
		return
	}
	// 激活码对了，修改权限,还要删除激活码。
	user.Auth = 7
	tx := models.DB.Begin()
	tx.Where("id = ?", license.Id).Delete(&license)
	tx.Save(&user)
	tx.Commit()
	tools.Success(ctx, gin.H{
		"user": user,
	}, "教师权限激活成功")
}

func (RegisterCon) VIPRegister(ctx *gin.Context) {
	userInfo, _ := ctx.Get("user")
	user := userInfo.(models.User)
	vipKeyInfo := models.VIPKey{}
	if err := ctx.ShouldBind(&vipKeyInfo); err != nil {
		tools.Fail(ctx, gin.H{
			"ERROR": err.Error(),
		}, "获取信息失败")
		return
	}
	// 获取到vipKeyInfo，在数据库里找一下是不是真的。
	vipKey := models.VIPKey{}
	models.DB.Where("account = ? AND password = ?", vipKeyInfo.Account, vipKeyInfo.Password).First(&vipKey)
	// 如果是假的
	if vipKey.Id == 0 {
		tools.Fail(ctx, gin.H{
			"user":   user,
			"vipKey": vipKeyInfo,
		}, "激活码错误")
		return
	}
	// 如果是真的
	user.Auth = 2
	tx := models.DB.Begin()
	tx.Where("id = ?", vipKey.Id).Delete(&vipKey)
	tx.Save(&user)
	tx.Commit()
	tools.Success(ctx, gin.H{
		"user": user,
	}, "会员权限激活成功")
}

// func (RegisterCon) AdminRegister(ctx *gin.Context) {

// }
