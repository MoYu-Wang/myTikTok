package common

// ResCode int32
type ResCode int32
type Action int32

const (
	CodeSuccess          ResCode = 0
	CodeInvalidParam     ResCode = 1001
	CodeInvalidLoginInfo ResCode = 1004
	CodeServerBusy       ResCode = 1005
	CodeNeedLogin        ResCode = 1006
	CodeInvalidToken     ResCode = 1007
	CodeRegisterFailed   ResCode = 1008
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:          "success",
	CodeInvalidParam:     "请求参数错误",
	CodeInvalidLoginInfo: "查不到该用户信息",
	CodeServerBusy:       "服务繁忙",
	CodeNeedLogin:        "需要登录",
	CodeInvalidToken:     "无效的token",
	CodeRegisterFailed:   "注册失败",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
