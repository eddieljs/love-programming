package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"untitled/routers"
)

//	func handler(w http.ResponseWriter, r *http.Request) {
//		log.Print("收到访问请求！")
//		target := "欢迎使用微信云托管"
//		fmt.Fprintf(w, "Hello, %s!\n", target)
//	}
//
//	func main() {
//		log.Print("微信云托管服务启动成功")
//		// 记录到文件。
//		f, _ := os.Create("gin.log")
//		// 如果需要同时将日志写入文件和控制台，请使用以下代码。
//		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
//		http.HandleFunc("/", handler)
//		port := os.Getenv("PORT")
//		if port == "" {
//			port = "80"
//		}
//		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
//
//		r := gin.Default()
//		// CORS中间件
//		cors := func(c *gin.Context) {
//			// 允许特定的域进行跨域请求
//			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//			// 允许特定的请求方法
//			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
//			// 允许特定的请求头
//			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
//			// 允许携带身份凭证（如Cookie）
//			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
//			if c.Request.Method == http.MethodOptions {
//				c.AbortWithStatus(http.StatusOK)
//			} else {
//				// 继续处理请求
//				c.Next()
//			}
//
//		}
//		r.Use(cors)
//		routers.RoutersInit(r)
//		routers.LearnRouters(r)
//		routers.ResourceRouters(r)
//		routers.VideoRouters(r)
//		routers.ExerciseRouters(r)
//		routers.DiscussionRouters(r)
//		routers.CompileRouters(r)
//		routers.RecordRouters(r)
//		routers.CourseRouters(r)
//		routers.AuthRouters(r)
//		routers.AdminRouters(r)
//
//		r.Run(":80")
//	}
func main() {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	// 创建日志文件
	f, err := os.Create("gin.log")
	if err != nil {
		log.Fatalf("无法创建日志文件: %v", err)
	}
	defer f.Close()

	// 同时将日志写入文件和控制台
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 创建默认的 Gin 路由器
	r := gin.Default()

	// CORS 中间件
	cors := func(c *gin.Context) {
		// 允许所有域进行跨域请求（根据需要调整）
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 允许特定的请求方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 允许特定的请求头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// 允许携带身份凭证（如 Cookie）
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 如果是预检请求，直接返回 200
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// 继续处理请求
		c.Next()
	}
	r.Use(cors)

	// 添加根路径处理器
	r.GET("/", func(c *gin.Context) {
		log.Print("收到访问请求！")
		target := "欢迎使用微信云托管"
		c.String(http.StatusOK, "Hello, %s!\n", target)
	})

	// 初始化您的自定义路由
	routers.RoutersInit(r)
	routers.LearnRouters(r)
	routers.ResourceRouters(r)
	routers.VideoRouters(r)
	routers.ExerciseRouters(r)
	routers.DiscussionRouters(r)
	routers.CompileRouters(r)
	routers.RecordRouters(r)
	routers.CourseRouters(r)
	routers.AuthRouters(r)
	routers.AdminRouters(r)
	// routers.CodeRouters(r) // 如果需要，可以取消注释

	// 获取端口号，默认为 80
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	log.Print("微信云托管服务启动成功")

	// 启动 Gin 服务器
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
