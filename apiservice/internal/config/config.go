package config

import (
	conf "github.com/dokidokikoi/go-common/config"
	"github.com/spf13/viper"
)

var configInfo *config

type config struct {
	conf.PGConfig
	conf.EsConfig
	conf.RabbitMqConfig
	conf.RedisConfig
	conf.ServerConfig
	ServiceConfig
	EtcdConfig
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
		PGConfig:       conf.GetPgInfo(),
		EsConfig:       conf.GetEsInfo(),
		RabbitMqConfig: conf.GetRabbitMqInfo(),
		RedisConfig:    conf.GetRadisInfo(),
		ServerConfig:   conf.GetServerInfo(),
		ServiceConfig:  GetServiceInfo(),
		EtcdConfig:     GetEtcdConfig(),
	}
}
