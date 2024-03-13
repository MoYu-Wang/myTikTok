package mysql

import (
	"WebVideoServer/dao"
	"context"
	"time"
)

// 查询用户历史记录
func QueryUserHistory(ctx context.Context, userID int64) ([]dao.History, error) {
	db := GetDB(ctx)
	var ret []dao.History
	err := db.Table("History").Where("UserID=?", userID).Find(&ret).Error
	return ret, err
}

// 查询用户是否观看过此视频
func QueryUserIsWatch(ctx context.Context, userID int64, videoID int64) (bool, error) {
	db := GetDB(ctx)
	var ret int64
	err := db.Table("History").Where("UserID=? AND VideoID=?", userID, videoID).Count(&ret).Error
	if ret != 0 {
		return true, err
	}
	return false, err
}

// 查询用户打开过该视频多少次
func QueryUserWatchCount(ctx context.Context, userID int64, videoID int64) (int64, error) {
	db := GetDB(ctx)
	var ret int64
	err := db.Table("History").Select("Cnt").Where("UserID=? AND VideoID=?", userID, videoID).Find(&ret).Error
	return ret, err
}

// 增加该用户打开视频次数
func AddUserWatch(ctx context.Context, userID int64, videoID int64) error {
	db := GetDB(ctx)
	f, _ := QueryUserIsWatch(ctx, userID, videoID)
	if f {
		ret, _ := QueryUserWatchCount(ctx, userID, videoID)
		return db.Table("History").Where("UserID=? AND VideoID=?", userID, videoID).Update("Cnt", ret+1).Error
	}
	ret := dao.History{
		UserID:   userID,
		VideoID:  videoID,
		Cnt:      1,
		LastTime: time.Now().UnixNano(),
	}
	err := db.Table("History").Create(&ret).Error
	return err
}

// 删除用户所有历史记录
func DeleteUserALLHistory(ctx context.Context, userID int64) error {
	db := GetDB(ctx)
	history := dao.History{}
	return db.Table("History").Where("UserID=?", userID).Delete(&history).Error
}

// 删除观看过该视频的所有用户记录
func DeleteVideoALLHistory(ctx context.Context, videoID int64) error {
	db := GetDB(ctx)
	history := dao.History{}
	return db.Table("History").Where("VideoID=?", videoID).Delete(&history).Error
}
