package config

import (
	conf "github.com/dokidokikoi/go-common/config"
	"github.com/spf13/viper"
)

var configInfo *config

type config struct {
	conf.EsConfig
	conf.RedisConfig
	conf.ServerConfig
	RpcConfig
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
		EsConfig:     conf.GetEsInfo(),
		RedisConfig:  conf.GetRadisInfo(),
		ServerConfig: conf.GetServerInfo(),
		RpcConfig:    GetRpcInfo(),
		EtcdConfig:   GetEtcdInfo(),
	}
}
