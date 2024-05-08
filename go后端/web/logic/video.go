package logic

import (
	"WebVideoServer/common"
	"WebVideoServer/dao"
	"WebVideoServer/io"
	"WebVideoServer/jwt"
	"WebVideoServer/snowflake"
	"WebVideoServer/web/model/mysql"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 根据视频id获取视频信息
func GetVideoInfoByVID(ctx *gin.Context, videoID int64) (*io.VideoInfo, common.ResCode) {
	ret, err := mysql.QueryVideoInfoByVID(ctx, videoID)
	if err != nil {
		return nil, common.CodeMysqlFailed
	}

	videoInfo := &io.VideoInfo{
		VideoID:   strconv.FormatInt(ret.VideoID, 10),
		VideoName: ret.VideoName,
		VideoTags: ret.Tags,
		UserID:    ret.UserID,
		VideoLink: ret.VideoLink,
	}
	return videoInfo, common.CodeSuccess
}

// 获取与用户相关的视频信息
func GetVideoOperateInfo(ctx *gin.Context, userID int64, videoID int64) (*io.VideoOperateInfo, common.ResCode) {
	vfnum, err := mysql.QueryVideoFavoriteNum(ctx, videoID)
	if err != nil {
		return nil, common.CodeMysqlFailed
	}
	vcnum, err := mysql.QueryVideoCommentNum(ctx, videoID)
	if err != nil {
		return nil, common.CodeMysqlFailed
	}
	isf, err := mysql.QueryUserIsLikeVedio(ctx, userID, videoID)
	if err != nil {
		return nil, common.CodeMysqlFailed
	}

	videoOperateInfo := &io.VideoOperateInfo{
		VideoFavoriteNum: vfnum,
		VideoCommitNum:   vcnum,
		IsFavorite:       isf,
	}
	return videoOperateInfo, common.CodeSuccess
}

// 获取用户发布的所有视频id
func GetUserVideoIDs(ctx *gin.Context, userID int64) ([]int64, common.ResCode) {
	ret, err := mysql.QueryVideoIDByUserID(ctx, userID)
	if err != nil {
		return nil, common.CodeMysqlFailed
	}
	// 使用切片进行倒序
	reversedSlice := make([]int64, len(ret))
	for i, v := range ret {
		reversedSlice[len(ret)-1-i] = v
	}
	return reversedSlice, common.CodeSuccess
}

// 排序历史记录(根据上次观看时间排序)
func SortHistory(v []dao.History) {
	sort.Slice(v, func(i, j int) bool {
		return v[i].LastTime > v[j].LastTime
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
	// 使用切片进行倒序
	reversedSlice := make([]int64, len(ret))
	for i, v := range ret {
		reversedSlice[len(ret)-1-i] = v
	}
	return reversedSlice, common.CodeSuccess
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
	videoID, err := strconv.ParseInt(p.VideoID, 10, 64)
	if err != nil {
		return common.CodeDataTypeChangeError
	}
	//对视频进行基础权值增加处理
	if err := mysql.AddVideoWeight(ctx, videoID, float64(p.WatchTime)); err != nil {
		return common.CodeMysqlFailed
	}
	//判断是否登录,若未登录,则直接返回
	if _, err := jwt.ParseToken(p.Token); err != nil {
		return common.CodeSuccess
	}
	//增加用户打开视频次数
	if err := mysql.AddUserWatch(ctx, claim.UserID, videoID); err != nil {
		return common.CodeMysqlFailed
	}
	//增加用户观看标签时长
	//获取标签数组
	tags, err := mysql.QueryTagArrByVideoID(ctx, videoID)
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
			VideoID: videoID,
		}); err != nil {
			return common.CodeMysqlFailed
		}
	}
	if p.IsFavorite < 0 {
		if err := mysql.DeleteUserLikeVedio(ctx, dao.Favorite{
			UserID:  claim.UserID,
			VideoID: videoID,
		}); err != nil {
			return common.CodeMysqlFailed
		}
	}
	// //判断是否评论
	// if p.CommentNum > 0 {
	// 	for _, commentText := range p.CommentTexts {
	// 		comment := dao.CommentList{
	// 			CommentID:   snowflake.GenID(),
	// 			UserID:      claim.UserID,
	// 			VideoID:     videoID,
	// 			CommentText: commentText,
	// 			CommentTime: GetNowTime(),
	// 		}
	// 		if err := mysql.InsertVideoComment(ctx, comment); err != nil {
	// 			return common.CodeMysqlFailed
	// 		}
	// 	}
	// }
	return common.CodeSuccess
}

// 获取模糊搜索视频id列表
func GetSearchVideoIDs(ctx *gin.Context, searchText string) ([]int64, common.ResCode) {
	searchIDs, err := mysql.QueryVIDByName(ctx, searchText)
	if err != nil {
		return nil, common.CodeMysqlFailed
	}
	//去重
	encountered := map[int64]bool{}
	result := []int64{}
	for _, v := range searchIDs {
		if !encountered[v] {
			encountered[v] = true
			result = append(result, v)
		}
	}
	return result, common.CodeSuccess
}

