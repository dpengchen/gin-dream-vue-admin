package casbin

import (
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func InitCasbin(db *gorm.DB) *casbin.Enforcer {

	gormAdapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Panicf("初始化Casbin失败：%s", err.Error())
	}
	e, err := casbin.NewEnforcer("config/rbac_model.conf", gormAdapter)
	if err != nil {
		log.Panicf("初始化Casbin失败：%s", err.Error())

	}
	return e
}
