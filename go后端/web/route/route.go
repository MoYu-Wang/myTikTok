package route

import (
	"WebVideoServer/web/service"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func OpenRoute() {
	r := gin.Default()

	//配置cors头
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}                                       // 允许所有来源访问，可以根据需要设置具体的来源
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // 允许的 HTTP 方法
	r.Use(cors.New(config))

	//定义路由和处理函数
	apiRouter := r.Group("/myTikTok")
	{
		userRouter := apiRouter.Group("/user")
		{
			userRouter.POST("/register", service.UserRegister) //注册用户
			userRouter.POST("/login", service.UserLogin)       //用户登录
			updateRouter := userRouter.Group("/update")
			{
				updateRouter.POST("/info", service.UserUpdateInfo)         //用户修改信息
				updateRouter.POST("/password", service.UserUpdatePassword) //用户修改密码
			}
			userRouter.POST("/info", service.UserInfo)            //获取用户信息
			userRouter.POST("/forgetpwd", service.PasswordForget) //找回密码
			userRouter.POST("/delete", service.UserDelete)        //用户注销
			userRouter.POST("/works", service.UserWorks)          //用户作品
			userRouter.POST("/care", service.CareUser)            //关注用户

			userRouter.GET("/carelist", service.CareList)       //关注列表
			userRouter.GET("/updatetoken", service.UpdateToken) //更新用户登录信息
			userRouter.GET("/base", service.UserBase)           //获取本用户基本信息
			userRouter.GET("/favorite", service.UserFavorite)   //用户点赞视频列表
			userRouter.GET("/history", service.UserHistory)     //用户观看历史记录

		}

		videoRouter := apiRouter.Group("/video")
		{
			videoRouter.GET("/top", service.TopVideo)         //获取热点视频
			videoRouter.GET("/care", service.CareVideo)       //获取关注视频
			videoRouter.GET("/referee", service.RefereeVideo) //获取推荐视频

			videoRouter.GET("/getsign", service.GetSign)            //获取上传签名
			videoRouter.GET("/getcomment", service.GetVideoComment) //获取视频评论

			videoRouter.POST("/info", service.VideoOperateInfo)            //获取视频操作信息
			videoRouter.POST("/upload", service.UpLoadVideo)               //上传视频
			videoRouter.POST("/search", service.SearchVideo)               //模糊查询视频
			videoRouter.POST("/favorite", service.FavoriteVideo)           //视频点赞
			videoRouter.POST("/comment", service.CommentVideo)             //评论视频
			videoRouter.POST("/deletecomment", service.DeleteVideoComment) //视频删除评论

			videoRouter.POST("/operate", service.OperateVideo) //划走视频后对视频的操作
		}
		apiRouter.GET("/broadcast", service.Broadcast) //直播
		apiRouter.GET("/shopping", service.Shopping)   //商城
	}

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	//启动(端口为11316)
	r.Run(":11316")
}
