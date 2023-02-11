package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Config() *viper.Viper {
	return config
}

func SetConfig(filename string) {
	config = viper.New()
	config.SetConfigFile(filename)
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func init() {
	SetConfig("./internal/conf/application.yml")
}

func GetSpecConfig(key string) *viper.Viper {
	return config.Sub(key)
}
