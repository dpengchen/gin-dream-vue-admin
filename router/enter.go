package router

import (
	"dream-vue-admin/global"
	"dream-vue-admin/middleware"
	"dream-vue-admin/router/initalize"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// InitRouterAndRun 初始化路由并运行
func InitRouterAndRun() {
	engine := gin.Default()

	{
		group := engine.Group("")
		// 使用鉴权中间件
		group.Use(middleware.Auth())
		//注册路由
		//v1版本
		initalize.InitV1Router(group)
	}

	//启动服务
	if err := engine.Run(fmt.Sprintf("0.0.0.0:%d", global.Cfg.Server.Port)); err != nil {
		log.Panicln("启动服务失败：", err.Error())
	}
}
