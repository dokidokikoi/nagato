package heartbeat

import (
	"nagato/common/rabbitmq"
	"nagato/dataservice/internal/config"
	"time"
)

func StartHeartbeat() {
	q := rabbitmq.New(rabbitmq.RABBITMQ_SERVER)
	defer q.Close()

	for {
		q.Publish("apiServers", config.LISTEN_ADDRESS)
		time.Sleep(5 * time.Second)
	}
}
