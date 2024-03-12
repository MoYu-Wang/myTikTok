package service

import (
	"WebVideoServer/common"
	"WebVideoServer/io"
	"WebVideoServer/jwt"
	"WebVideoServer/web/logic"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 用户注册
func UserRegister(ctx *gin.Context) {
	//1.获取参数和参数校验
	//绑定Query参数
	p := new(io.ParamRegister)
	//针对GET method 的操作
	if err := ctx.ShouldBindJSON(&p); err != nil {
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

	fmt.Println(p)

	//2.服务调用
	if err := logic.UserRegister(ctx, p); err != common.CodeSuccess {
		io.ResponseError(ctx, err)
		return
	}
	userID, _ := ctx.Get("UserID")

	//自动登录,获取Token
	user := &io.ParamLogin{
		UserID:   userID.(int64),
		PassWord: p.PassWord,
		IphoneID: p.IphoneID,
	}
	token, err := logic.UserIDLogin(ctx, user)
	if err != common.CodeSuccess {
		io.ResponseError(ctx, err)
		return
	}
	//4.返回成功响应
	io.ResponseSuccessLogin(ctx, token)
}

// 用户登录
func UserLogin(ctx *gin.Context) {
	//1.获取参数和参数校验
	//绑定Query参数
	p := new(io.ParamLogin)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		//请求参数有误,直接返回相应
		zap.L().Error("register with invalid param", zap.Error(err))
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	//2.服务调用
	//判断是根据哪个ID登录
	var token string
	var err common.ResCode
	if p.IphoneID != "" {
		var userID int64
		token, userID, err = logic.IphoneIDLogin(ctx, p)
		if err != common.CodeSuccess {
			io.ResponseError(ctx, err)
			return
		}
		ctx.Set("UserID", userID)
	} else {
		token, err = logic.UserIDLogin(ctx, p)
		if err != common.CodeSuccess {
			io.ResponseError(ctx, err)
			return
		}
		ctx.Set("UserID", p.UserID)
	}

	//3.返回成功响应
	io.ResponseSuccessLogin(ctx, token)
}

// 获取用户信息
func UserInfo(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.UserInfoReq)
	//绑定Query参数
	if err := ctx.ShouldBindJSON(&p); err != nil {
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
	//获取用户信息
	userResp, err := logic.GetUserInfo(ctx, p, claim)
	if err != nil {
		io.ResponseError(ctx, common.CodeInvalidLoginInfo)
		return
	}
	resp := io.UserInfoResp{
		Response: io.Response{StatusCode: 0, StatusMsg: "success"},
		UserInfo: *userResp,
	}

	//3.返回成功响应
	io.ResponseSuccessUserInfo(ctx, &resp)
}

// 获取本用户基本信息
func UserBase(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.UserInfoReq)
	//绑定Query参数
	if err := ctx.ShouldBindJSON(&p); err != nil {
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
	//获取本用户基本信息
	userBase, code := logic.GetUser(ctx, claim)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	resp := io.UserBaseResp{
		Response: io.Response{StatusCode: 0, StatusMsg: "success"},
		User:     *userBase,
	}
	//3.返回成功响应
	io.ResponseSuccessUserBase(ctx, &resp)
}

// 修改用户基本信息
func UserUpdate(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.ParamUpdate)
	//绑定Query参数
	if err := ctx.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
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
	//更新用户基本信息
	ret, code := logic.UpdateUserBase(ctx, p, claim)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	resp := io.UserBaseResp{
		Response: io.Response{StatusCode: 0, StatusMsg: "success"},
		User:     *ret,
	}
	//3.返回成功响应
	io.ResponseSuccessUserBase(ctx, &resp)
}

// 找回密码
func PasswordForget(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.ParamForgetpwd)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	//未来加上短信验证

	//2.服务调用
	ret, code := logic.QueryPassword(ctx, p)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
	}
	resp := io.PasswordResp{
		Response: io.Response{StatusCode: 0, StatusMsg: "success"},
		Password: ret,
	}
	//3.返回成功响应
	io.ResponseSuccessPassword(ctx, &resp)
}
