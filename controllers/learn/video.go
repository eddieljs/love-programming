package learn

import (
	"net/http"
	"untitled/models"
	"untitled/tools"

	"github.com/gin-gonic/gin"
)

type VideoCon struct{}

// func (VideoCon) Cate(ctx *gin.Context) {
// 	cateList := []models.VideoCate{}
// 	if err := models.DB.Find(&cateList).Error; err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// 	// ctx.JSON(http.StatusOK, cateList)
// 	tools.Success(ctx, gin.H{
// 		"category": cateList,
// 	}, "教学视频分类")
// }

func (VideoCon) Video(ctx *gin.Context) {
	// 接收课程id。
	videoList := models.Course{}
	if err := ctx.ShouldBind(&videoList); err == nil {
		models.DB.Where("id = ?", videoList.Id).Preload("Videos").First(&videoList)
		// ctx.JSON(http.StatusOK, videoList)
		tools.Success(ctx, gin.H{
			"video": videoList.Videos,
		}, "教学视频")
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	}
}
