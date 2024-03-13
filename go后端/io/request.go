package io

//请求参数

//注册参数
type ParamRegister struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	IphoneID string `json:"iphoneID"`
}

//登录参数
type ParamLogin struct {
	UserID   int64  `json:"userID"`
	PassWord string `json:"password"`
	IphoneID string `json:"iphoneID"`
}

//修改参数
type ParamUpdate struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	IphoneID string `json:"iphoneID"`
	Token    string `json:"token"`
}

//找回密码参数
type ParamForgetpwd struct {
	UserID   int64  `json:"userID"`
	IphoneID string `json:"iphoneID"`
}

//用户信息请求参数
type UserInfoReq struct {
	UserID   int64  `json:"userID"`
	UserName string `json:"userName"`
	Token    string `json:"token"`
}
