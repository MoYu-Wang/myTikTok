package common

// ResCode int32
type ResCode int32
type Action int32

const (
	CodeSuccess                 ResCode = 0
	CodeInvalidParam            ResCode = 1010
	CodeDataTypeChangeError     ResCode = 1011
	CodeGetVideoTagsWeightError ResCode = 1020
	CodeInvalidLoginInfo        ResCode = 1040
	CodeInvalidLoginUserID      ResCode = 1041
	CodeInvalidLoginPassword    ResCode = 1042
	CodeUpdateUserInfoSuccess   ResCode = 1043
	CodeUserNotExist            ResCode = 1044
	CodeServerBusy              ResCode = 1050
	CodeNeedLogin               ResCode = 1060
	CodeCommentNotOwn           ResCode = 1061
	CodeInvalidToken            ResCode = 1070
	CodeRegisterFailed          ResCode = 1080
	CodeUserRegisterSuccess     ResCode = 1081
	CodeIphoneIsExist           ResCode = 1082
	CodeIphoneNotExist          ResCode = 1083
	CodeMysqlFailed             ResCode = 1090
	CodeUserDeleteSuccess       ResCode = 1100
	CodeUserUpLoadVideoSuccess  ResCode = 1101
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:                 "success",
	CodeInvalidParam:            "请求参数错误",
	CodeDataTypeChangeError:     "数据类型转化失败",
	CodeGetVideoTagsWeightError: "获取视频标签权值算法错误",
	CodeInvalidLoginInfo:        "查不到该用户信息",
	CodeInvalidLoginUserID:      "用户id不存在",
	CodeInvalidLoginPassword:    "用户密码错误",
	CodeUpdateUserInfoSuccess:   "修改用户信息成功",
	CodeUserNotExist:            "用户已注销或用户id不存在",
	CodeServerBusy:              "服务繁忙",
	CodeNeedLogin:               "用户需要登录或登录信息已过期",
	CodeCommentNotOwn:           "该评论不属于自己",
	CodeInvalidToken:            "无效的token",
	CodeRegisterFailed:          "注册失败",
	CodeUserRegisterSuccess:     "注册成功",
	CodeIphoneIsExist:           "手机号已被注册",
	CodeIphoneNotExist:          "手机号未注册",
	CodeMysqlFailed:             "mysql操作错误",
	CodeUserDeleteSuccess:       "用户注销成功",
	CodeUserUpLoadVideoSuccess:  "用户上传视频成功",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
