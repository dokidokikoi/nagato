package heartbeat

import (
	"fmt"
	"nagato/common/rabbitmq"
	"nagato/dataservice/internal/config"
	"time"
)

func StartHeartbeat() {
	dns := fmt.Sprintf(rabbitmq.RABBITMQ_SERVER_TEMPLATE,
		config.RabbitMqConfig.Username,
		config.RabbitMqConfig.Password,
		config.RabbitMqConfig.Host,
		config.RabbitMqConfig.Port)
	q := rabbitmq.New(dns)
	defer q.Close()

	for {
		q.Publish("apiServers", config.ServerConfig.Address())
		time.Sleep(5 * time.Second)
	}
}
