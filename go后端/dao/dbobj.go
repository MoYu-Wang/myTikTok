package dao

import "gorm.io/gorm"

//用户表
type User struct {
	gorm.Model
	UserID   int64  `gorm:"column:UserID;not null;comment:用户名uid"`
	Username string `gorm:"column:Username;not null;comment:昵称"`
	PassWord string `gorm:"column:PassWord;not null;comment:密码"`
}

//视频表
type Vedio struct {
	gorm.Model
	VedioID    int64    `gorm:"column:VedioID;comment:视频vid"`
	Vedioname  string   `gorm:"column:Vedioname;comment:视频名称"`
	Vediolink  string   `gorm:"column:Vediolink;comment:视频链接"`
	UserID     int64    `gorm:"column:UserID;comment:发布人uid"`
	Tag        []string `gorm:"column:Tag;comment:视频标签数组"`
	Weight     float64  `gorm:"column:Weight;comment:视频初始权重"`
	PublicTime int64    `gorm:"column:PublicTime;comment:视频开始发布时间"`
}

//关注列表
type CareList struct {
	gorm.Model
	UserID     int64 `gorm:"column:UserID;comment:用户uid"`
	CareUserID int64 `gorm:"column:CareUserID;comment:被关注用户uid"`
}

//用户观看标签表
type UserLookTag struct {
	gorm.Model
	UserID   int64  `gorm:"column:UserID;comment:用户uid"`
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
	UserID  int64 `gorm:"column:UserID;not null;comment:用户uid"`
	VedioID int64 `gorm:"column:VedioID;not null;comment:视频vid"`
	Cnt     int64 `gorm:"column:Cnt;comment:播放次数"`
}
