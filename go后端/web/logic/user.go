package logic

import (
	"WebVideoServer/dao"
	"WebVideoServer/io"
	"WebVideoServer/snowflake"
	"WebVideoServer/web/model"

	"github.com/gin-gonic/gin"
)

// 用户注册
func UserRegister(ctx *gin.Context, p *io.ParamRegister) error {
	//判断用户是否存在
	ret, _ := model.UserIsExist(ctx, p)
	if ret {
		return io.ErrorUserNameIsExist
	}
	//生成UID
	userID := snowflake.GenID()
	//构建一个user实例
	user := &dao.User{
		UserID:   userID,
		Username: p.Username,
		PassWord: p.Password,
	}
	//保存进数据库
	return model.InsertUser(ctx, user)
}

// 用户登录
func UserLogin(ctx *gin.Context) {
	//判断用户ID是否存在

	//判断用户密码是否正确

}
