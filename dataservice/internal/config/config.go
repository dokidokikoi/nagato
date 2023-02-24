package config

import (
	"flag"
	"fmt"

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
	flag.StringVar(&configInfo.ServerConfig.Host, "h", "", "服务host")
	flag.IntVar(&configInfo.ServerConfig.Port, "p", 10100, "服务port")

	flag.StringVar(&configInfo.FileSystemConfig.StoreDir, "store", "/tmp/objects/", "文件存储目录")
	flag.StringVar(&configInfo.FileSystemConfig.TempDir, "temp", "/tmp/temp/", "临时文件存储目录")
	flag.Parse()

	conf.Parse(configFile, configInfo)

	fmt.Printf("fileSystem: %+v", configInfo.FileSystemConfig)
}
