package common

// ResCode int32
type ResCode int32
type Action int32

const (
	CodeSuccess              ResCode = 0
	CodeInvalidParam         ResCode = 1010
	CodeInvalidLoginInfo     ResCode = 1040
	CodeInvalidLoginUserID   ResCode = 1041
	CodeInvalidLoginPassword ResCode = 1042
	CodeServerBusy           ResCode = 1050
	CodeNeedLogin            ResCode = 1060
	CodeInvalidToken         ResCode = 1070
	CodeRegisterFailed       ResCode = 1080
	CodeIphoneIsExist        ResCode = 1081
	CodeIphoneNotExist       ResCode = 1082
	CodeMysqlFailed          ResCode = 1090
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:              "success",
	CodeInvalidParam:         "请求参数错误",
	CodeInvalidLoginInfo:     "查不到该用户信息",
	CodeInvalidLoginUserID:   "用户id不存在",
	CodeInvalidLoginPassword: "用户密码错误",
	CodeServerBusy:           "服务繁忙",
	CodeNeedLogin:            "需要登录",
	CodeInvalidToken:         "无效的token",
	CodeRegisterFailed:       "注册失败",
	CodeIphoneIsExist:        "手机号已被注册",
	CodeIphoneNotExist:       "手机号未注册",
	CodeMysqlFailed:          "mysql操作错误",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
