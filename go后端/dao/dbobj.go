package dao

import "gorm.io/gorm"

//用户表
type User struct {
	gorm.Model
	UID      int64  `gorm:"column:UID;not null;comment:用户名uid"`
	Uname    string `gorm:"column:Uname;not null;comment:昵称"`
	PassWord string `gorm:"column:PassWord;not null;comment:密码"`
}

//视频表
type Vedio struct {
	gorm.Model
	VID        int64    `gorm:"column:VID;comment:视频vid"`
	Vname      string   `gorm:"column:Vname;comment:视频名称"`
	Vlink      string   `gorm:"column:Vlink;comment:视频链接"`
	UID        int64    `gorm:"column:UID;comment:发布人uid"`
	Tag        []string `gorm:"column:Tag;comment:视频标签数组"`
	Weight     float64  `gorm:"column:Weight;comment:视频初始权重"`
	PublicTime int64    `gorm:"column:PublicTime;comment:视频开始发布时间"`
}

//关注列表
type CareList struct {
	gorm.Model
	UID  int64 `gorm:"column:UID;comment:用户uid"`
	CUID int64 `gorm:"column:CUID;comment:被关注用户uid"`
}

//用户观看标签表
type UserLookTag struct {
	gorm.Model
	UID      int64  `gorm:"column:UID;comment:用户uid"`
	Tag      string `gorm:"column:Tag;comment:用户观看标签"`
	PlayTime int    `gorm:"column:PlayTime;comment:标签被播放时间(单位:时间戳)"`
}

// //视频热度表
// type Top struct {
// 	VID        int64   `gorm:"column:VID;comment:视频vid"`
// 	PublicTime int64   `gorm:"column:PublicTime;comment:视频发布时间戳"`
// }

//用户观看历史记录
type History struct {
	gorm.Model
	UID int64 `gorm:"column:UID;not null;comment:用户uid"`
	VID int64 `gorm:"column:VID;not null;comment:视频vid"`
	Cnt int64 `gorm:"column:Cnt;comment:播放次数"`
}
