package service

import (
	"WebVideoServer/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response_Sport(ctx *gin.Context) {
	ctx.String(http.StatusOK, logic.GetSportStr())
}
