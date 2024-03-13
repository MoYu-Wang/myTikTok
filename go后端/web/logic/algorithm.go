package logic

import (
	"WebVideoServer/web/model/mysql"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

const MAX_VedioLink = 10 //最大缓存视频链接

type Vvalue struct {
	VID   int64
	Value float64
}

func SortByValue(v []Vvalue) {
	sort.Slice(v, func(i, j int) bool {
		return v[i].Value > v[j].Value
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
func TopAlgorithm(ctx *gin.Context) ([]string, error) {
	//获取所有视频VID
	AllVID, err := mysql.QueryAllVID(ctx)
	if err != nil {
		return nil, err
	}
	//判断该用户是否观看过该视频(暂时不写)

	//计算所有视频权值
	var vv []Vvalue
	for _, val := range AllVID {
		var flag Vvalue
		flag.VID = val
		weight, _ := mysql.QueryWeightByVID(ctx, val)
		stime, _ := mysql.QueryStartTimeByVID(ctx, val)
		flag.Value = weight * TimeToWeight(GetNowTime()-stime)
		vv = append(vv, flag)
	}
	//排序找出权值最大的最大缓存量个视频
	SortByValue(vv)
	var TopArr []string
	for i := 0; i < MAX_VedioLink; i++ {
		if len(vv) > i {
			var str string
			str, err = mysql.QueryVlinkByVID(ctx, vv[i].VID)
			TopArr = append(TopArr, str)
			if err != nil {
				return TopArr, err
			}
		}
	}
	return TopArr, err
}

func CareAlgorithm(ctx *gin.Context) ([]string, error) {
	var err error
	var CareArr []string
	return CareArr, err
}

func RefereeAlgorithm(ctx *gin.Context) ([]string, error) {
	var err error
	var RefereeArr []string
	return RefereeArr, err
}

// 获取当前时间戳(ms)
func GetNowTime() int64 {
	return time.Now().UnixNano()
}
