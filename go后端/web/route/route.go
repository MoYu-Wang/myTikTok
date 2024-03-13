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
			userRouter.POST("register", service.UserRegister)    //注册用户
			userRouter.POST("login", service.UserLogin)          //用户登录
			userRouter.POST("update", service.UserUpdate)        //用户修改信息
			userRouter.POST("base", service.UserBase)            //获取本用户基本信息
			userRouter.POST("forgetpwd", service.PasswordForget) //找回密码
			userRouter.POST("delete", service.UserDelete)        //用户注销

			userRouter.GET("info", service.UserInfo) //获取用户信息

		}

		apiRouter.GET("/top", service.Response_Top)           //热点
		apiRouter.GET("/care", service.Response_Care)         //关注
		apiRouter.GET("/dBc", service.Response_DBc)           //直播
		apiRouter.GET("/shopping", service.Response_Shopping) //商城
		apiRouter.GET("/referee", service.Response_Referee)   //推荐
	}

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	//启动(端口为11316)
	r.Run(":11316")
}
