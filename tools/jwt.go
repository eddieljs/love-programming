package tools

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"untitled/models"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId int
	jwt.RegisteredClaims
}

func ReleaseToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(100 * 60 * time.Second)
	claims := Claims{
		UserId: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			// 这个Issuer可能要放进配置文件中。
			Issuer:  "faruzan.cn",
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
