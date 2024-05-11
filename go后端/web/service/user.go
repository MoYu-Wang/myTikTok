package service

import (
	"WebVideoServer/common"
	"WebVideoServer/io"
	"WebVideoServer/jwt"
	"WebVideoServer/web/logic"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 用户注册
func UserRegister(ctx *gin.Context) {
	//1.获取参数和参数校验
	//绑定Query参数
	p := new(io.ParamRegister)
	//针对GET method 的操作
	if err := ctx.ShouldBindJSON(&p); err != nil {
		//请求参数有误,直接返回相应
		zap.L().Error("register with invalid param", zap.Error(err))
		//判断err是不是validationError类型
		errors := err.(validator.ValidationErrors)
		if errors != nil {
			//返回参数错误相应
			io.ResponseError(ctx, common.CodeInvalidParam)
			return
		}
		return
	}

	fmt.Println("请求参数:")
	fmt.Println(p)

	//2.服务调用
	if err := logic.UserRegister(ctx, p); err != common.CodeSuccess {
		io.ResponseError(ctx, err)
		return
	}
	userID, _ := ctx.Get("UserID")

	//3.返回成功响应
	io.ResponseSuccessRegister(ctx, userID.(int64))
}

// 用户登录
func UserLogin(ctx *gin.Context) {
	//1.获取参数和参数校验
	//绑定Query参数
	p := new(io.ParamLogin)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		//请求参数有误,直接返回相应
		zap.L().Error("register with invalid param", zap.Error(err))
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	fmt.Println("请求参数:")
	fmt.Println(p)
	//2.服务调用
	//判断是根据哪个ID登录
	var token string
	var code common.ResCode
	if p.IphoneID != "" {
		token, code = logic.IphoneIDLogin(ctx, p)
		if code != common.CodeSuccess {
			io.ResponseError(ctx, code)
			return
		}
	} else {
		token, code = logic.UserIDLogin(ctx, p)
		if code != common.CodeSuccess {
			io.ResponseError(ctx, code)
			return
		}
	}
	//3.返回成功响应
	io.ResponseSuccessLogin(ctx, token)
}

// 获取用户信息
func UserInfo(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.UserInfoReq)
	//绑定Query参数
	if err := ctx.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("get user info invalid param", zap.Error(err))
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	fmt.Println("请求参数:")
	fmt.Println(p)
	//登录校验,解析Token里的参数
	claim, _ := jwt.ParseToken(p.Token)
	//2.服务调用
	//获取用户信息
	userResp, err := logic.GetUserInfo(ctx, p, claim)
	if err != nil {
		io.ResponseError(ctx, common.CodeInvalidLoginInfo)
		return
	}
	resp := io.UserInfoResp{
		Response: io.Response{StatusCode: 0, StatusMsg: "success"},
		UserInfo: *userResp,
	}

	//3.返回成功响应
	io.ResponseSuccessUserInfo(ctx, &resp)
}

// 获取本用户基本信息
func UserBase(ctx *gin.Context) {
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
	//获取本用户基本信息
	userBase, code := logic.GetUser(ctx, claim)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	resp := io.UserBaseResp{
		Response: io.Response{StatusCode: 0, StatusMsg: "success"},
		UserBase: *userBase,
	}
	//3.返回成功响应
	io.ResponseSuccessUserBase(ctx, &resp)
}

// 修改用户基本信息
func UserUpdateInfo(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.ParamUpdate)
	//绑定Query参数
	if err := ctx.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	fmt.Println("请求参数:")
	fmt.Println(p)
	//登录校验,解析Token里的参数
	claim, err := jwt.ParseToken(p.Token)
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
	//更新用户基本信息
	code := logic.UpdateUserBase(ctx, p, claim)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	//3.返回成功响应
	io.ResponseSuccess(ctx, common.CodeSuccess)
}

// 修改密码
func UserUpdatePassword(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.ParamUpdatepwd)
	//绑定Query参数
	if err := ctx.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	fmt.Println("请求参数:")
	fmt.Println(p)
	//登录校验,解析Token里的参数
	claim, err := jwt.ParseToken(p.Token)
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
	//更新用户基本信息
	code := logic.UpdateUserPassword(ctx, p, claim)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	//3.返回成功响应
	io.ResponseSuccess(ctx, common.CodeSuccess)
}

