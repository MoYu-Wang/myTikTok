package service

import (
	"WebVideoServer/web/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response_Shopping(ctx *gin.Context) {
	ctx.String(http.StatusOK, logic.GetShoppingStr())
}
