package config

type RabbitMqConf struct {
	Host     string
	Port     int
	Username string `validate:"required"`
	Password string
}

const (
	rabbitMqKey string = "rabbitmq"
)

var RabbitMqConfig = &RabbitMqConf{
	Port: 5672, Host: "127.0.0.1", Username: "harukaze", Password: "123456",
}

func init() {
	rabbitMqConfig := GetSpecConfig(rabbitMqKey)
	if rabbitMqConfig != nil {
		rabbitMqConfig.Unmarshal(RabbitMqConfig)
	}
}
