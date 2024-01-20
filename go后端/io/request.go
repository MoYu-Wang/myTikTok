package io

import "errors"

//用户未登录错误
var ErrorUserNotLogin = errors.New("user not login")

//请求参数

//注册参数
type ParamRegister struct {
	Username string
	Password string
}

//登录参数
type ParamLogin struct {
	Username string
	Password string
}

//用户信息请求参数
type UserInfoReq struct {
	UID   int64
	Token string
}
