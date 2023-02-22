package config

import (
	conf "github.com/dokidokikoi/go-common/config"
	"github.com/spf13/viper"
)

var configInfo = &config{}

type config struct {
	EsConfig     conf.EsConfig     `mapstructure:"elasticsearch"`
	RedisConfig  conf.RedisConfig  `mapstructure:"redis"`
	ServerConfig conf.ServerConfig `mapstructure:"server"`
	RpcConfig    RpcConfig         `mapstructure:"rpc"`
	EtcdConfig   conf.EtcdConfig   `mapstructure:"etcd"`
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
