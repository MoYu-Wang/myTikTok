package service

import (
	"WebVideoServer/web/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response_DBc(ctx *gin.Context) {
	ctx.String(http.StatusOK, logic.GetDBcStr())
}
