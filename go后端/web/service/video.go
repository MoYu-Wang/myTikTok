package service

import (
	"WebVideoServer/common"
	"WebVideoServer/io"
	"WebVideoServer/jwt"
	"WebVideoServer/web/logic"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 上传视频
func UpLoadVideo(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.UserUpLoadVideoReq)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	fmt.Printf("请求参数:")
	fmt.Println(p)
	//登录校验,解析Token里的参数
	claim, err := jwt.ParseToken(p.Token)
	if err != nil {
		fmt.Println("token解析失败")
		io.ResponseError(ctx, common.CodeNeedLogin)
		return
	}
	//判断token解析出来的用户信息是否正确
	if code := logic.UserIsExist(ctx, claim); code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	//2.服务调用
	if code := logic.UpLoadVideo(ctx, p, claim); code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}

	//3.返回成功响应
	io.ResponseSuccess(ctx, common.CodeUserUpLoadVideoSuccess)
}

// 获取签名
func GetSign(ctx *gin.Context) {
	//1.获取参数和参数校验
	token := ctx.DefaultQuery("token", "")
	//登录校验,解析Token里的参数
	claim, err := jwt.ParseToken(token)
	if err != nil {
		fmt.Println("token解析失败")
		io.ResponseError(ctx, common.CodeNeedLogin)
		return
	}
	//判断token解析出来的用户信息是否正确
	if code := logic.UserIsExist(ctx, claim); code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	//2.服务调用
	//获取签名
	mysign := logic.GetSign()
	fmt.Println("mysign:" + mysign)

	resp := io.GetSignResp{
		Response: io.Response{StatusCode: 0, StatusMsg: "success"},
		Sign:     mysign,
	}
	//3.返回成功响应
	io.ResponseSuccessGetSign(ctx, &resp)
}

// 划走视频后对视频的操作
func OperateVideo(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.OperateVideoReq)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		//请求参数有误,直接返回响应
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	fmt.Println("请求参数:")
	fmt.Println(p)
	//登录校验
	claim, _ := jwt.ParseToken(p.Token)
	//2.服务调用
	code := logic.OperateVideo(ctx, p, claim)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	//3.返回响应
	io.ResponseSuccess(ctx, common.CodeSuccess)
}

// 获取热点视频
func TopVideo(ctx *gin.Context) {
	//1.获取参数和参数校验
	token := ctx.DefaultQuery("token", "")
	var claim *jwt.MyClaims
	var userID int64
	if token != "" {
		//登录校验,解析Token里的参数
		var err error
		claim, err = jwt.ParseToken(token)
		if err != nil {
			io.ResponseError(ctx, common.CodeNeedLogin)
			return
		}
		userID = claim.UserID
	} else {
		userID = 0
	}
	//2.服务调用
	vids, code := logic.GetTopVideoIDs(ctx, userID)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	var videoInfos []io.VideoInfo
	for _, vid := range vids {
		videoInfo, code := logic.GetVideoInfoByVID(ctx, vid, userID)
		if code != common.CodeSuccess {
			io.ResponseError(ctx, code)
			return
		}
		videoInfos = append(videoInfos, *videoInfo)
	}
	resp := &io.TopVideoResp{
		Response:   io.Response{StatusCode: 0, StatusMsg: "success"},
		VideoInfos: videoInfos,
	}
	//3.返回成功响应
	io.ResponseSuccessTopVideo(ctx, resp)
}

// 获取推荐视频
func RefereeVideo(ctx *gin.Context) {
	//1.获取参数和参数校验
	token := ctx.DefaultQuery("token", "")
	//登录校验,解析Token里的参数
	claim, err := jwt.ParseToken(token)
	if err != nil {
		io.ResponseError(ctx, common.CodeNeedLogin)
		return
	}
	//判断token解析出来的用户信息是否正确
	if code := logic.UserIsExist(ctx, claim); code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	//2.服务调用
	vids, code := logic.GetRefereeVideoIDs(ctx, claim)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	var videoInfos []io.VideoInfo
	for _, vid := range vids {
		videoInfo, code := logic.GetVideoInfoByVID(ctx, vid, claim.UserID)
		if code != common.CodeSuccess {
			io.ResponseError(ctx, code)
			return
		}
		videoInfos = append(videoInfos, *videoInfo)
	}
	resp := &io.RefereeVideoResp{
		Response:   io.Response{StatusCode: 0, StatusMsg: "success"},
		VideoInfos: videoInfos,
	}
	//3.返回成功响应
	io.ResponseSuccessRefereeVideo(ctx, resp)
}

// 获取关注视频
func CareVideo(ctx *gin.Context) {
	//1.获取参数和参数校验
	token := ctx.DefaultQuery("token", "")
	//登录校验,解析Token里的参数
	claim, err := jwt.ParseToken(token)
	if err != nil {
		io.ResponseError(ctx, common.CodeNeedLogin)
		return
	}
	//判断token解析出来的用户信息是否正确
	if code := logic.UserIsExist(ctx, claim); code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	//2.服务调用
	vids, code := logic.GetCareVideoIDs(ctx, claim)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	var videoInfos []io.VideoInfo
	for _, vid := range vids {
		videoInfo, code := logic.GetVideoInfoByVID(ctx, vid, claim.UserID)
		if code != common.CodeSuccess {
			io.ResponseError(ctx, code)
			return
		}
		videoInfos = append(videoInfos, *videoInfo)
	}
	resp := &io.CareVideoResp{
		Response:   io.Response{StatusCode: 0, StatusMsg: "success"},
		VideoInfos: videoInfos,
	}
	//3.返回成功响应
	io.ResponseSuccessCareVideo(ctx, resp)
}

