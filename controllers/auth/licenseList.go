package auth

// import (
// 	"wcapp/models"
// 	"wcapp/tools"

// 	"github.com/gin-gonic/gin"
// )

// type LicenseListCon struct{}

// func (LicenseListCon) LicenseList(ctx *gin.Context) {
// 	licenseList := []models.License{}
// 	models.DB.Find(&licenseList)
// 	tools.Success(ctx, gin.H{
// 		"secretKey": licenseList,
// 	}, "现有注册码")
// }
