package learn

import (
	"net/http"
	"time"
	"untitled/models"
	"untitled/tools"

	"github.com/gin-gonic/gin"
)

type RecordCon struct{}

// 添加学习记录
func (RecordCon) AddRecord(ctx *gin.Context) {
	userInfo, _ := ctx.Get("user")
	user := userInfo.(models.User)
	recordInfo := struct {
		Type  string `json:"type" form:"type"`
		Title string `json:"title" form:"title"`
		Url   string `json:"url" form:"url"`
	}{}
	if err := ctx.ShouldBind(&recordInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	record := models.Record{}
	models.DB.Where("user_id = ? AND url = ?", user.Id, recordInfo.Url).First(&record)
	if record.Id == 0 {
		// 之前没有学习过，那么加上这条记录。
		record.Title = recordInfo.Title
		record.Url = recordInfo.Url
		record.Type = recordInfo.Type
		record.UserId = user.Id
		record.Time = time.Now().Format("2006-01-02 15:04:05")
		models.DB.Create(&record)
		tools.Success(ctx, gin.H{
			"record": record,
		}, "添加记录成功")
	} else {
		// 之前有过记录，更新一下时间。
		record.Time = time.Now().Format(time.DateTime)
		models.DB.Save(&record)
		tools.Success(ctx, gin.H{
			"record": record,
		}, "修改记录成功")
	}

}

// 查看学习记录
func (RecordCon) Record(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	recordList := []models.Record{}
	models.DB.Where("user_id = ?", user.(models.User).Id).Find(&recordList)
	tools.Success(ctx, gin.H{
		"record": recordList,
	}, "学习记录")
}

// 添加做题记录
func (RecordCon) AddExerRecord(ctx *gin.Context) {
	userInfo, _ := ctx.Get("user")
	user := userInfo.(models.User)
	recordInfo := struct {
		ChoiceId int    `json:"choiceId" form:"choiceId"`
		UserAns  string `json:"userAns" form:"userAns"`
	}{}
	if err := ctx.ShouldBind(&recordInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	record := models.ChoiceRecord{}
	models.DB.Where("choice_id = ? AND user_id = ?", recordInfo.ChoiceId, user.Id).First(&record)
	if record.Id == 0 {
		record.ChoiceId = recordInfo.ChoiceId
		record.UserId = user.Id
		record.UserAns = recordInfo.UserAns
		models.DB.Create(&record)
		tools.Success(ctx, gin.H{
			"choiceRecord": record,
		}, "做题记录添加成功")
	} else {
		record.UserAns = recordInfo.UserAns
		models.DB.Save(&record)
		tools.Success(ctx, gin.H{
			"choiceRecord": record,
		}, "做题记录修改成功")
	}

}

// 查看记录
func (RecordCon) ChoiceRecord(ctx *gin.Context) {
	userInfo, _ := ctx.Get("user")
	user := userInfo.(models.User)
	recordInfo := struct {
		ChoiceId int `json:"choiceId" form:"choiceId"`
	}{}
	if err := ctx.ShouldBind(&recordInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	record := models.ChoiceRecord{}
	models.DB.Where("choice_id = ? AND user_id = ?", recordInfo.ChoiceId, user.Id).First(&record)
	if record.Id != 0 {
		tools.Success(ctx, gin.H{
			"record": record,
			"f":      true,
		}, "已做过")
	} else {
		tools.Success(ctx, gin.H{
			"record": nil,
			"f":      false,
		}, "没做过")
	}
}