// 获取视频评论
func GetVideoComment(ctx *gin.Context, videoID int64) ([]io.VideoComment, common.ResCode) {
	comments, err := mysql.QueryVideoComments(ctx, videoID)
	if err != nil {
		return nil, common.CodeMysqlFailed
	}
	var vcomments []io.VideoComment
	for _, val := range comments {
		//查询评论人昵称
		userName, _ := mysql.QueryUserName(ctx, val.UserID)
		vcomment := &io.VideoComment{
			CommentID:   strconv.FormatInt(val.CommentID, 10),
			UserID:      val.UserID,
			UserName:    userName,
			CommentTime: val.CommentTime,
			CommentText: val.CommentText,
		}
		vcomments = append(vcomments, *vcomment)
	}
	// 使用切片进行倒序
	reversedSlice := make([]io.VideoComment, len(vcomments))
	for i, v := range vcomments {
		reversedSlice[len(vcomments)-1-i] = v
	}
	return reversedSlice, common.CodeSuccess
}

// 视频点赞
func FavoriteVideo(ctx *gin.Context, p *io.FavoriteVideoReq, claim *jwt.MyClaims) common.ResCode {
	//将videoID转化为int64
	videoID, err := strconv.ParseInt(p.VideoID, 10, 64)
	if err != nil {
		return common.CodeDataTypeChangeError
	}
	//判断是否点赞 0:未进行操作 1:点赞操作 -1:取消点赞操作
	if p.IsFavorite > 0 {
		if err := mysql.InsertUserLikeVedio(ctx, dao.Favorite{
			UserID:  claim.UserID,
			VideoID: videoID,
		}); err != nil {
			return common.CodeMysqlFailed
		}
	}
	if p.IsFavorite < 0 {
		if err := mysql.DeleteUserLikeVedio(ctx, dao.Favorite{
			UserID:  claim.UserID,
			VideoID: videoID,
		}); err != nil {
			return common.CodeMysqlFailed
		}
	}
	return common.CodeSuccess
}

// 评论视频
func CommentVideo(ctx *gin.Context, p *io.CommentVideoReq, claim *jwt.MyClaims) (int64, common.ResCode) {
	//将videoID转化为int64
	videoID, err := strconv.ParseInt(p.VideoID, 10, 64)
	if err != nil {
		return 0, common.CodeDataTypeChangeError
	}

	comment := dao.CommentList{
		CommentID:   snowflake.GenID(),
		UserID:      claim.UserID,
		VideoID:     videoID,
		CommentText: p.CommentText,
		CommentTime: GetNowTime(),
	}
	if err := mysql.InsertVideoComment(ctx, comment); err != nil {
		return 0, common.CodeMysqlFailed
	}
	return comment.CommentID, common.CodeSuccess
}

// 删除评论
func DeleteComment(ctx *gin.Context, p *io.DeleteCommentReq, claim *jwt.MyClaims) common.ResCode {
	commentID, err := strconv.ParseInt(p.CommentID, 10, 64)
	if err != nil {
		return common.CodeDataTypeChangeError
	}
	//判断该评论是否是自己的
	f, err := mysql.QueryCommentFromUser(ctx, commentID, claim.UserID)
	if err != nil {
		return common.CodeMysqlFailed
	}
	if !f {
		return common.CodeCommentNotOwn
	}
	//删除评论
	err = mysql.DeleteVideoComment(ctx, commentID)
	if err != nil {
		return common.CodeMysqlFailed
	}
	return common.CodeSuccess
}

func DeleteVideo(ctx *gin.Context, p *io.DeleteVideoReq, claim *jwt.MyClaims) common.ResCode {
	//将videoID转化为int64
	videoID, err := strconv.ParseInt(p.VideoID, 10, 64)
	if err != nil {
		return common.CodeDataTypeChangeError
	}
	//删除视频操作
	//判断该视频是否是自己的
	uid, err := mysql.QueryPublicUserIDByVideoID(ctx, videoID)
	if err != nil {
		return common.CodeMysqlFailed
	}
	if uid != claim.UserID {
		return common.CodeVideoNotOwn
	}
	//删除视频所有点赞
	err = mysql.DeleteVideoALLFavorite(ctx, videoID)
	if err != nil {
		return common.CodeMysqlFailed
	}
	//删除视频所有评论
	err = mysql.DeleteVideoALLComment(ctx, videoID)
	if err != nil {
		return common.CodeMysqlFailed
	}
	//删除该视频历史记录
	err = mysql.DeleteVideoALLHistory(ctx, videoID)
	if err != nil {
		return common.CodeMysqlFailed
	}
	//删除视频
	err = mysql.DeleteVideoByVID(ctx, videoID)
	if err != nil {
		return common.CodeMysqlFailed
	}
	return common.CodeSuccess
}
