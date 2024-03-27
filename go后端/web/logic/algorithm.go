package logic

import (
	"WebVideoServer/common"
	"WebVideoServer/jwt"
	"WebVideoServer/web/model/mysql"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

const MAX_VedioLink = 10 //最大缓存视频链接

type TopVideo struct {
	VID   int64
	Value float64
	Cnt   int64
}

type RefereeVideo struct {
	VID   int64
	Value float64
	Cnt   int64
}

type CareVideo struct {
	VID        int64
	PublicTime int64
}

// 热点视频排序算法:按照播放次数从小到大,再按权值从大到小
func SortTopVID(v []TopVideo) {
	sort.Slice(v, func(i, j int) bool {
		if v[i].Cnt != v[j].Cnt {
			return v[i].Cnt < v[j].Cnt
		}
		return v[i].Value > v[j].Value
	})
}

// 推荐视频排序算法:按照播放次数从小到大,再按新权值从大到小
func SortRefereeVID(v []RefereeVideo) {
	sort.Slice(v, func(i, j int) bool {
		if v[i].Cnt != v[j].Cnt {
			return v[i].Cnt < v[j].Cnt
		}
		return v[i].Value > v[j].Value
	})
}

// 关注视频排序算法:直接按照发布时间距当前时间从小到大排序
func SortCareVID(v []CareVideo) {
	sort.Slice(v, func(i, j int) bool {
		return v[i].PublicTime < v[j].PublicTime
	})
}

// 时间占权重比例
func TimeToWeight(Vediotime int64) float64 {
	if Vediotime < 3*60*60*1000 {
		return 10000
	}
	if Vediotime < 6*60*60*1000 {
		return 5000
	}
	if Vediotime < 12*60*60*1000 {
		return 1000
	}
	if Vediotime < 24*60*60*1000 {
		return 500
	}
	if Vediotime < 2*24*60*60*1000 {
		return 100
	}
	if Vediotime < 5*24*60*60*1000 {
		return 10
	}
	if Vediotime < 10*24*60*60*1000 {
		return 1
	}
	if Vediotime < 20*24*60*60*1000 {
		return 0.5
	}
	if Vediotime < 30*24*60*60*1000 {
		return 0.1
	}
	return 0.01
}

// 热点推送算法
func GetTopVideoIDs(ctx *gin.Context, claim *jwt.MyClaims) ([]int64, common.ResCode) {
	//获取所有视频VID
	AllVID, err := mysql.QueryAllVID(ctx)
	if err != nil {
		return nil, common.CodeMysqlFailed
	}
	//计算所有视频权值并排序
	var vv []TopVideo
	for _, val := range AllVID {
		var flag TopVideo

		weight, _ := mysql.QueryWeightByVID(ctx, val)
		stime, _ := mysql.QueryStartTimeByVID(ctx, val)
		cnt, _ := mysql.QueryUserWatchCount(ctx, claim.UserID, val)

		flag.VID = val
		flag.Cnt = cnt
		flag.Value = weight * TimeToWeight(GetNowTime()-stime)
		vv = append(vv, flag)
	}
	SortTopVID(vv)
	//找出权值最大的最大缓存量个视频
	var TopIDs []int64
	for i := 0; i < MAX_VedioLink; i++ {
		if len(vv) > i {
			TopIDs = append(TopIDs, vv[i].VID)
		}
	}
	return TopIDs, common.CodeSuccess
}

func GetRefereeVideoIDs(ctx *gin.Context, claim *jwt.MyClaims) ([]int64, common.ResCode) {
	//获取所有视频VID
	AllVID, err := mysql.QueryAllVID(ctx)
	if err != nil {
		return nil, common.CodeMysqlFailed
	}
	//计算所有视频权值并排序
	var vv []RefereeVideo
	for _, val := range AllVID {
		var flag RefereeVideo

		weight, _ := mysql.QueryWeightByVID(ctx, val)
		stime, _ := mysql.QueryStartTimeByVID(ctx, val)
		cnt, _ := mysql.QueryUserWatchCount(ctx, claim.UserID, val)

		tagsWeight, code := GetWeightByUserLookVideo(ctx, claim.UserID, val)
		if code != common.CodeSuccess {
			return nil, common.CodeGetVideoTagsWeightError
		}

		flag.VID = val
		flag.Cnt = cnt
		flag.Value = (weight + float64(tagsWeight)) * TimeToWeight(GetNowTime()-stime)

		vv = append(vv, flag)
	}
	SortRefereeVID(vv)
	//找出权值最大的最大缓存量个视频
	var RefereeIDs []int64
	for i := 0; i < MAX_VedioLink; i++ {
		if len(vv) > i {
			RefereeIDs = append(RefereeIDs, vv[i].VID)
		}
	}
	return RefereeIDs, common.CodeSuccess

}

func GetCareVideoIDs(ctx *gin.Context, claim *jwt.MyClaims) ([]int64, common.ResCode) {
	//获取所有关注人uid
	AllCareUID, err := mysql.QueryUserCareList(ctx, claim.UserID)
	if err != nil {
		return nil, common.CodeMysqlFailed
	}
	//查询所有关注用户的所有视频vid
	var careVIDs []int64
	for _, val := range AllCareUID {
		vids, _ := mysql.QueryVideoIDByUserID(ctx, val)
		careVIDs = append(careVIDs, vids...)
	}
	//获取视频VID并排序
	var vv []CareVideo
	for _, val := range careVIDs {
		var flag CareVideo

		stime, _ := mysql.QueryStartTimeByVID(ctx, val)

		flag.VID = val
		flag.PublicTime = GetNowTime() - stime
		vv = append(vv, flag)
	}
	SortCareVID(vv)
	//找出权值最大的最大缓存量个视频
	var CareIDs []int64
	for i := 0; i < MAX_VedioLink; i++ {
		if len(vv) > i {
			CareIDs = append(CareIDs, vv[i].VID)
		}
	}
	return CareIDs, common.CodeSuccess
}

// 获取当前时间戳(ms)
func GetNowTime() int64 {
	return time.Now().UnixNano()
}
