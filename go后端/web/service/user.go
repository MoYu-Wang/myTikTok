package service

import (
	"WebVideoServer/common"
	"WebVideoServer/io"
	"WebVideoServer/jwt"
	"WebVideoServer/web/logic"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 用户注册
func UserRegister(ctx *gin.Context) {
	//1.获取参数和参数校验
	//绑定Query参数
	p := new(io.ParamRegister)
	//针对GET method 的操作
	if err := ctx.ShouldBindWith(p, binding.Form); err != nil {
		//请求参数有误,直接返回相应
		zap.L().Error("register with invalid param", zap.Error(err))
		//判断err是不是validationError类型
		errors := err.(validator.ValidationErrors)
		if errors != nil {
			//返回参数错误相应
			io.ResponseError(ctx, common.CodeInvalidParam)
			return
		}
		return
	}

	//2.服务调用
	userID, err := logic.UserRegister(ctx, p)
	if err != nil {
		zap.L().Error("register failed", zap.Error(err))
		io.ResponseError(ctx, common.CodeRegisterFailed)
		return
	}

	//自动登录,获取Token
	user := &io.ParamLogin{
		UserID:   userID,
		PassWord: p.PassWord,
	}
	token, err := logic.UserLogin(ctx, user)
	if err != nil {
		io.ResponseError(ctx, common.CodeInvalidLoginInfo)
		return
	}
	ctx.Set("UserID", userID)

	//4.返回成功响应
	io.ResponseSuccessLogin(ctx, token)
}

// 用户登录
func UserLogin(ctx *gin.Context) {
	//1.获取参数和参数校验
	//绑定Query参数
	p := new(io.ParamLogin)
	if err := ctx.ShouldBindWith(p, binding.Form); err != nil {
		//请求参数有误,直接返回相应
		zap.L().Error("register with invalid param", zap.Error(err))
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}

	//2.服务调用
	token, err := logic.UserLogin(ctx, p)
	if err != nil {
		io.ResponseError(ctx, common.CodeInvalidLoginInfo)
		return
	}
	ctx.Set("UserID", p.UserID)

	//3.返回成功响应
	io.ResponseSuccessLogin(ctx, token)
}

// 获取用户信息
func UserInfo(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.UserInfoReq)
	//绑定Query参数
	if err := ctx.ShouldBindWith(p, binding.Form); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("get user info invalid param", zap.Error(err))
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	//登录校验,解析Token里的参数
	claim, err := jwt.ParseToken(p.Token)
	if err != nil {
		io.ResponseError(ctx, common.CodeNeedLogin)
		return
	}
	//2.服务调用
	userResp, err := logic.GetUserInfo(ctx, p, claim)
	if err != nil {
		io.ResponseError(ctx, common.CodeInvalidLoginInfo)
		return
	}
	resp := io.UserInfoResp{
		Response: io.Response{StatusCode: 0, StatusMsg: "success"},
		User:     *userResp,
	}

	//3.返回成功响应
	io.ResponseSuccessUserInfo(ctx, &resp)

}
