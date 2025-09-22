package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func InitApplicationCfg() (cfg *ApplicationConfig) {
	v := viper.New()
	v.SetConfigName("application")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")

	_, err := os.Stat("./config/application.yaml")
	if err != nil {
		log.Println("未找到配置文件！/config/application.yaml")
		return
	}

	if err = v.ReadInConfig(); err != nil {
		log.Panicln("读取配置文件错误！", err.Error())
		return
	}

	cfg = new(ApplicationConfig)
	if err = v.Unmarshal(cfg); err != nil {
		log.Panicln("配置文件解析错误！", err.Error())
		return
	}
	return
}
