package mysql

import (
	"WebVideoServer/dao"
	"context"
)

// 删除用户所观看的所有标签
func DeleteUserLookALLTag(ctx context.Context, userID int64) error {
	db := GetDB(ctx)
	tag := dao.UserLookTag{}
	return db.Table("UserLookTag").Where("UserID=?", userID).Delete(&tag).Error
}

// 添加用户所观看的标签时长
func InsertUserLookTagTime(ctx context.Context, userID int64, tag string, addTime int64) error {
	db := GetDB(ctx)
	f, _ := QueryUserIsLookTag(ctx, userID, tag)
	if !f {
		ult := &dao.UserLookTag{
			UserID:   userID,
			Tag:      tag,
			PlayTime: addTime,
		}
		return db.Table("UserLookTag").Create(&ult).Error
	}
	playTime, _ := QueryUserLookTagTime(ctx, userID, tag)
	return db.Table("UserLookTag").Where("UserID=? AND Tag=?", userID, tag).Update("PlayTime", playTime+addTime).Error
}

// 判断该用户是否观看该标签
func QueryUserIsLookTag(ctx context.Context, userID int64, tag string) (bool, error) {
	db := GetDB(ctx)
	var cnt int64
	err := db.Table("UserLookTag").Where("UserID=? AND Tag=?", userID, tag).Count(&cnt).Error
	if cnt > 0 {
		return true, err
	}
	return false, err
}

// 查询用户观看该标签时长
func QueryUserLookTagTime(ctx context.Context, userID int64, tag string) (int64, error) {
	db := GetDB(ctx)
	var ret int64
	err := db.Table("UserLookTag").Select("PlayTime").Where("UserID=?", userID).Find(&ret).Error
	return ret, err
}
