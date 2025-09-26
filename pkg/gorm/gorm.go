package gorm

import (
	"dream-vue-admin/global"
	"dream-vue-admin/models/system_model"
	"fmt"
	"log"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitGorm() *gorm.DB {

	database := global.Cfg.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		database.Username,
		database.Password,
		database.IP,
		database.Port,
		database.Database,
	)
	dialector := mysql.Open(dsn)

	//修改日志级别
	gormLogger := logger.Default
	gormLogger.LogMode(logger.Info)

	//链接数据库
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: gormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
	})
	if err != nil {
		log.Panicln("初始化Gorm失败：", err.Error())
		return nil
	}

	//注册模型
	initModel(db)

	return db
}

func initModel(db *gorm.DB) {
	err := db.AutoMigrate(
		// casbin记录权限
		&gormadapter.CasbinRule{},

		//字典
		&system_model.Dict{},
		&system_model.DictValue{},
	)
	if err != nil {
		log.Panicf("初始化Gorm模型失败：%s", err.Error())
	}
}
