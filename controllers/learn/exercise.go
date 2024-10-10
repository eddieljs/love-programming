package learn

import (
	"net/http"
	"wcapp/models"
	"wcapp/tools"

	"github.com/gin-gonic/gin"
)

type ExerciseCon struct{}

// func (ExerciseCon) Cate(ctx *gin.Context) {
// 	cateList := []models.ExerciseCate{}
// 	if err := models.DB.Find(&cateList).Error; err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// 	// ctx.JSON(http.StatusOK, cateList)
// 	tools.Success(ctx, gin.H{
// 		"category": cateList,
// 	}, "习题语言分类")
// }

func (ExerciseCon) Point(ctx *gin.Context) {
	pointList := models.Course{}
	if err := ctx.ShouldBind(&pointList); err == nil {
		models.DB.Where("id = ?", pointList.Id).Preload("Points").First(&pointList)
		// ctx.JSON(http.StatusOK, exerciseList)
		tools.Success(ctx, gin.H{
			"point": pointList.Points,
		}, "知识点")
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	}

}

func (ExerciseCon) Exercise(ctx *gin.Context) {
	exerciseInfo := struct {
		// CateId  int `form:"cid"`
		PointId int `form:"id"`
	}{}
	choiceList := models.ExercisePoint{}
	codeList := models.ExercisePoint{}
	if err := ctx.ShouldBind(&exerciseInfo); err == nil {
		models.DB.Where("id = ?", exerciseInfo.PointId).Preload("Choices").First(&choiceList)
		models.DB.Where("id = ?", exerciseInfo.PointId).Preload("Codes").First(&codeList)
		tools.Success(ctx, gin.H{
			"choice": choiceList.Choices,
			"code":   codeList.Codes,
		}, "练习题")
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	}
}
