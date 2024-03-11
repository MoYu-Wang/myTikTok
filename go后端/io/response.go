package io

import (
	"net/http"

	"WebVideoServer/common"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode common.ResCode `json:"status_code"`
	StatusMsg  string         `json:"status_msg"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type User struct {
	CareCount int64  `json:"care_count"` // 关注总数
	FansCount int64  `json:"fans_count"` // 粉丝总数
	ID        int64  `json:"id"`         // 用户id
	IsCare    bool   `json:"is_care"`    // true-已关注，false-未关注
	Name      string `json:"name"`       // 用户名称
}

// ResponseData 通用的响应内容
type ResponseData struct {
	Response
	Msg  interface{} `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// UserInfoResp 用户信息返回值
type UserInfoResp struct {
	Response
	User User `json:"user"`
}

// ResponseSuccessLogin 登录成功
func ResponseSuccessLogin(c *gin.Context, token string) {
	userId, _ := c.Get("userId")
	c.JSON(http.StatusOK, &UserLoginResponse{
		Response: Response{common.CodeSuccess, common.CodeSuccess.Msg()},
		UserId:   userId.(int64),
		Token:    token,
	})
}

// ResponseError 响应错误
func ResponseError(c *gin.Context, code common.ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Response: Response{code, code.Msg()},
	})
}

// ResponseSuccessUserInfo 返回用户信息
func ResponseSuccessUserInfo(c *gin.Context, resp *UserInfoResp) {
	c.JSON(http.StatusOK, resp)
}
