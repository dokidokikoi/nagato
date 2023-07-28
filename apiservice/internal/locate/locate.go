package locate

import (
	"encoding/json"
	"time"

	"nagato/apiservice/internal/config"
	"nagato/common/rabbitmq"
	"nagato/common/types"
)

type LocateMessage struct {
	Addr string
	Id   int
}

func Locate(hash string) map[int]string {
	q := rabbitmq.New(config.Config().RabbitMqConfig.Dns())
	q.Publish("dataServers", hash)
	c := q.Consume()

	// 一秒之后关闭消息队列
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()

	// 获取各 dataservice 的分片 id 和地址
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

// 文件是否存在
// 因为将一个文件分成了 3 个数据片和 1 个校验片
// 所以最多允许丢失一个分片
func Exist(hash string) bool {
	return len(Locate(hash)) >= 3
}
