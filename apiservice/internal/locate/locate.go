package locate

import (
	"encoding/json"
	"nagato/apiservice/internal/config"
	"nagato/common/rabbitmq"
	"nagato/common/types"
	"time"
)

type LocateMessage struct {
	Addr string
	Id   int
}

func Locate(name string) map[int]string {
	q := rabbitmq.New(config.Config().RabbitMqConfig.Dns())
	q.Publish("dataServers", name)
	c := q.Consume()

	// 一秒之后关闭消息队列
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()

	locateInfo := make(map[int]string)
	for i := 0; i < 4; i++ {
		msg := <-c
		if len(msg.Body) == 0 {
			return locateInfo
		}
		var info types.LocateMessage
		json.Unmarshal(msg.Body, &info)
		locateInfo[info.Id] = info.Addr
	}

	return locateInfo
}

func Exist(name string) bool {
	return len(Locate(name)) >= 3
}
