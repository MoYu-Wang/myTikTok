package route

import (
	"WebVideoServer/web/service"

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
		apiRouter.POST("/user/register") //注册用户
		apiRouter.POST("/user/login")    //用户登录

		apiRouter.GET("/top", service.Response_Top)           //热点
		apiRouter.GET("/care", service.Response_Care)         //关注
		apiRouter.GET("/dBc", service.Response_DBc)           //直播
		apiRouter.GET("/shopping", service.Response_Shopping) //商城
		apiRouter.GET("/referee", service.Response_Referee)   //推荐
	}

	//启动(端口为11316)
	r.Run(":11316")
}
