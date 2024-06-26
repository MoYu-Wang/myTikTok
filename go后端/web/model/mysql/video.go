package mysql

import (
	"WebVideoServer/dao"
	"context"
	"strings"
)

// 根据视频ID查找视频链接
func QueryVlinkByVID(ctx context.Context, videoID int64) (string, error) {
	db := GetDB(ctx)
	var ret string
	err := db.Table("Video").Select("VideoLink").Where("VideoID=?", videoID).Find(&ret).Error
	return ret, err
}

// 根据视频ID查找视频发布时间
func QueryStartTimeByVID(ctx context.Context, videoID int64) (int64, error) {
	db := GetDB(ctx)
	var ret int64
	err := db.Table("Video").Select("PublicTime").Where("VideoID=?", videoID).Find(&ret).Error
	return ret, err
}

// 根据视频ID查找视频基础权重
func QueryWeightByVID(ctx context.Context, videoID int64) (float64, error) {
	db := GetDB(ctx)
	var ret float64
	err := db.Table("Video").Select("Weight").Where("VideoID=?", videoID).Find(&ret).Error
	return ret, err
}

// 增加用户观看视频时间
func AddVideoWeight(ctx context.Context, videoID int64, watchTime float64) error {
	db := GetDB(ctx)
	weight, _ := QueryWeightByVID(ctx, videoID)
	return db.Table("Video").Where("VideoID=?", videoID).Update("Weight", weight+watchTime).Error
}

// 根据视频ID查找视频所有信息
func QueryVideoInfoByVID(ctx context.Context, videoID int64) (*dao.Video, error) {
	db := GetDB(ctx)
	var ret *dao.Video
	err := db.Table("Video").Where("VideoID=?", videoID).Find(&ret).Error
	return ret, err
}

// 根据视频名称模糊查找类似视频所有VID(暂时这样写,不确定对不对)
func QueryVIDByName(ctx context.Context, Name string) ([]int64, error) {
	db := GetDB(ctx)

	var vIDs1 []int64
	err := db.Table("Video").Select("VideoID").Where("VideoName LIKE ?", "%"+Name+"%").Find(&vIDs1).Error
	if err != nil {
		return nil, err
	}
	var uIDs []int64
	err = db.Table("User").Select("UserID").Where("UserName LIKE ?", "%"+Name+"%").Find(&uIDs).Error
	if err != nil {
		return nil, err
	}
	var vIDs2 []int64
	for _, uID := range uIDs {
		vids, _ := QueryVideoIDByUserID(ctx, uID)
		vIDs2 = append(vIDs2, vids...)
	}
	var vIDs3 []int64
	err = db.Table("Video").Select("VideoId").Where("Tags LIKE ?", "%"+Name+"%").Find(&vIDs3).Error
	if err != nil {
		return nil, err
	}
	ret := append(append(vIDs1, vIDs2...), vIDs3...)
	return ret, err
}

// 获取所有视频ID
func QueryAllVID(ctx context.Context) ([]int64, error) {
	db := GetDB(ctx)
	var ret []int64
	err := db.Table("Video").Select("VideoID").Find(&ret).Error
	return ret, err
}

// 根据用户id查找所发布的所有视频id
func QueryVideoIDByUserID(ctx context.Context, userID int64) ([]int64, error) {
	db := GetDB(ctx)
	var ret []int64
	err := db.Table("Video").Select("VideoID").Where("UserID=?", userID).Find(&ret).Error
	return ret, err
}

// 根据视频id删除视频
func DeleteVideoByVID(ctx context.Context, videoID int64) error {
	db := GetDB(ctx)
	video := dao.Video{}
	return db.Table("Video").Where("VideoID=?", videoID).Delete(&video).Error
}

// 上传视频
func InsertVideo(ctx context.Context, video *dao.Video) error {
	db := GetDB(ctx)
	return db.Table("Video").Create(&video).Error
}

// 获取视频标签
func QueryTagsByVideoID(ctx context.Context, videoID int64) (string, error) {
	db := GetDB(ctx)
	var tags string
	err := db.Table("Video").Select("Tags").Where("VideoID=?", videoID).Find(&tags).Error
	return tags, err
}

// 获取视频标签数组
func QueryTagArrByVideoID(ctx context.Context, videoID int64) ([]string, error) {
	tags, err := QueryTagsByVideoID(ctx, videoID)
	if err != nil {
		return nil, err
	}
	tagArr := strings.Split(tags, "#")
	tagArrs := make([]string, 0, len(tagArr))
	for _, tag := range tagArr {
		if tag != "" {
			tagArrs = append(tagArrs, tag)
		}
	}
	return tagArrs, nil
}

// 根据视频id查找视频发布人
func QueryPublicUserIDByVideoID(ctx context.Context, videoID int64) (int64, error) {
	db := GetDB(ctx)
	var ret int64
	err := db.Table("Video").Select("UserID").Where("VideoID=?", videoID).Find(&ret).Error
	return ret, err
}
