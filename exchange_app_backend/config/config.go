package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}
	Redis struct {
		Addr     string
		DB       int
		Password string
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config") // 配置文件名 (不带后缀)
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config") // 在 config 子目录查找
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config file failed, err:%v\n", err)
	}

	AppConfig = &Config{}
	// 将读取到的配置反序列化到 Conf 结构体中
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("unmarshal config failed, err:%v\n", err)
	}

	// 监控配置文件变化并热加载
	viper.WatchConfig()

	initDb()
}
