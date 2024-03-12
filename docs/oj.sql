#1.重新创建表空间
drop database mytiktok;
create database mytiktok;
#2.创建用户
#CREATE USER 'MyTikTok'@'localhost' IDENTIFIED BY '2020tiktok';

#3.赋予用户表空间的权限
GRANT ALL PRIVILEGES ON mytiktok.* TO 'MyTikTok'@'localhost';
#4.创建表
use mytiktok

#此处存储用户的基本信息
create table User(
    UserID bigint primary key comment '用户id',
    UserName varchar(20) comment '用户昵称',
    PassWord varchar(20) comment '用户密码',
    IphoneID varchar(20) comment '手机号'
);

#此处存储视频的基本信息
create table Video(
    VideoID bigint primary key comment '视频id',
    VideoName varchar(100) comment '视频名称',
    VideoLink varchar(100) comment '视频链接',
    UserID bigint comment '视频发布人id',
    Tags  text comment '视频标签',
    Weight float(32,16) comment'视频初始权重',
    PublicTime bigint comment'视频发布时间',

    FOREIGN KEY (UserID) REFERENCES User(UserID)
);

#喜爱列表
create table Favorite(
    UserID bigint comment '用户id',
    VideoID bigint comment '视频id',
    
    FOREIGN KEY (UserID) REFERENCES User(UserID),
    FOREIGN KEY (VideoID) REFERENCES Video(VideoID)
);

#视频评论表
create table CommentList(
    UserID bigint comment '用户id',
    VideoID bigint comment '视频id',
    CommentText text comment '评论文本',
    CommentTime bigint comment '评论时间',

    FOREIGN KEY (UserID) REFERENCES User(UserID),
    FOREIGN KEY (VideoID) REFERENCES Video(VideoID)
);

#关注列表
create table CareList(
    UserID bigint comment '用户id',
    CareUserID bigint comment '被关注用户id',

    FOREIGN KEY (UserID) REFERENCES User(UserID),
    FOREIGN KEY (CareUserID) REFERENCES User(UserID)
);

#用户观看标签表
create table UserLookTag(
    UserID bigint comment '用户id',
    Tag varchar(50) comment '标签',
    PlayTime bigint comment '标签被播放时间',

    FOREIGN KEY (UserID) REFERENCES User(UserID)
);

#用户观看历史记录
create table History(
    UserID bigint comment '用户id',
    VideoID bigint comment '视频id',
    Cnt int comment '播放次数',
    LastTime bigint comment '上一次播放时间',

    FOREIGN KEY (UserID) REFERENCES User(UserID),
    FOREIGN KEY (VideoID) REFERENCES Video(VideoID)
);