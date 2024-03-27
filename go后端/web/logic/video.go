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

// 获取用户观看视频标签总权值
func GetWeightByUserLookVideo(ctx *gin.Context, userID int64, videoID int64) (int64, common.ResCode) {
	var weight int64
	weight = 0
	tags, err := mysql.QueryTagArrByVideoID(ctx, videoID)
	if err != nil {
		return 0, common.CodeMysqlFailed
	}
	for _, tag := range tags {
		tagTime, err := mysql.QueryUserLookTagTime(ctx, userID, tag)
		if err != nil {
			return 0, common.CodeMysqlFailed
		}
		weight += tagTime
	}
	return weight, common.CodeSuccess
}

// 对视频进行操作
func OperateVideo(ctx *gin.Context, p *io.OperateVideoReq, claim *jwt.MyClaims) common.ResCode {
	//对视频进行基础权值增加处理
	if err := mysql.AddVideoWeight(ctx, p.VideoID, float64(p.WatchTime)); err != nil {
		return common.CodeMysqlFailed
	}
	//判断是否登录,若未登录,则直接返回
	if _, err := jwt.ParseToken(p.Token); err != nil {
		return common.CodeSuccess
	}
	//增加用户打开视频次数
	if err := mysql.AddUserWatch(ctx, claim.UserID, p.VideoID); err != nil {
		return common.CodeMysqlFailed
	}
	//增加用户观看标签时长
	//获取标签数组
	tags, err := mysql.QueryTagArrByVideoID(ctx, p.VideoID)
	if err != nil {
		return common.CodeMysqlFailed
	}
	for _, tag := range tags {
		if err := mysql.InsertUserLookTagTime(ctx, claim.UserID, tag, p.WatchTime); err != nil {
			return common.CodeMysqlFailed
		}
	}
	//判断是否点赞 0:未进行操作 1:点赞操作 -1:取消点赞操作
	if p.IsFavorite > 0 {
		if err := mysql.InsertUserLikeVedio(ctx, dao.Favorite{
			UserID:  claim.UserID,
			VideoID: p.VideoID,
		}); err != nil {
			return common.CodeMysqlFailed
		}
	}
	if p.IsFavorite < 0 {
		if err := mysql.DeleteUserLikeVedio(ctx, dao.Favorite{
			UserID:  claim.UserID,
			VideoID: p.VideoID,
		}); err != nil {
			return common.CodeMysqlFailed
		}
	}
	//判断是否评论
	if p.CommentNum > 0 {
		for _, commentText := range p.CommentTexts {
			comment := dao.CommentList{
				UserID:      claim.UserID,
				VideoID:     p.VideoID,
				CommentText: commentText,
				CommentTime: GetNowTime(),
			}
			if err := mysql.InsertVideoComment(ctx, comment); err != nil {
				return common.CodeMysqlFailed
			}
		}
	}
	return common.CodeSuccess
}

// 获取模糊搜索视频id列表
func GetSearchVideoIDs(ctx *gin.Context, searchText string) ([]int64, common.ResCode) {
	searchIDs, err := mysql.QueryVIDByVName(ctx, searchText)
	if err != nil {
		return nil, common.CodeMysqlFailed
	}
	return searchIDs, common.CodeSuccess
}
