package dao

type User struct {
	UID      string //用户名uid
	Uname    string //昵称
	PassWord string //密码

}

type Vedio struct {
	VID        string   //视频vid
	Vname      string   //视频名称
	Vlink      string   //视频链接
	UID        string   //发布人uid
	Tag        []string //视频标签数组
	Start_Time int64    //视频开始发布时间

}

type VedioWeight struct {
	VID         string  //视频vid
	Weight      float64 //视频初始权重
	WeightMul   float64 //视频权重倍率
	Public_Time int64   //视频发布时间戳

}
