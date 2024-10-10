package routers

import (
	"wcapp/controllers"

	"github.com/gin-gonic/gin"
)

func RoutersInit(r *gin.Engine) {
	Routers := r.Group("/")
	{
		Routers.GET("/", controllers.Controller{}.Index)
	}
}
