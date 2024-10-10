package auth

// import (
// 	"crypto/sha256"
// 	"encoding/base64"
// 	"fmt"
// 	"math/rand"
// 	"time"
// 	"wcapp/models"
// 	"wcapp/tools"

// 	"github.com/gin-gonic/gin"
// )

// type GetLicenseCon struct{}

// func (GetLicenseCon) GetLicense(ctx *gin.Context) {
// 	linceseInfo := struct {
// 		Num int    `json:"num" form:"num"`
// 		Key string `json:"key" form:"key"`
// 	}{}
// 	if err := ctx.ShouldBind(&linceseInfo); err != nil {
// 		tools.Fail(ctx, gin.H{
// 			"ERROR": err.Error(),
// 		}, "获取信息失败")
// 	}
// 	// fmt.Printf("linceseInfo: %v\n", linceseInfo)
// 	secretKey := CreateSK(linceseInfo.Num, linceseInfo.Key)
// 	tools.Success(ctx, gin.H{
// 		"SK": secretKey,
// 	}, "注册码生成成功")
// }

// func CreateSK(num int, key string) []string {
// 	var result []string
// 	for i := 0; i < num; i++ {
// 		// 生成随机盐值
// 		rand.New(rand.NewSource(time.Now().UnixNano()))
// 		salt := fmt.Sprintf("%d", rand.Intn(2024))

// 		// 将输入字符串和随机盐值拼接后进行MD5加密
// 		hs := sha256.Sum256([]byte(key + salt))
// 		// fmt.Printf("hs: %v\n", hs)
// 		encode := base64.StdEncoding.EncodeToString(hs[:])
// 		license := models.License{
// 			SecretKey: encode,
// 		}
// 		if err := models.DB.Create(&license).Error; err != nil {
// 			fmt.Println("数据添加失败！")
// 			fmt.Printf("err: %v\n", err)
// 		}
// 		result = append(result, encode)
// 	}
// 	return result
// }
