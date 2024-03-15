package service

import (
	"WebVideoServer/web/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TopVideo(ctx *gin.Context) {
	ctx.String(http.StatusOK, logic.GetTopStr())
}