// 找回密码
func PasswordForget(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.ParamForgetpwd)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	fmt.Println("请求参数:")
	fmt.Println(p)
	//未来加上短信验证

	//2.服务调用
	ret, code := logic.QueryPassword(ctx, p)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	resp := io.PasswordResp{
		Response: io.Response{StatusCode: 0, StatusMsg: "success"},
		Password: ret,
	}
	//3.返回成功响应
	io.ResponseSuccessPassword(ctx, &resp)
}

// 用户注销
func UserDelete(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.ParamUserDelete)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	fmt.Println("请求参数:")
	fmt.Println(p)
	//登录校验,解析Token里的参数
	claim, err := jwt.ParseToken(p.Token)
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
	code := logic.DeleteUser(ctx, claim.UserID, p.Password)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	//3.返回成功响应
	io.ResponseSuccess(ctx, common.CodeUserDeleteSuccess)
}

// 更新token
func UpdateToken(ctx *gin.Context) {
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
	token, err = jwt.GenToken(claim.UserID, claim.UserName)
	if err != nil {
		fmt.Println("token生成失败")
		io.ResponseError(ctx, common.CodeNeedLogin)
		return
	}
	//3.返回成功响应
	io.ResponseSuccessLogin(ctx, token)
}

// 用户作品
func UserWorks(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.UserWorkReq)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	//2.服务调用
	//获取用户发布的所有视频
	vids, code := logic.GetUserVideoIDs(ctx, p.UserID)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	var videoInfos []io.VideoInfo
	for _, vid := range vids {
		videoInfo, code := logic.GetVideoInfoByVID(ctx, vid)
		if code != common.CodeSuccess {
			io.ResponseError(ctx, code)
			return
		}
		videoInfos = append(videoInfos, *videoInfo)
	}
	resp := &io.UserWorkResp{
		Response:   io.Response{StatusCode: 0, StatusMsg: "success"},
		VideoInfos: videoInfos,
	}
	//3.返回成功响应
	io.ResponseSuccessUserWork(ctx, resp)
}

// 用户历史记录
func UserHistory(ctx *gin.Context) {
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
	vids, code := logic.GetUserHistoryVideoIDs(ctx, claim)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	var videoInfos []io.VideoInfo
	for _, vid := range vids {
		videoInfo, code := logic.GetVideoInfoByVID(ctx, vid)
		if code != common.CodeSuccess {
			io.ResponseError(ctx, code)
			return
		}
		videoInfos = append(videoInfos, *videoInfo)
	}
	resp := &io.UserHistoryResp{
		Response:   io.Response{StatusCode: 0, StatusMsg: "success"},
		VideoInfos: videoInfos,
	}
	//3.返回成功响应
	io.ResponseSuccessUserHistory(ctx, resp)
}

// 用户喜爱视频列表
func UserFavorite(ctx *gin.Context) {
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
	vids, code := logic.GetUserFavoriteVideoIDs(ctx, claim)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	var videoInfos []io.VideoInfo
	for _, vid := range vids {
		videoInfo, code := logic.GetVideoInfoByVID(ctx, vid)
		if code != common.CodeSuccess {
			io.ResponseError(ctx, code)
			return
		}
		videoInfos = append(videoInfos, *videoInfo)
	}
	resp := &io.UserFavoriteResp{
		Response:   io.Response{StatusCode: 0, StatusMsg: "success"},
		VideoInfos: videoInfos,
	}
	//3.返回成功响应
	io.ResponseSuccessUserFavorite(ctx, resp)
}

func CareUser(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(io.CareUserReq)
	//绑定Query参数
	if err := ctx.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		io.ResponseError(ctx, common.CodeInvalidParam)
		return
	}
	fmt.Println("请求参数:")
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
	if code := logic.CareUser(ctx, p, claim); code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	//3.返回响应
	io.ResponseSuccess(ctx, common.CodeSuccess)
}

// 获取关注列表
func CareList(ctx *gin.Context) {
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
	}
	//2.服务调用
	careList, code := logic.GetUserCareList(ctx, claim)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	resp := &io.CareListResp{
		Response: io.Response{StatusCode: 0, StatusMsg: "success"},
		CareList: careList,
	}
	//3.返回响应
	io.ResponseSuccessCareList(ctx, resp)
}

func FansList(ctx *gin.Context) {
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
	}
	//2.服务调用
	fansList, code := logic.GetUserFansList(ctx, claim)
	if code != common.CodeSuccess {
		io.ResponseError(ctx, code)
		return
	}
	resp := &io.FansListResp{
		Response: io.Response{StatusCode: 0, StatusMsg: "success"},
		FansList: fansList,
	}
	//3.返回响应
	io.ResponseSuccessFansList(ctx, resp)
}
