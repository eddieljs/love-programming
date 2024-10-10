package auth

import (
	"net/http"
	"untitled/models"
	"untitled/tools"

	"github.com/gin-gonic/gin"
)

type SpaceCon struct{}

func (SpaceCon) Space(ctx *gin.Context) {
	// fmt.Println("lllllllkkkkkkkk")
	user, _ := ctx.Get("user")
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"code": 200,
	// 	"user": user,
	// })
	tools.Success(ctx, gin.H{
		"username":  user.(models.User).Username,
		"auth":      user.(models.User).Auth,
		"name":      user.(models.User).Name,
		"studentId": user.(models.User).StudentId,
		"avatar":    user.(models.User).Avatar,
	}, "用户基本信息")
}

func (SpaceCon) ChangeInfo(ctx *gin.Context) {
	userInfo, _ := ctx.Get("user")
	cgInfo := struct {
		Username  string `json:"username" form:"username"`
		Name      string `json:"name" form:"name"`
		StudentId string `json:"sId" form:"sId"`
		Avatar    string `json:"avatar" form:"avatar"`
	}{}
	if err := ctx.ShouldBind(&cgInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	user := userInfo.(models.User)
	models.DB.Where("id = ?", user.Id).First(&user)
	user.Name = cgInfo.Name
	user.Username = cgInfo.Username
	user.StudentId = cgInfo.StudentId
	user.Avatar = cgInfo.Avatar
	models.DB.Save(&user)
	tools.Success(ctx, gin.H{
		"info": cgInfo,
	}, "信息修改成功")
}
