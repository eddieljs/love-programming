package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wcapp/models"
	"wcapp/tools"

	"github.com/gin-gonic/gin"
)

type LoginCon struct{}

// func (LoginCon) Login(ctx *gin.Context) {
// 	// print("aaaaa")
// 	userInfo := struct {
// 		Username string `json:"username" form:"username"`
// 		Password string `json:"password" form:"password"`
// 	}{}
// 	if err := ctx.ShouldBind(&userInfo); err != nil {
// 		// ctx.JSON(http.StatusBadRequest, gin.H{
// 		// 	"ERROR": err.Error(),
// 		// })
// 		tools.Fail(ctx, gin.H{
// 			"ERROR": err.Error(),
// 		}, "获取信息失败")
// 	}
// 	// fmt.Printf("userInfo: %v\n", userInfo)
// 	// ctx.JSON(http.StatusOK, userInfo)
// 	user := models.User{}
// 	models.DB.Where("username = ?", userInfo.Username).First(&user)
// 	if user.Id == 0 {
// 		// 用户不存在
// 		// ctx.JSON(http.StatusOK, gin.H{
// 		// 	"message": "登录失败",
// 		// 	"error":   "用户不存在",
// 		// })
// 		tools.Response(ctx, http.StatusOK, 401, gin.H{}, "用户不存在")
// 		return
// 	} else {
// 		if user.Password != userInfo.Password {
// 			// 密码错误
// 			// ctx.JSON(http.StatusOK, gin.H{
// 			// 	"message": "登录失败",
// 			// 	"error":   "密码错误",
// 			// })
// 			tools.Response(ctx, http.StatusOK, 401, gin.H{}, "密码错误")
// 			return
// 		}
// 	}
// 	// 登录成功
// 	// ctx.SetCookie("user", user.Username, 3600, "/", "localhost", false, false)
// 	token, err := tools.ReleaseToken(user)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"code": 500,
// 			"msg":  "系统异常",
// 		})
// 		log.Printf("token get error:%v", err)
// 		return
// 	}
// 	// ctx.JSON(http.StatusOK, gin.H{
// 	// 	"code":    200,
// 	// 	"message": "登录成功",
// 	// 	"token":   token,
// 	// })
// 	tools.Success(ctx, gin.H{
// 		"token": token,
// 	}, "登录成功")
// }

// 用微信的openid自动登录
func (LoginCon) OpenIdLogin(ctx *gin.Context) {
	resCodeInfo := struct {
		Code string `form:"code" json:"code"`
	}{}

	if err := ctx.ShouldBind(&resCodeInfo); err != nil {
		tools.Fail(ctx, gin.H{
			"ERROR": err.Error(),
		}, "获取code失败")
		print("获得code===")
		return
	}
	openid, err := getOpenid(resCodeInfo.Code)
	print("openid为===" + openid)
	if err != nil {
		tools.Fail(ctx, gin.H{
			"ERROR": err.Error(),
		}, "获取openid失败")
		return
	}
	user := models.User{}
	/*models.DB.Where("openid = ?", openid).First(&user)
	if user.Id == 0 {
		// 用户不存在，帮他注册一下。
		print("用户不存在")
		user.Auth = 1
		user.VipTime = 0
		user.Openid = openid
		if err := models.DB.Create(&user).Error; err != nil {
			fmt.Println("数据添加失败！")
			// fmt.Printf("err: %v\n", err)
			tools.Fail(ctx, gin.H{
				"error": err,
			}, "异常")
		} else {
			fmt.Println("数据添加成功！")
		}
	}*/

	result := models.DB.Where("openid = ?", openid).First(&user)
	if result.Error != nil {
		// 处理查询错误
		tools.Fail(ctx, gin.H{"error": result.Error.Error()}, "数据库查询异常")
	} else if result.RowsAffected == 0 {
		// 用户不存在，帮他注册一下。
		print("用户不存在")
		user.Auth = 1
		user.VipTime = 0
		user.Openid = openid
		if err := models.DB.Create(&user).Error; err != nil {
			fmt.Println("数据添加失败！", err)
			tools.Fail(ctx, gin.H{"error": err}, "异常")
		} else {
			fmt.Println("数据添加成功！")
		}
	} else {
		// 用户存在
		fmt.Println("用户已存在")
	}
	// 登录成功返回token。
	token, err := tools.ReleaseToken(user)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		log.Printf("token get error:%v", err)
		return
	}
	tools.Success(ctx, gin.H{
		"token":     token,
		"openid":    user.Openid,
		"id":        user.Id,
		"auth":      user.Auth,
		"username":  user.Username,
		"studentId": user.StudentId,
		"name":      user.Name,
		"avatar":    user.Avatar,
	}, "登录成功"+user.StudentId)

}

// 获取openid
func getOpenid(code string) (string, error) {
	print("====" + code)
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=wx48df8e6699587719&secret=448464f1823039f96726c66e428cc373&js_code=" + code + "&grant_type=authorization_code"
	print("===" + url)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var data map[string]any
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}
	return data["openid"].(string), nil
}

// 超级管理员登录
func (LoginCon) AdminLogin(ctx *gin.Context) {
	aaInfo := struct {
		Account  string `json:"account" form:"account"`
		Password string `json:"password" form:"password"`
	}{}
	if err := ctx.ShouldBind(&aaInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	// 查找account
	aa := models.AdminAccount{}
	models.DB.Where("account = ?", aaInfo.Account).First(&aa)
	if aa.Id == 0 {
		// 账号不存在
		tools.Fail(ctx, gin.H{}, "账号不存在")
		return
	}
	if aaInfo.Password != aa.Password {
		// 密码错误
		tools.Fail(ctx, gin.H{}, "密码错误")
		return
	}
	// 登录成功，返回token。
	user := models.User{
		Id:   -1,
		Auth: 9,
	}
	token, err := tools.ReleaseToken(user)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		log.Printf("token get error:%v", err)
		return
	}
	tools.Success(ctx, gin.H{
		"token":    token,
		"openid":   user.Openid,
		"id":       user.Id,
		"username": user.Username,
	}, "登录成功")
}
