package config

import (
	"flag"
	"fmt"

	conf "github.com/dokidokikoi/go-common/config"
	"github.com/spf13/viper"
)

var configInfo = &config{}

type config struct {
	PGConfig       conf.PGConfig       `mapstructure:"postgresql"`
	RabbitMqConfig conf.RabbitMqConfig `mapstructure:"rabbitmq"`
	RedisConfig    conf.RedisConfig    `mapstructure:"redis"`
	ServerConfig   conf.ServerConfig   `mapstructure:"server"`
	ServiceConfig  ServiceConfig       `mapstructure:"services"`
	EtcdConfig     conf.EtcdConfig     `mapstructure:"etcd"`

	EsConfig conf.EsConfig `mapstructure:"elasticsearch"`
}

func Config() *config {
	return configInfo
}

func GetSpecConfig(key string) *viper.Viper {
	return conf.Config().Sub(key)
}

func Init(configFile string) {
	flag.StringVar(&configInfo.ServerConfig.Host, "h", "", "服务host")
	flag.IntVar(&configInfo.ServerConfig.Port, "p", 10000, "服务port")
	flag.Parse()

	conf.Parse(configFile, configInfo)
	fmt.Printf("config: %+v", configInfo.ServerConfig)
}
