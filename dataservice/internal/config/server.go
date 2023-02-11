package config

import "fmt"

type ServerConf struct {
	Host string
	Port int
}

func (s ServerConf) Address() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

const (
	severKey string = "server"
)

var ServerConfig = &ServerConf{
	Host: "127.0.0.1", Port: 8100,
}

func init() {
	serverConfig := GetSpecConfig(severKey)
	if serverConfig != nil {
		serverConfig.Unmarshal(ServerConfig)
	}
}
