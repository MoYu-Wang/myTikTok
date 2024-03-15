package service

import (
	"WebVideoServer/web/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CareVideo(ctx *gin.Context) {
	ctx.String(http.StatusOK, logic.GetCareStr())
}
