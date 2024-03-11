package io

//请求参数

//注册参数
type ParamRegister struct {
	UserName string
	PassWord string
}

//登录参数
type ParamLogin struct {
	UserID   int64
	PassWord string
}

//修改参数
type ParamUpdate struct {
	UserID   int64
	PassWord string
}

//用户信息请求参数
type UserInfoReq struct {
	UserID int64
	Token  string
}
