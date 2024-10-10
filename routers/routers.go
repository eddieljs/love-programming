package routers

import (
	"github.com/gin-gonic/gin"
	"untitled/controllers"
)

func RoutersInit(r *gin.Engine) {
	Routers := r.Group("/")
	{
		Routers.GET("/", controllers.Controller{}.Index)
	}
}
