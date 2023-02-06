package locate

import (
	"nagato/common/rabbitmq"
	"strconv"
	"time"
)

func Locate(name string) string {
	q := rabbitmq.New(rabbitmq.RABBITMQ_SERVER)
	q.Publish("dataServers", name)
	c := q.Consume()

	// 一秒之后关闭消息队列
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()

	msg := <-c
	s, _ := strconv.Unquote(string(msg.Body))

	return s
}

func Exist(name string) bool {
	return Locate(name) != ""
}
