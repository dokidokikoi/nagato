package config

import (
	conf "github.com/dokidokikoi/go-common/config"
	"github.com/spf13/viper"
)

var configInfo *config

type config struct {
	FileSystemConfig
	conf.RabbitMqConfig
	conf.ServerConfig
}

func Config() *config {
	return configInfo
}

func GetSpecConfig(key string) *viper.Viper {
	return conf.Config().Sub(key)
}

func init() {
	conf.SetConfig("./internal/conf/application.yml")
	configInfo = &config{
		FileSystemConfig: GetFileSystemInfo(),
		ServerConfig:     conf.GetServerInfo(),
		RabbitMqConfig:   conf.GetRabbitMqInfo(),
	}
}
