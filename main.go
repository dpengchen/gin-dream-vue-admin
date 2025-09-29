package main

import (
	"dream-vue-admin/config"
	_ "dream-vue-admin/docs"
	"dream-vue-admin/global"
	"dream-vue-admin/pkg/casbin"
	"dream-vue-admin/pkg/gorm"
	"dream-vue-admin/pkg/zap"
	"dream-vue-admin/router"
)

// @title Dream Vue Admin API
// @version 1.0
// @description 这是一个基于Gin框架的API服务
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://github.com/swaggo/gin-swagger
// @contact.email

// @license.name MIT

// @host localhost:8080
// @BasePath /v1
// @schemes http
func main() {
	//初始化配置
	global.Cfg = config.InitApplicationCfg()
	//初始化日志
	global.Log = zap.InitZapLog()
	//初始化数据库
	global.Db = gorm.InitGorm()
	//初始化casbin
	global.Enforcer = casbin.InitCasbin(global.Db)

	//启动服务
	router.InitRouterAndRun()
}
