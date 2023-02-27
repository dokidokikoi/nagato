package config

import "fmt"

type RpcConfig struct {
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func (r RpcConfig) Address() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}
