package gorm

import (
	"dream-vue-admin/global"
	"fmt"
	"log"

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

	return db
}
