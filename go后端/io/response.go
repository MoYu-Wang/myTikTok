package io

import (
	"net/http"

	"WebVideoServer/common"
	"WebVideoServer/dao"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode common.ResCode `json:"status_code"`
	StatusMsg  string         `json:"status_msg"`
}

type UserLoginResponse struct {
	Response
	UserID   int64  `json:"userID"`
	UserName string `json:"userName"`
	Token    string `json:"token"`
}

type UserRegisterResponse struct {
	Response
	UserID int64 `json:"userID"`
}

type UserInfo struct {
	CareCount int64  `json:"careCount"` // 关注总数
	FansCount int64  `json:"fansCount"` // 粉丝总数
	ID        int64  `json:"id"`        // 用户id
	IsCare    bool   `json:"isCare"`    // true-已关注，false-未关注
	Name      string `json:"name"`      // 用户名称
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
	UserInfo UserInfo `json:"userInfo"`
}

// UserBaseResp 本用户基本信息返回值
type UserBaseResp struct {
	Response
	User dao.User `json:"user"`
}

type PasswordResp struct {
	Response
	Password string
}

// ResponseError 响应错误
func ResponseError(c *gin.Context, code common.ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Response: Response{code, code.Msg()},
	})
}

// ResponseSuccess 响应成功
func ResponseSuccess(c *gin.Context, code common.ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Response: Response{code, code.Msg()},
	})
}

func ResponseSuccessRegister(c *gin.Context, userID int64) {
	c.JSON(http.StatusOK, &UserRegisterResponse{
		Response: Response{common.CodeUserRegisterSuccess, common.CodeUserRegisterSuccess.Msg()},
		UserID:   userID,
	})
}

// ResponseSuccessLogin 登录成功
func ResponseSuccessLogin(c *gin.Context, token string) {
	userID, _ := c.Get("UserID")
	userName, _ := c.Get("UserName")
	c.JSON(http.StatusOK, &UserLoginResponse{
		Response: Response{common.CodeSuccess, common.CodeSuccess.Msg()},
		UserID:   userID.(int64),
		UserName: userName.(string),
		Token:    token,
	})
}

// ResponseSuccessUserInfo 返回用户信息
func ResponseSuccessUserInfo(c *gin.Context, resp *UserInfoResp) {
	c.JSON(http.StatusOK, resp)
}

// ResponseSuccessUserBase 返回用户基本信息
func ResponseSuccessUserBase(c *gin.Context, resp *UserBaseResp) {
	c.JSON(http.StatusOK, resp)
}

// 返回用户密码
func ResponseSuccessPassword(c *gin.Context, resp *PasswordResp) {
	c.JSON(http.StatusOK, resp)
}
