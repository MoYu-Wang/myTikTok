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
func UserRegister(ctx *gin.Context, p *io.ParamRegister) (int64, error) {
	//判断用户密码是否合法
	if len(p.PassWord) < 5 || len(p.PassWord) > 18 {
		return 0, common.ErrorPasswordNotOK
	}
	//生成UID
	userID := snowflake.GenID()
	//构建一个user实例
	user := &dao.User{
		UserID:   userID,
		UserName: p.UserName,
		PassWord: p.PassWord,
	}
	//保存进数据库
	return user.UserID, mysql.InsertUser(ctx, user)
}

// 用户登录
func UserLogin(ctx *gin.Context, p *io.ParamLogin) (token string, err error) {
	//判断用户ID是否存在
	f, err := mysql.QueryUserIsExist(ctx, p.UserID)
	if err != nil {
		return "", err
	}
	if !f {
		return "", common.ErrorUserNotLogin
	}
	//判断用户密码是否正确
	pwd, err := mysql.QueryUserPassword(ctx, p.UserID)
	if err != nil {
		return "", err
	}
	if p.PassWord != pwd {
		return "", common.ErrorPassword
	}
	//生成Token
	userName, err := mysql.QueryUserName(ctx, p.UserID)
	if err != nil {
		return "", err
	}
	token, err = jwt.GenToken(p.UserID, userName)
	if err != nil {
		return
	}
	return token, err
}

func GetUserInfo(ctx *gin.Context, p *io.UserInfoReq, claim *jwt.MyClaims) (*io.User, error) {
	userResp := new(io.User)
	userResp.ID = p.UserID
	//查询用户信息
	user, err := mysql.QueryUserInfo(ctx, p.UserID)
	if err != nil {
		return nil, err
	}
	userResp.Name = user.UserName
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

func UserUpdatePassword(ctx *gin.Context, p *io.ParamUpdate) error {
	//修改密码
	if err := mysql.UpdateUserPassword(ctx, p); err != nil {
		return err
	}
	return nil
}
