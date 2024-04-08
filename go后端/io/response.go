package io

import (
	"fmt"
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
	UserID   int64  `json:"userID"`
	UserName string `json:"userName"`
	Token    string `json:"token"`
}

type UserRegisterResponse struct {
	Response
	UserID int64 `json:"userID"`
}

type VideoCommentResp struct {
	Response
	VideoComments []VideoComment `json:"videoComments"` //视频所有评论
}

// 用户信息
type UserInfo struct {
	CareCount int64  `json:"careCount"` // 关注总数
	FansCount int64  `json:"fansCount"` // 粉丝总数
	GetLikes  int64  `json:"getLikes"`  //获取点赞数
	ID        int64  `json:"id"`        // 用户id
	IsCare    bool   `json:"isCare"`    // true-已关注，false-未关注
	Name      string `json:"name"`      // 用户名称
}

// 视频信息
type VideoInfo struct {
	VideoID          string `json:"videoID"`          //视频id
	VideoName        string `json:"videoName"`        //视频名称
	VideoTags        string `json:"videoTags"`        //视频标签
	UserID           int64  `json:"userID"`           //视频发布人id
	VideoLink        string `json:"videoLink"`        //视频链接
	VideoFavoriteNum int64  `json:"videoFavoriteNum"` //视频点赞人数
	VideoCommitNum   int64  `json:"videoCommitNum"`   //视频评论人数

	IsFavorite bool `json:"isFavorite"` //当前登录账号是否点赞该视频
}

type VideoComment struct {
	CommentID  string `json:"commentID"`  //评论ID
	UserID     int64  `json:"userID"`     //评论者
	CommitText string `json:"commitText"` //评论文本
	CommitTime int64  `json:"commitTime"` //评论时间
}

type UserBase struct {
	UserID   int64  `json:"userID"`
	UserName string `json:"userName"`
	PassWord string `json:"password"`
	IphoneID string `json:"iphoneID"`
}

type CareUser struct {
	UserID   int64  `json:"userID"`
	UserName string `json:"userName"`
}

type CareListResp struct {
	Response
	CareList []CareUser `json:"careList"`
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
	UserInfo
}

// UserBaseResp 用户基本信息返回值
type UserBaseResp struct {
	Response
	UserBase
}

// UserWorkResp 用户作品返回值
type UserWorkResp struct {
	Response
	VideoInfos []VideoInfo `json:"videoInfos"`
}

// UserHistoryResp 用户观看历史记录返回值
type UserHistoryResp struct {
	Response
	VideoInfos []VideoInfo `json:"videoInfos"`
}

// UserFavoriteResp 用户点赞视频返回值
type UserFavoriteResp struct {
	Response
	VideoInfos []VideoInfo `json:"videoInfos"`
}

// TopVideoResp 热点视频返回值
type TopVideoResp struct {
	Response
	VideoInfos []VideoInfo `json:"videoInfos"`
}

// CareVideoResp 关注视频返回值
type CareVideoResp struct {
	Response
	VideoInfos []VideoInfo `json:"videoInfos"`
}

// RefereeVideoResp 推荐视频返回值
type RefereeVideoResp struct {
	Response
	VideoInfos []VideoInfo `json:"videoInfos"`
}

// SearchVideoResp 搜索视频返回值
type SearchVideoResp struct {
	Response
	VideoInfos []VideoInfo `json:"videoInfos"`
}

// PasswordResp 密码返回值
type PasswordResp struct {
	Response
	Password string
}

// GetSignResp 获取上传签名返回值
type GetSignResp struct {
	Response
	Sign string `json:"mysign"`
}

type CommentVideoResp struct {
	Response
	CommentID string `json:"commentID"`
}

// ResponseError 响应错误
func ResponseError(c *gin.Context, code common.ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Response: Response{code, code.Msg()},
	})
}

// ResponseSuccess 响应成功
func ResponseSuccess(c *gin.Context, code common.ResCode) {
	fmt.Println(code)
	c.JSON(http.StatusOK, &ResponseData{
		Response: Response{code, code.Msg()},
	})

}

func ResponseSuccessRegister(c *gin.Context, userID int64) {
	fmt.Println(userID)
	c.JSON(http.StatusOK, &UserRegisterResponse{
		Response: Response{common.CodeUserRegisterSuccess, common.CodeUserRegisterSuccess.Msg()},
		UserID:   userID,
	})
}

// ResponseSuccessLogin 登录成功
func ResponseSuccessLogin(c *gin.Context, token string) {
	fmt.Println(token)
	c.JSON(http.StatusOK, &UserLoginResponse{
		Response: Response{common.CodeSuccess, common.CodeSuccess.Msg()},
		Token:    token,
	})
}

// ResponseSuccessUserInfo 返回用户信息
func ResponseSuccessUserInfo(c *gin.Context, resp *UserInfoResp) {
	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// ResponseSuccessUserBase 返回用户基本信息
func ResponseSuccessUserBase(c *gin.Context, resp *UserBaseResp) {
	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// 返回用户密码
func ResponseSuccessPassword(c *gin.Context, resp *PasswordResp) {
	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// 返回获取上传视频签名
func ResponseSuccessGetSign(c *gin.Context, resp *GetSignResp) {
	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// 返回用户所有作品
func ResponseSuccessUserWork(c *gin.Context, resp *UserWorkResp) {
	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// 返回用户观看历史记录
func ResponseSuccessUserHistory(c *gin.Context, resp *UserHistoryResp) {
	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// 返回用户点赞视频列表
func ResponseSuccessUserFavorite(c *gin.Context, resp *UserFavoriteResp) {
	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// 返回热点视频列表
func ResponseSuccessTopVideo(c *gin.Context, resp *TopVideoResp) {
	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// 返回关注视频列表
func ResponseSuccessCareVideo(c *gin.Context, resp *CareVideoResp) {
	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// 返回推荐视频列表
func ResponseSuccessRefereeVideo(c *gin.Context, resp *RefereeVideoResp) {
	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// 返回搜索视频列表
func ResponseSuccessSearchVideo(c *gin.Context, resp *SearchVideoResp) {
	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// 返回视频评论id
func ResponseSuccessCommentVideo(c *gin.Context, resp *CommentVideoResp) {
	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// 返回获取视频所有评论
func ResponseSuccessGetVideoComments(c *gin.Context, resp *VideoCommentResp) {
	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// 返回关注列表
func ResponseSuccessCareList(c *gin.Context, resp *CareListResp) {
	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}
