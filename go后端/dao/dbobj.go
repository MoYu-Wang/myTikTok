package dao

type User struct {
	UID      int64  //用户名uid
	Uname    string //昵称
	PassWord string //密码
}

type Vedio struct {
	VID        int64    //视频vid
	Vname      string   //视频名称
	Vlink      string   //视频链接
	UID        int64    //发布人uid
	Tag        []string //视频标签数组
	Start_Time int64    //视频开始发布时间
}

type VedioTime struct {
	VID       int64 //视频vid
	Play_Time int64 //视频被播放时间
}

type TagTime struct {
	Tag       string //标签
	Play_Time int64  //标签被播放时间
}

type VedioWeight struct {
	VID         string  //视频vid
	Weight      float64 //视频初始权重
	WeightMul   float64 //视频权重倍率
	Public_Time int64   //视频发布时间戳

}
