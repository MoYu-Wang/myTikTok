package io

//请求参数

//注册参数
type ParamRegister struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	IphoneID string `json:"iphoneID"`
}

//登录参数
type ParamLogin struct {
	UserID   int64  `json:"userID"`
	PassWord string `json:"password"`
	IphoneID string `json:"iphoneID"`
}

//修改个人信息参数
type ParamUpdate struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	IphoneID string `json:"iphoneID"`
	Token    string `json:"token"`
}

//找回密码参数
type ParamForgetpwd struct {
	UserID   int64  `json:"userID"`
	IphoneID string `json:"iphoneID"`
}

//用户注销请求
type ParamUserDelete struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

// //用户基本请求
// type UserBaseReq struct {
// 	Token string `json:"token"`
// }

//用户信息请求参数
type UserInfoReq struct {
	UserID   int64  `json:"userID"`
	UserName string `json:"userName"`
	Token    string `json:"token"`
}

//用户上传视频请求参数
type UserUpLoadVideoReq struct {
	Token string `json:"token"`

	VideoName string `json:"videoName"`
	VideoTags string `json:"videoTags"`
	VideoLink string `json:"videoLink"`
}

//点赞视频请求
type FavoriteVideoReq struct {
	Token      string `json:"token"`
	VideoID    string `json:"videoID"`
	IsFavorite int64  `json:"isFavorite"`
}

//评论视频请求
type CommentVideoReq struct {
	Token       string `json:"token"`
	VideoID     string `json:"videoID"`
	CommentText string `json:"commentText"`
}

//删除评论请求
type DeleteCommentReq struct {
	Token     string `json:"token"`
	VideoID   string `json:"videoID"`
	CommentID string `json:"commentID"`
}

//用户操作视频请求参数
type OperateVideoReq struct {
	Token        string   `json:"token"`
	VideoID      string   `json:"videoID"`
	WatchTime    int64    `json:"watchTime"`
	IsFavorite   int64    `json:"isFavorite"`
	CommentNum   int64    `json:"commentNum"`
	CommentTexts []string `json:"commentTexts"`
}

//模糊查询视频请求
type SearchVideoReq struct {
	SearchText string `json:"searchText"`
}

//用户作品请求
type UserWorkReq struct {
	Token  string `json:"token"`
	UserID int64  `json:"userID"`
}

//关注用户请求
type CareUserReq struct {
	Token   string `json:"token"`
	UserID  int64  `json:"userID"`
	Operate int64  `json:"operate"`
}
