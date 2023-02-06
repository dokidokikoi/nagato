package heartbeat

import (
	"nagato/common/rabbitmq"
	"nagato/dataservice/internal/config"
	"os"
	"time"
)

func StartHeartbeat() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()

	for {
		q.Publish("apiServers", config.LISTEN_ADDRESS)
		time.Sleep(5 * time.Second)
	}
}
