package learn

// import (
// 	"fmt"
// 	"net/http"
// 	"wcapp/models"
// 	"wcapp/tools"

// 	"github.com/gin-gonic/gin"
// )

// type CodeCon struct{}

// func (CodeCon) Cate(ctx *gin.Context) {
// 	CateList := []models.CodeCate{}
// 	if err := models.DB.Find(&CateList).Error; err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// 	tools.Success(ctx, gin.H{
// 		"category": CateList,
// 	}, "代码语言分类")
// }

// func (CodeCon) Code(ctx *gin.Context) {
// 	codeInfo := struct {
// 		CateId int `form:"id"`
// 	}{}
// 	CodeList := []models.CodeCate{}
// 	if err := ctx.ShouldBind(&codeInfo); err == nil {
// 		models.DB.Where("id = ?", codeInfo.CateId).Preload("Codes").Find(&CodeList)
// 		tools.Success(ctx, gin.H{
// 			"codeList": CodeList,
// 		}, "代码补全习题")
// 	} else {
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"err": err.Error(),
// 		})
// 	}
// }