// 模糊查询视频
func SearchVideo(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.SearchVideoReq)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		//请求参数有误,直接返回响应
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	fmt.Println("请求参数:")
	fmt.Println(p)
	//登录校验
	claim, _ := jwt.ParseToken(p.Token)
	//2.服务调用
	vids, code := logic.GetSearchVideoIDs(ctx, p.SearchText)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	var videoInfos []io.VideoInfo
	for _, vid := range vids {
		videoInfo, code := logic.GetVideoInfoByVID(ctx, vid, claim.UserID)
		if code != common.CodeSuccess {
			io.ResponseError(ctx, code)
			return
		}
		videoInfos = append(videoInfos, *videoInfo)
	}
	resp := &io.SearchVideoResp{
		Response:   io.Response{StatusCode: 0, StatusMsg: "success"},
		VideoInfos: videoInfos,
	}
	//3.返回成功响应
	io.ResponseSuccessSearchVideo(ctx, resp)
}

// 获取视频评论
func GetVideoComment(ctx *gin.Context) {
	//1.获取参数和参数校验
	videoID, err := strconv.ParseInt(ctx.DefaultQuery("videoID", ""), 10, 64)
	if err != nil {
		fmt.Println("转换失败:", err)
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	//2.服务调用
	comments, code := logic.GetVideoComment(ctx, videoID)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	resp := &io.VideoCommentResp{
		Response:      io.Response{StatusCode: 0, StatusMsg: "success"},
		VideoComments: comments,
	}
	//3.返回成功响应
	io.ResponseSuccessGetVideoComments(ctx, resp)
}

// 视频点赞
func FavoriteVideo(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.FavoriteVideoReq)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	fmt.Printf("请求参数:")
	fmt.Println(p)
	//登录校验,解析Token里的参数
	claim, err := jwt.ParseToken(p.Token)
	if err != nil {
		fmt.Println("token解析失败")
		io.ResponseError(ctx, common.CodeNeedLogin)
		return
	}
	//判断token解析出来的用户信息是否正确
	if code := logic.UserIsExist(ctx, claim); code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	//2.服务调用
	if code := logic.FavoriteVideo(ctx, p, claim); code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}

	//3.返回成功响应
	io.ResponseSuccess(ctx, common.CodeSuccess)
}

// 评论视频
func CommentVideo(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.CommentVideoReq)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	fmt.Printf("请求参数:")
	fmt.Println(p)
	//登录校验,解析Token里的参数
	claim, err := jwt.ParseToken(p.Token)
	if err != nil {
		fmt.Println("token解析失败")
		io.ResponseError(ctx, common.CodeNeedLogin)
		return
	}
	//判断token解析出来的用户信息是否正确
	if code := logic.UserIsExist(ctx, claim); code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	//2.服务调用
	commentID, code := logic.CommentVideo(ctx, p, claim)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	resp := &io.CommentVideoResp{
		Response:  io.Response{StatusCode: 0, StatusMsg: "success"},
		CommentID: strconv.FormatInt(commentID, 10),
	}
	//3.返回成功响应
	io.ResponseSuccessCommentVideo(ctx, resp)
}

// 删除视频评论
func DeleteVideoComment(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.DeleteCommentReq)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	fmt.Printf("请求参数:")
	fmt.Println(p)
	//登录校验,解析Token里的参数
	claim, err := jwt.ParseToken(p.Token)
	if err != nil {
		fmt.Println("token解析失败")
		io.ResponseError(ctx, common.CodeNeedLogin)
		return
	}
	//判断token解析出来的用户信息是否正确
	if code := logic.UserIsExist(ctx, claim); code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	//2.服务调用
	if code := logic.DeleteComment(ctx, p, claim); code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}

	//3.返回成功响应
	io.ResponseSuccess(ctx, common.CodeSuccess)
}

func VideoOperateInfo(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.VideoOperateInfoReq)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	fmt.Printf("请求参数:")
	fmt.Println(p)
	//登录校验
	var userID int64
	claim, err := jwt.ParseToken(p.Token)
	if err != nil {
		userID = 0
	} else {
		userID = claim.UserID
	}
	videoID, err := strconv.ParseInt(p.VideoID, 10, 64)
	if err != nil {
		io.ResponseError(ctx, common.CodeDataTypeChangeError)
	}
	//2.服务调用
	videoOperateInfo, code := logic.GetVideoOperateInfo(ctx, userID, videoID)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	resp := &io.VideoOperateInfoResp{
		Response:         io.Response{StatusCode: 0, StatusMsg: "success"},
		VideoOperateInfo: *videoOperateInfo,
	}
	//3.返回响应
	io.ResponseSuccessVideoOperate(ctx, resp)
}
