package router

import (
	"dream-vue-admin/global"
	"dream-vue-admin/router/initalize"
	"fmt"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// InitRouterAndRun 初始化路由并运行
func InitRouterAndRun() {
	engine := gin.Default()

	{
		group := engine.Group("/v1")
		// 使用鉴权中间件
		//group.Use(middleware.Auth())
		//注册路由
		//v1版本
		initalize.InitV1Router(group)

		// 注册swagger路由
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	//启动服务
	if err := engine.Run(fmt.Sprintf(":%d", global.Cfg.Server.Port)); err != nil {
		log.Panicln("启动服务失败：", err.Error())
	}
}