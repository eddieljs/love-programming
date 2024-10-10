package learn

import (
	"net/http"
	"wcapp/models"
	"wcapp/tools"

	"github.com/gin-gonic/gin"
)

type ResourceCon struct{}

// func (ResourceCon) Cate(ctx *gin.Context) {
// 	cateList := []models.ResourceCate{}
// 	if err := models.DB.Find(&cateList).Error; err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// 	// ctx.JSON(http.StatusOK, cateList)
// 	tools.Success(ctx, gin.H{
// 		"category": cateList,
// 	}, "学习资源分类")
// }

func (ResourceCon) Resource(ctx *gin.Context) {
	// 接收课程的id
	resourceInfo := struct {
		Id int `json:"id" form:"id"`
	}{}
	resourceList := models.Course{}
	if err := ctx.ShouldBind(&resourceInfo); err == nil {
		// 预加载资料
		models.DB.Where("id = ?", resourceInfo.Id).Preload("Resources").First(&resourceList)
		// ctx.JSON(http.StatusOK, resourceList)
		tools.Success(ctx, gin.H{
			"resource": resourceList.Resources,
		}, "学习资料")
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	}
}
