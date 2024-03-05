package mysql

import (
	"WebVideoServer/dao"
	"WebVideoServer/io"
	"context"
)

// 添加用户
func InsertUser(ctx context.Context, user *dao.User) error {
	db := GetDB(ctx)

	return db.Table("User").Create(&user).Error
}

// 判断用户是否存在
func UserIsExist(ctx context.Context, user *io.ParamRegister) (bool, error) {
	db := GetDB(ctx)
	var ret int64
	err := db.Table("User").Where("Username=?", user.Username).Count(&ret).Error
	if err != nil {
		return false, err
	}
	if ret != 0 {
		return true, err
	}
	return false, err
}

// 更改用户密码
func UpdateUserPassword(ctx context.Context, user *dao.User) error {
	db := GetDB(ctx)

	return db.Table("User").Where("UserID=?", user.UserID).Update("PassWord", user.PassWord).Error
}
