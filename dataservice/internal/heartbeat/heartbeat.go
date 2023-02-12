package heartbeat

import (
	"nagato/common/rabbitmq"
	"nagato/dataservice/internal/config"
	"time"
)

func StartHeartbeat() {
	q := rabbitmq.New(config.Config().RabbitMqConfig.Dns())
	defer q.Close()

	for {
		q.Publish("apiServers", config.Config().ServerConfig.Address())
		time.Sleep(5 * time.Second)
	}
}
