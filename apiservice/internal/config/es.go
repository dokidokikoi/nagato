package config

import "fmt"

type EsConf struct {
	Host     string
	Port     int
	Cert     string `validate:"required"`
	Username string `validate:"required"`
	Password string
}

func (s EsConf) Address() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

func (s EsConf) Url() string {
	return fmt.Sprintf("https://%s", s.Address())
}

const (
	esKey string = "elasticsearch"
)

var EsConfig = &EsConf{
	Port: 9200, Host: "127.0.0.1", Cert: "http_ca.crt", Username: "elastic", Password: "ADR*piFezssmbUhhN8*S",
}

func init() {
	esConfig := GetSpecConfig(esKey)
	if esConfig != nil {
		esConfig.Unmarshal(EsConfig)
	}
}
