package mysql

import (
	"WebVideoServer/dao"
	"context"
)

// 查询用户的点赞视频vid列表
func QueryUserFavoriteVIDs(ctx context.Context, userID int64) ([]int64, error) {
	db := GetDB(ctx)
	var ret []int64
	err := db.Table("Favorite").Select("VideoID").Where("UserID=?", userID).Find(&ret).Error
	return ret, err
}

// 查询用户点赞总数
func QueryUserFavoriteNum(ctx context.Context, userID int64) (int64, error) {
	db := GetDB(ctx)
	var ret int64
	err := db.Table("Favorite").Where("UserID=?", userID).Count(&ret).Error
	return ret, err
}

// 查询视频喜爱人数
func QueryVideoFavoriteNum(ctx context.Context, videoID int64) (int64, error) {
	db := GetDB(ctx)
	var ret int64
	err := db.Table("Favorite").Where("VideoID=?", videoID).Count(&ret).Error
	return ret, err
}

// 查询该用户是否点赞该视频
func QueryUserIsLikeVedio(ctx context.Context, userID int64, videoID int64) (bool, error) {
	db := GetDB(ctx)
	var ret int64
	err := db.Table("Favorite").Where("UserID=? AND VideoID=?", userID, videoID).Count(&ret).Error
	if err != nil {
		return false, err
	}
	if ret > 0 {
		return true, err
	}
	return false, err
}

// 用户点赞
func InsertUserLikeVedio(ctx context.Context, ret dao.Favorite) error {
	db := GetDB(ctx)
	return db.Table("Favorite").Create(&ret).Error
}

// 取消点赞
func DeleteUserLikeVedio(ctx context.Context, ret dao.Favorite) error {
	db := GetDB(ctx)
	favorite := dao.Favorite{}
	return db.Table("Favorite").Delete(&favorite, ret).Error
}

// 删除用户所有点赞
func DeleteUserALLFavorite(ctx context.Context, userID int64) error {
	db := GetDB(ctx)
	favorite := dao.Favorite{}
	return db.Table("Favorite").Where("UserID=?", userID).Delete(&favorite).Error
}

// 删除视频所有点赞
func DeleteVideoALLFavorite(ctx context.Context, videoID int64) error {
	db := GetDB(ctx)
	favorite := dao.Favorite{}
	return db.Table("Favorite").Where("VideoID=?", videoID).Delete(&favorite).Error
}
