package mysql

import (
	"WebVideoServer/dao"
	"context"
)

// 根据视频ID查找视频链接
func VIDGetVlink(ctx context.Context, videoID int64) (string, error) {
	db := GetDB(ctx)
	var ret string
	err := db.Table("Video").Select("VideoLink").Where("VideoID=?", videoID).Find(&ret).Error
	return ret, err
}

// 根据视频ID查找视频发布时间
func VIDGetStartTime(ctx context.Context, videoID int64) (int64, error) {
	db := GetDB(ctx)
	var ret int64
	err := db.Table("Video").Select("StartTime").Where("VideoID=?", videoID).Find(&ret).Error
	return ret, err
}

// 根据视频ID查找视频基础权重
func VIDGetWeight(ctx context.Context, videoID int64) (float64, error) {
	db := GetDB(ctx)
	var ret float64
	err := db.Table("Video").Select("Weight").Where("VideoID=?", videoID).Find(&ret).Error
	return ret, err
}

// 根据视频ID查找视频所有信息
func VIDGetVideo(ctx context.Context, videoID int64) (*dao.Video, error) {
	db := GetDB(ctx)
	var ret *dao.Video
	err := db.Table("Video").Where("VideoID=?", videoID).Find(&ret).Error
	return ret, err
}

// 根据视频名称模糊查找类似视频所有VID(暂时这样写,不确定对不对)
func VNameGetVID(ctx context.Context, videoName string) ([]int64, error) {
	db := GetDB(ctx)
	var ret []int64
	err := db.Table("Video").Where("VideoName LIKE ?", "%"+videoName+"%").Find(&ret).Error
	return ret, err
}

// 获取所有视频ID
func GetAllVID(ctx context.Context) ([]int64, error) {
	db := GetDB(ctx)
	var ret []int64
	err := db.Table("Video").Find(&ret).Error
	return ret, err
}
