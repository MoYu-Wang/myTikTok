package logic

import (
	"WebVideoServer/common"
	"WebVideoServer/dao"
	"WebVideoServer/io"
	"WebVideoServer/jwt"
	"WebVideoServer/web/model/mysql"
	"sort"

	"github.com/gin-gonic/gin"
)

// 根据视频id获取视频信息
func GetVideoInfoByVID(ctx *gin.Context, videoID int64, claim *jwt.MyClaims) (*io.VideoInfo, common.ResCode) {
	ret, err := mysql.QueryVideoInfoByVID(ctx, videoID)
	if err != nil {
		return nil, common.CodeMysqlFailed
	}
	vfnum, err := mysql.QueryVideoFavoriteNum(ctx, videoID)
	if err != nil {
		return nil, common.CodeMysqlFailed
	}
	vcnum, err := mysql.QueryVideoCommentNum(ctx, videoID)
	if err != nil {
		return nil, common.CodeMysqlFailed
	}
	isf, err := mysql.QueryUserIsLikeVedio(ctx, claim.UserID, videoID)
	if err != nil {
		return nil, common.CodeMysqlFailed
	}
	videoInfo := &io.VideoInfo{
		VideoID:          ret.VideoID,
		UserID:           ret.UserID,
		VideoLink:        ret.VideoLink,
		VideoFavoriteNum: vfnum,
		VideoCommitNum:   vcnum,
		IsFavorite:       isf,
	}
	return videoInfo, common.CodeSuccess
}

// 获取用户发布的所有视频id
func GetUserVideoIDs(ctx *gin.Context, claim *jwt.MyClaims) ([]int64, common.ResCode) {
	ret, err := mysql.QueryVideoIDByUserID(ctx, claim.UserID)
	if err != nil {
		return nil, common.CodeMysqlFailed
	}
	return ret, common.CodeSuccess
}

// 排序历史记录(根据上次观看时间排序)
func SortHistory(v []dao.History) {
	sort.Slice(v, func(i, j int) bool {
		return v[i].LastTime < v[j].LastTime
	})
}

// 获取用户历史记录id
func GetUserHistoryVideoIDs(ctx *gin.Context, claim *jwt.MyClaims) ([]int64, common.ResCode) {
	ret, err := mysql.QueryUserHistory(ctx, claim.UserID)
	if err != nil {
		return nil, common.CodeMysqlFailed
	}
	SortHistory(ret)
	var IDs []int64
	for _, his := range ret {
		IDs = append(IDs, his.VideoID)
	}
	return IDs, common.CodeSuccess
}

// 获取用户点赞视频id
func GetUserFavoriteVideoIDs(ctx *gin.Context, claim *jwt.MyClaims) ([]int64, common.ResCode) {
	ret, err := mysql.QueryUserFavoriteVIDs(ctx, claim.UserID)
	if err != nil {
		return nil, common.CodeMysqlFailed
	}
	return ret, common.CodeSuccess
}
