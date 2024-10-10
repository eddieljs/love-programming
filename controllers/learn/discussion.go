package learn

import (
	"net/http"
	"time"
	"wcapp/models"
	"wcapp/tools"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DiscussionCon struct{}

func (DiscussionCon) Index(ctx *gin.Context) {
	// token := ctx.GetHeader("Authorization")
	// fmt.Printf("token: %v\n", token)
	// 接收课程id
	disList := models.Course{}
	if err := ctx.ShouldBind(&disList); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	}
	models.DB.Where("id = ?", disList.Id).Preload("Discussions", func(db *gorm.DB) *gorm.DB {
		return db.Preload("User")
	}).Find(&disList)
	// ctx.JSON(http.StatusOK, disList)
	tools.Success(ctx, gin.H{
		"discussion": disList.Discussions,
	}, "讨论区内容")
}

func (DiscussionCon) Detail(ctx *gin.Context) {
	disInfo := struct {
		DisId int `form:"id"`
	}{}
	commentList := models.Discussion{}
	if err := ctx.ShouldBind(&disInfo); err == nil {
		models.DB.Where("id = ?", disInfo.DisId).Preload("User").Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User")
		}).First(&commentList)
		// models.DB.Where("id = ?", disInfo.DisId).Preload("Comments").Preload("User").Find(&commentList)
		tools.Success(ctx, gin.H{
			"detail": commentList,
		}, "问题详情")
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	}

}

func (DiscussionCon) PulishDiscussion(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	disInfo := struct {
		Title    string `json:"title" form:"title"`
		Content  string `json:"content" form:"content"`
		CourseId int    `json:"courseId" form:"courseId"`
	}{}
	discussion := models.Discussion{}
	if err := ctx.ShouldBind(&disInfo); err == nil {
		discussion.CourseId = disInfo.CourseId
		discussion.Title = disInfo.Title
		discussion.Content = disInfo.Content
		discussion.UserId = user.(models.User).Id
		discussion.UploadTime = time.Now().Format("2006-01-02 15:04:05")
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	models.DB.Create(&discussion)
	discussion.User = user.(models.User)
	tools.Success(ctx, gin.H{
		"dis": discussion,
	}, "发布成功")
}

func (DiscussionCon) PulishComment(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	commentInfo := struct {
		Content string `json:"content" form:"content"`
		DisId   int    `json:"disId" form:"disId"`
	}{}
	comment := models.Comment{}
	if err := ctx.ShouldBind(&commentInfo); err == nil {
		comment.Content = commentInfo.Content
		comment.DisId = commentInfo.DisId
		comment.UserId = user.(models.User).Id
		comment.UploadTime = time.Now().Format("2006-01-02 15:04:05")
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	models.DB.Create(&comment)
	comment.User = user.(models.User)
	tools.Success(ctx, gin.H{
		"comment": comment,
	}, "发布成功")
}
