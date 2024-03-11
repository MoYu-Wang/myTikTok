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

// 注销用户
func DeleteUser(ctx context.Context, user *dao.User) error {
	db := GetDB(ctx)
	return db.Table("User").Delete(&dao.User{}, user).Error
}

// 判断用户是否存在
func QueryUserIsExist(ctx context.Context, userID int64) (bool, error) {
	db := GetDB(ctx)
	var ret int64
	err := db.Table("User").Where("UserID=?", userID).Count(&ret).Error
	if err != nil {
		return false, err
	}
	if ret > 0 {
		return true, err
	}
	return false, err
}

// 更改用户密码
func UpdateUserPassword(ctx context.Context, p *io.ParamUpdate) error {
	db := GetDB(ctx)
	return db.Table("User").Where("UserID=?", p.UserID).Update("PassWord", p.PassWord).Error
}

// 根据用户昵称查找用户id
func QueryUserIDByUserName(ctx context.Context, userName string) ([]int64, error) {
	db := GetDB(ctx)
	var ret []int64
	err := db.Table("User").Select("UserID").Where("UserName=?", userName).Find(&ret).Error
	return ret, err
}

// 根据用户id查找密码
func QueryUserPassword(ctx context.Context, userID int64) (string, error) {
	db := GetDB(ctx)
	var ret string
	err := db.Table("User").Select("PassWord").Where("UserID=?", userID).Find(&ret).Error
	return ret, err
}

// 根据用户id查找用户昵称
func QueryUserName(ctx context.Context, userID int64) (string, error) {
	db := GetDB(ctx)
	var ret string
	err := db.Table("User").Select("UserName").Where("UserID=?", userID).Find(&ret).Error
	return ret, err
}

func QueryUserInfo(ctx context.Context, userID int64) (*dao.User, error) {
	db := GetDB(ctx)
	var ret *dao.User
	err := db.Table("User").Where("UserID=?", userID).Find(&ret).Error
	return ret, err
}
