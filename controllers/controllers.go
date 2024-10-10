package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (con Controller) Index(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, "/learn")
}
