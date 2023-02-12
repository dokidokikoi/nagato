package heartbeat

import (
	"fmt"
	"math/rand"
	"nagato/apiservice/internal/config"
	"nagato/common/rabbitmq"
	"strconv"
	"sync"
	"time"
)

var (
	dataServers = make(map[string]time.Time)
	mutex       sync.Mutex
)

func ListenHeartbeat() {
	dns := fmt.Sprintf(rabbitmq.RABBITMQ_SERVER_TEMPLATE,
		config.RabbitMqConfig.Username,
		config.RabbitMqConfig.Password,
		config.RabbitMqConfig.Host,
		config.RabbitMqConfig.Port)
	q := rabbitmq.New(dns)
	defer q.Close()

	q.Bind("apiServers")
	c := q.Consume()

	go removeExpireDataServer()
	for msg := range c {
		fmt.Println("---------------got message--------------")
		dataServer, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}

		mutex.Lock()
		dataServers[dataServer] = time.Now()
		mutex.Unlock()
	}
}

func removeExpireDataServer() {
	for {
		time.Sleep(time.Second * 5)
		mutex.Lock()
		for k, v := range dataServers {
			if v.Add(10 * time.Second).Before(time.Now()) {
				delete(dataServers, k)
			}
		}
		mutex.Unlock()
	}
}

func GetDataServers() []string {
	mutex.Lock()
	defer mutex.Unlock()

	ds := make([]string, len(dataServers))
	for k := range dataServers {
		ds = append(ds, k)
	}

	return ds
}

func ChooseRandomDataServer() string {
	ds := GetDataServers()
	n := len(ds)
	if n == 0 {
		return ""
	}

	return ds[rand.Intn(n)]
}
