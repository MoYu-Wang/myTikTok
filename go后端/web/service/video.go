package service

import (
	"WebVideoServer/common"
	"WebVideoServer/io"
	"WebVideoServer/jwt"
	"WebVideoServer/web/logic"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 上传视频
func UpLoadVideo(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.UserUpLoadVideoReq)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	fmt.Printf("请求参数:")
	fmt.Println(p)
	//登录校验,解析Token里的参数
	claim, err := jwt.ParseToken(p.Token)
	if err != nil {
		fmt.Println("token解析失败")
		io.ResponseError(ctx, common.CodeNeedLogin)
		return
	}
	//2.服务调用
	if code := logic.UpLoadVideo(ctx, p, claim); code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}

	//3.返回成功响应
	io.ResponseSuccess(ctx, common.CodeUserUpLoadVideoSuccess)
}

// 获取签名
func GetSign(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.UserBaseReq)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	fmt.Printf("请求参数:")
	fmt.Println(p)
	//登录校验,解析Token里的参数
	_, err := jwt.ParseToken(p.Token)
	if err != nil {
		fmt.Println("token解析失败")
		io.ResponseError(ctx, common.CodeNeedLogin)
		return
	}
	//2.服务调用
	//获取签名
	mysign := logic.GetSign()
	fmt.Println("mysign:" + mysign)

	resp := io.GetSignResp{
		Response: io.Response{StatusCode: 0, StatusMsg: "success"},
		Sign:     mysign,
	}
	//3.返回成功响应
	io.ResponseSuccessGetSign(ctx, &resp)
}

// 评论视频
func CommitVideo(ctx *gin.Context) {

}

// 划走视频后对视频的操作
func OperateVideo(Ctx *gin.Context) {

}

// 获取热点视频
func TopVideo(ctx *gin.Context) {
}

// 获取推荐视频
func RefereeVideo(ctx *gin.Context) {

}

// 获取关注视频
func CareVideo(ctx *gin.Context) {

}
