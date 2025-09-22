package global

import (
	"dream-vue-admin/config"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Cfg         *config.ApplicationConfig
	Db          *gorm.DB
	RedisClient *redis.Client
	Log         *zap.Logger
)
