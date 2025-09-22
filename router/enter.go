package router

import (
	"dream-vue-admin/global"
	"dream-vue-admin/router/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// InitRouterAndRun 初始化路由并运行
func InitRouterAndRun() {
	engine := gin.Default()

	//注册路由
	//自动CURD之后需要在此注册路由才可以
	{
		v1group := engine.Group("/v1")
		//用户接口
		user.InitUserApiRouter(v1group)
	}

	//启动服务
	if err := engine.Run(fmt.Sprintf(":%d", global.Cfg.Server.Port)); err != nil {
		log.Panicln("启动服务失败：", err.Error())
	}
}
