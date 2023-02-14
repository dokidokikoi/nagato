package config

import "fmt"

type EtcdConfig struct {
	Host string
	Port int
}

func (e EtcdConfig) Address() string {
	return fmt.Sprintf("%s:%d", e.Host, e.Port)
}

const (
	etcdKey string = "etcd"
)

func GetEtcdConfig() EtcdConfig {
	etcdConfig := &EtcdConfig{}
	conf := GetSpecConfig(etcdKey)
	if conf != nil {
		conf.Unmarshal(etcdConfig)
	}

	return *etcdConfig
}
