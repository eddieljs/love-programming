package main

import (
	"github.com/gin-gonic/gin"
	"io"
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
	r.Use(cors)

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

	// routers.CodeRouters(r)
	r.Run(":80")
}
