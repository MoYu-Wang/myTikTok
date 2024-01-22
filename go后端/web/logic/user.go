package logic

import (
	"WebVideoServer/dao"
	"WebVideoServer/io"
	"WebVideoServer/snowflake"

	"github.com/gin-gonic/gin"
)

//用户注册
func UserRegister(ctx *gin.Context, p *io.ParamRegister) error {

	//判断用户是否存在
	//生成UID
	userID := snowflake.GenID()
	//构建一个user实例
	user := &dao.User{
		UID:      userID,
		Uname:    p.Username,
		PassWord: p.Password,
	}
	//保存进数据库
	// return model.InsertUser(ctx, user)//这里有问题

	user.UID = 1
	return nil
}

//用户登录
func UserLogin(ctx *gin.Context) {

}
