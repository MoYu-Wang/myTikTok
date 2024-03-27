package dao

//用户表
type User struct {
	UserID   int64  `gorm:"column:UserID;not null;comment:用户uid"`
	UserName string `gorm:"column:UserName;not null;comment:昵称"`
	PassWord string `gorm:"column:PassWord;not null;comment:密码"`
	IphoneID string `gorm:"column:IphoneID;not null;comment:手机号"`
}

//视频表
type Video struct {
	VideoID    int64   `gorm:"column:VideoID;comment:视频vid"`
	VideoName  string  `gorm:"column:VideoName;comment:视频名称"`
	VideoLink  string  `gorm:"column:VideoLink;comment:视频链接"`
	UserID     int64   `gorm:"column:UserID;comment:发布人uid"`
	Tags       string  `gorm:"column:Tags;comment:视频所有标签"`
	Weight     float32 `gorm:"column:Weight;comment:视频初始权重"`
	PublicTime int64   `gorm:"column:PublicTime;comment:视频开始发布时间"`
}

//喜爱列表
type Favorite struct {
	UserID  int64 `gorm:"column:UserID;not null;comment:用户uid"`
	VideoID int64 `gorm:"column:VideoID;comment:视频vid"`
}

//视频评论表
type CommentList struct {
	UserID      int64  `gorm:"column:UserID;not null;comment:评论用户uid"`
	VideoID     int64  `gorm:"column:VideoID;comment:视频vid"`
	CommentText string `gorm:"column:CommentText;comment:评论文本"`
	CommentTime int64  `gorm:"column:CommentTime;comment:评论时间"`
}

//关注列表
type CareList struct {
	UserID     int64 `gorm:"column:UserID;comment:用户uid"`
	CareUserID int64 `gorm:"column:CareUserID;comment:被关注用户uid"`
}

//用户观看标签表
type UserLookTag struct {
	UserID   int64  `gorm:"column:UserID;comment:用户uid"`
	Tag      string `gorm:"column:Tag;comment:用户观看标签"`
	PlayTime int64  `gorm:"column:PlayTime;comment:标签被播放时间(单位:时间戳)"`
}

//用户观看历史记录
type History struct {
	UserID   int64 `gorm:"column:UserID;not null;comment:用户uid"`
	VideoID  int64 `gorm:"column:VideoID;not null;comment:视频vid"`
	Cnt      int64 `gorm:"column:Cnt;comment:播放次数"`
	LastTime int64 `gorm:"column:LastTime;comment:上一次播放时间"`
}
