package mysql

import (
	"WebVideoServer/dao"
	"context"
)

// 查询视频所有评论信息
func QueryVideoComments(ctx context.Context, videoID int64) ([]dao.CommentList, error) {
	db := GetDB(ctx)
	var ret []dao.CommentList
	err := db.Table("CommentList").Where("VideoID=?", videoID).Find(&ret).Error
	return ret, err
}

// 查询视频评论总数
func QueryVideoCommentNum(ctx context.Context, videoID int64) (int64, error) {
	db := GetDB(ctx)
	var ret int64
	err := db.Table("CommentList").Where("VideoID=?", videoID).Count(&ret).Error
	return ret, err
}

// 添加评论
func InsertVideoComment(ctx context.Context, comment dao.CommentList) error {
	db := GetDB(ctx)
	return db.Table("CommentList").Create(&comment).Error
}

// 删除评论(不确定)
func DeleteVideoComment(ctx context.Context, comment dao.CommentList) error {
	db := GetDB(ctx)
	commentList := dao.CommentList{}
	return db.Table("CommentList").Delete(&commentList, comment).Error
}

// 删除用户所有评论
func DeleteUserALLComment(ctx context.Context, userID int64) error {
	db := GetDB(ctx)
	commentList := dao.CommentList{}
	return db.Table("CommentList").Where("UserID=?", userID).Delete(&commentList).Error
}

// 删除视频所有评论
func DeleteVideoALLComment(ctx context.Context, videoID int64) error {
	db := GetDB(ctx)
	commentList := dao.CommentList{}
	return db.Table("CommentList").Where("VideoID=?", videoID).Delete(&commentList).Error
}
