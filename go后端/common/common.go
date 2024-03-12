package common

import "errors"

const (
	KCtxUserIDKey   = "userID"   // userId 上下文的 userId
	KCtxUserNameKey = "username" // username 上下文的 username
	Kmd5Secret      = "暂时先写在这"   // 用于用户信息加密
)

var (
	ErrorUserNotLogin        = errors.New("用户不存在")
	ErrorPasswordNotOK       = errors.New("密码不合法")
	ErrorUserIDNotExist      = errors.New("用户id不存在")
	ErrorPassword            = errors.New("密码错误")
	ErrorMysql               = errors.New("Mysql错误")
	ErrorIphoneID            = errors.New("手机号位数错误")
	ErrorIphoneIDNotRegister = errors.New("手机号尚未注册")
	ErrorIphoneIDIsRegister  = errors.New("手机号已被注册")
)
