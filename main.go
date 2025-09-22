package main

import (
	"dream-vue-admin/config"
	"dream-vue-admin/global"
	"dream-vue-admin/pkg/gorm"
	"dream-vue-admin/pkg/zap"
	"dream-vue-admin/router"
)

func main() {
	//初始化配置
	global.Cfg = config.InitApplicationCfg()
	//初始化日志
	global.Log = zap.InitZapLog()
	//初始化数据库
	global.Db = gorm.InitGorm()

	//启动服务
	router.InitRouterAndRun()
}
