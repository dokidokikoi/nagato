package config

import (
	conf "github.com/dokidokikoi/go-common/config"
	"github.com/spf13/viper"
)

var configInfo = &config{}

type config struct {
	FileSystemConfig FileSystemConfig    `mapstructure:"filesystem"`
	RabbitMqConfig   conf.RabbitMqConfig `mapstructure:"rabbitmq"`
	ServerConfig     conf.ServerConfig   `mapstructure:"server"`
}

func Config() *config {
	return configInfo
}

func GetSpecConfig(key string) *viper.Viper {
	return conf.Config().Sub(key)
}

func Init(configFile string) {
	conf.Parse(configFile, configInfo)
}
