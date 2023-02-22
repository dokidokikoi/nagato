package config

import (
	"fmt"
	"testing"
)

type con struct {
	Pg struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Database string `mapstructure:"database"`
		Username string `mapstructure:"username"`
		TimeZone string `mapstructure:"timezone"`
		Password string `mapstructure:"password"`
	} `mapstructure:"postgresql"`
}

func TestInitConfig(t *testing.T) {
	Init("../conf/application.yml")
	fmt.Printf("config: %+v", Config())
}
