package mysql

import (
	"WebVideoServer/dao"
	"context"
)

// 查询用户历史记录
func QueryUserHistory(ctx context.Context, UID int64) ([]int64, error) {
	db := GetDB(ctx)
	var ret []int64
	err := db.Table("History").Where("UserID=?", UID).Find(&ret).Error
	return ret, err
}

// 查询用户是否观看过此视频
func QueryUserIsWatch(ctx context.Context, UID int64, VID int64) (bool, error) {
	db := GetDB(ctx)
	var ret int64
	err := db.Table("History").Where("UserID=?,VedioID=?", UID, VID).Count(&ret).Error
	if ret != 0 {
		return true, err
	}
	return false, err
}

// 查询用户打开过该视频多少次
func QueryUserWatchCount(ctx context.Context, UID int64, VID int64) (int64, error) {
	db := GetDB(ctx)
	var ret int64
	err := db.Table("History").Select("Cnt").Where("UserID=?,VedioID=?", UID, VID).Find(&ret).Error
	return ret, err
}

// 增加该用户打开该视频次数
func AddUserWatch(ctx context.Context, UserID int64, VedioID int64) error {
	db := GetDB(ctx)
	f, _ := QueryUserIsWatch(ctx, UserID, VedioID)
	if f {
		ret, _ := QueryUserWatchCount(ctx, UserID, VedioID)
		return db.Table("History").Where("UserID=?,VedioID=?", UserID, VedioID).Update("Cnt", ret+1).Error
	}
	ret := dao.History{
		UserID:  UserID,
		VedioID: VedioID,
		Cnt:     1,
	}
	err := db.Table("History").Create(&ret).Error
	return err
}
