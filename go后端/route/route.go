package route

import (
	"WebVideoServer/service"

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
	r.GET("/top", service.Response_Top)
	r.GET("/care", service.Response_Care)
	r.GET("/sport", service.Response_Sport)
	r.GET("/game", service.Response_Game)
	r.GET("/referee", service.Response_Referee)

	//启动(端口为11316)
	r.Run(":11316")
}
