package config

import "fmt"

type EtcdConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func (e EtcdConfig) Address() string {
	return fmt.Sprintf("%s:%d", e.Host, e.Port)
}
