package main

import (
	"io"
	"net/http"
	"os"
	"untitled/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	// 记录到文件。
	f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	// CORS中间件
	cors := func(c *gin.Context) {
		// 允许特定的域进行跨域请求
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 允许特定的请求方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 允许特定的请求头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// 允许携带身份凭证（如Cookie）
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
		} else {
			// 继续处理请求
			c.Next()
		}

	}
	print()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
