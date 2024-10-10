package tools

import (
	"fmt"
	"net/http"
	"strings"
	"wcapp/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(ctx *gin.Context) {
	// fmt.Print("123123")
	tokenString := ctx.GetHeader("Authorization")
	// fmt.Printf("tokenString: %v\n", tokenString)
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		// ctx.JSON(http.StatusOK, gin.H{
		// 	"code": 401,
		// 	"msg":  "权限不足",
		// })
		Response(ctx, http.StatusOK, 401, gin.H{}, "权限不足")
		ctx.Abort()
		return
	}
	tokenString = tokenString[7:]
	// fmt.Printf("tokenString: %T\n", tokenString)
	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		// fmt.Printf("err: %v\n", err)
		// ctx.JSON(http.StatusOK, gin.H{
		// 	"code": 401,
		// 	"msg":  "解析失败",
		// })
		Response(ctx, http.StatusOK, 401, gin.H{
			"error": err,
		}, "token解析失败")
		ctx.Abort()
		return
	}
	// 验证通过后获取userId
	userId := claims.UserId
	user := models.User{}
	models.DB.Where("id = ?", userId).First(&user)
	if user.Id == 0 {
		// ctx.JSON(http.StatusOK, gin.H{
		// 	"code": 401,
		// 	"msg":  "用户不存在",
		// })
		Response(ctx, http.StatusOK, 401, gin.H{}, "用户不存在")
		ctx.Abort()
		return
	}
	ctx.Set("user", user)
	ctx.Next()
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return token, claims, err
}
