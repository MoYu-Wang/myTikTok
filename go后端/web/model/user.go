package model

import (
	"WebVideoServer/dao"
	"WebVideoServer/io"
	"WebVideoServer/web/model/mysql"
	"context"
)

// 注册用户
func InsertUser(ctx context.Context, user *dao.User) error {
	//mysql 操作
	err := mysql.InsertUser(ctx, user)
	if err != nil {
		return err
	}
	//redis 操作

	return err
}

// 用户是否存在
func UserIsExist(ctx context.Context, user *io.ParamRegister) (bool, error) {
	ret, err := mysql.UserIsExist(ctx, user)
	if err != nil {
		return false, err
	}
	if ret {
		return true, err
	}
	return false, err

}

// 修改密码
func UpdateUserPassword(ctx context.Context, user *dao.User) error {
	err := mysql.UpdateUserPassword(ctx, user)
	if err != nil {
		return err
	}
	return err
}
