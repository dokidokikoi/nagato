package config

import "fmt"

type RpcConfig struct {
	Name string
	Host string
	Port int
}

func (r RpcConfig) Address() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}

const (
	rpcKey string = "rpc"
)

func GetRpcInfo() RpcConfig {
	rpcConfig := &RpcConfig{}
	conf := GetSpecConfig(rpcKey)
	if conf != nil {
		conf.Unmarshal(rpcConfig)
	}

	return *rpcConfig
}
