package mysql

import (
	"WebVideoServer/dao"
	"context"
)

// 查询用户粉丝列表
func QueryUserFansList(ctx context.Context, careUserID int64) ([]int64, error) {
	db := GetDB(ctx)
	var ret []int64
	err := db.Table("CareList").Select("UserID").Where("CareUserID=?", careUserID).Find(&ret).Error
	return ret, err
}

// 查询用户粉丝数(用户被多少人关注)
func QueryUserFansCount(ctx context.Context, careUserID int64) (int64, error) {
	db := GetDB(ctx)
	var ret int64
	err := db.Table("CareList").Where("CareUserID=?", careUserID).Count(&ret).Error
	return ret, err
}

// 查询用户关注列表
func QueryUserCareList(ctx context.Context, UserID int64) ([]int64, error) {
	db := GetDB(ctx)
	var ret []int64
	err := db.Table("CareList").Select("CareUserID").Where("UserID=?", UserID).Find(&ret).Error
	return ret, err
}

// 查询用户关注数(用户关注了多少人)
func QueryUserCareCount(ctx context.Context, userID int64) (int64, error) {
	db := GetDB(ctx)
	var ret int64
	err := db.Table("CareList").Where("UserID=?", userID).Count(&ret).Error
	return ret, err
}

// 查询是否关注
func QueryIsCare(ctx context.Context, userID int64, careUserID int64) (bool, error) {
	db := GetDB(ctx)
	var ret int64
	err := db.Table("CareList").Where("UserID=? AND CareUserID=?", userID, careUserID).Count(&ret).Error
	if err != nil {
		return false, err
	}
	if ret > 0 {
		return true, err
	}
	return false, err
}

// 删除用户所有关注
func DeleteUserALLCare(ctx context.Context, userID int64) error {
	db := GetDB(ctx)
	carelist := dao.CareList{}
	return db.Table("CareList").Where("UserID=?", userID).Delete(&carelist).Error
}

// 删除用户所有粉丝
func DeleteUserALLFans(ctx context.Context, careUserID int64) error {
	db := GetDB(ctx)
	carelist := dao.CareList{}
	return db.Table("CareList").Where("CareUserID=?", careUserID).Delete(&carelist).Error
}
