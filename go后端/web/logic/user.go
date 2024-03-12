package logic

import (
	"WebVideoServer/common"
	"WebVideoServer/dao"
	"WebVideoServer/io"
	"WebVideoServer/jwt"
	"WebVideoServer/snowflake"
	"WebVideoServer/web/model/mysql"

	"github.com/gin-gonic/gin"
)

// 用户注册
func UserRegister(ctx *gin.Context, p *io.ParamRegister) common.ResCode {
	//判断用户密码和手机号是否合法(交给前端)
	//判断该手机是否已经被注册
	f, err := mysql.QueryIphoneIDIsExist(ctx, p.IphoneID)
	if err != nil {
		return common.CodeMysqlFailed
	}
	if f {
		return common.CodeIphoneIsExist
	}
	//生成UID 雪花ID%10000000000
	userID := snowflake.GenID() % 10000000000
	//判断该UID是否存在,如果存在,重新生成
	for f, err := mysql.QueryUserIsExist(ctx, userID); f; {
		if err != nil {
			return common.CodeMysqlFailed
		}
		userID = snowflake.GenID() % 10000000000
	}
	//构建一个user实例
	user := &dao.User{
		UserID:   userID,
		UserName: p.UserName,
		PassWord: p.PassWord,
		IphoneID: p.IphoneID,
	}
	//传输UserID
	ctx.Set("UserID", userID)
	//保存进数据库
	err = mysql.InsertUser(ctx, user)
	if err != nil {
		return common.CodeMysqlFailed
	}
	return common.CodeSuccess
}

// 用户id登录
func UserIDLogin(ctx *gin.Context, p *io.ParamLogin) (string, common.ResCode) {
	//判断用户ID是否存在
	f, err := mysql.QueryUserIsExist(ctx, p.UserID)
	if err != nil {
		return "", common.CodeMysqlFailed
	}
	if !f {
		return "", common.CodeInvalidLoginUserID
	}
	//判断用户密码是否正确
	pwd, err := mysql.QueryPasswordByUID(ctx, p.UserID)
	if err != nil {
		return "", common.CodeMysqlFailed
	}
	if p.PassWord != pwd {
		return "", common.CodeInvalidLoginPassword
	}
	//生成Token
	userName, err := mysql.QueryUserName(ctx, p.UserID)
	if err != nil {
		return "", common.CodeMysqlFailed
	}
	var token string
	token, err = jwt.GenToken(p.UserID, userName)
	if err != nil {
		return "", common.CodeInvalidToken
	}
	return token, common.CodeSuccess
}

// 手机号登录
func IphoneIDLogin(ctx *gin.Context, p *io.ParamLogin) (string, int64, common.ResCode) {
	//判断手机号是否存在
	f, err := mysql.QueryIphoneIDIsExist(ctx, p.IphoneID)
	if err != nil {
		return "", 0, common.CodeMysqlFailed
	}
	if !f {
		return "", 0, common.CodeIphoneNotExist
	}
	//查询用户id
	userID, err := mysql.QueryUserIDByIphoneID(ctx, p.IphoneID)
	if err != nil {
		return "", 0, common.CodeMysqlFailed
	}
	//判断用户密码是否正确
	pwd, err := mysql.QueryPasswordByIphoneID(ctx, p.IphoneID)
	if err != nil {
		return "", 0, common.CodeMysqlFailed
	}
	if p.PassWord != pwd {
		return "", 0, common.CodeInvalidLoginPassword
	}
	//生成Token
	userName, err := mysql.QueryUserName(ctx, userID)
	if err != nil {
		return "", 0, common.CodeMysqlFailed
	}
	token, err := jwt.GenToken(userID, userName)
	if err != nil {
		return "", 0, common.CodeInvalidToken
	}
	return token, userID, common.CodeSuccess
}

// 获取用户信息
func GetUserInfo(ctx *gin.Context, p *io.UserInfoReq, claim *jwt.MyClaims) (*io.UserInfo, error) {
	userResp := new(io.UserInfo)
	userResp.ID = p.UserID
	//查询用户昵称
	userName, err := mysql.QueryUserName(ctx, p.UserID)
	if err != nil {
		return nil, err
	}
	userResp.Name = userName
	//获取用户粉丝数
	fansCount, err := mysql.QueryUserFansCount(ctx, p.UserID)
	if err != nil {
		return nil, err
	}
	userResp.FansCount = fansCount
	//获取用户关注数
	careCount, err := mysql.QueryUserCareCount(ctx, p.UserID)
	if err != nil {
		return nil, err
	}
	userResp.CareCount = careCount
	//获取是否关注
	userResp.IsCare, err = mysql.QueryIsCare(ctx, claim.UserID, p.UserID)
	if err != nil {
		return nil, err
	}
	return userResp, nil
}

// 获取用户基本信息
func GetUser(ctx *gin.Context, claim *jwt.MyClaims) (*dao.User, common.ResCode) {
	//查询用户基本信息
	ret, err := mysql.QueryUserByUID(ctx, claim.UserID)
	if err != nil {
		return nil, common.CodeMysqlFailed
	}
	return ret, common.CodeSuccess
}

// 修改用户基本信息
func UpdateUserBase(ctx *gin.Context, p *io.ParamUpdate, claim *jwt.MyClaims) (*dao.User, common.ResCode) {
	//获取当前用户基本信息
	user, err := mysql.QueryUserByUID(ctx, claim.UserID)
	if err != nil {
		return nil, common.CodeMysqlFailed
	}
	//修改昵称
	if p.UserName != "" {
		if err := mysql.UpdateUserName(ctx, claim.UserID, p.UserName); err != nil {
			return nil, common.CodeMysqlFailed
		}
		user.UserName = p.UserName
	}
	//修改密码
	if p.PassWord != "" {
		if err := mysql.UpdateUserPassword(ctx, claim.UserID, p.PassWord); err != nil {
			return nil, common.CodeMysqlFailed
		}
		user.PassWord = p.PassWord
	}
	//修改手机号
	if p.IphoneID != "" {
		if err := mysql.UpdateIphoneID(ctx, claim.UserID, p.IphoneID); err != nil {
			return nil, common.CodeMysqlFailed
		}
		user.IphoneID = p.IphoneID
	}

	return user, common.CodeSuccess
}

// 找回密码
func QueryPassword(ctx *gin.Context, p *io.ParamForgetpwd) (string, common.ResCode) {
	var pwd string
	var err error
	//判断是用户id找回还是手机号找回
	if p.IphoneID != "" {
		if pwd, err = mysql.QueryPasswordByIphoneID(ctx, p.IphoneID); err != nil {
			return "", common.CodeMysqlFailed
		}
	} else {
		if pwd, err = mysql.QueryPasswordByUID(ctx, p.UserID); err != nil {
			return "", common.CodeMysqlFailed
		}
	}
	return pwd, common.CodeSuccess
}
