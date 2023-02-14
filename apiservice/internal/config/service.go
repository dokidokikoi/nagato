package config

import conf "github.com/dokidokikoi/go-common/config"

type ServiceConfig []string

const (
	serviceKey string = "services"
)

func GetServiceInfo() ServiceConfig {
	serviceConfig := conf.Config().GetStringSlice(serviceKey)

	return serviceConfig
}
