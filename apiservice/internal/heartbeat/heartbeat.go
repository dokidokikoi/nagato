package heartbeat

import (
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
	q := rabbitmq.New(config.Config().RabbitMqConfig.Dns())
	defer q.Close()

	q.Bind("apiServers")
	c := q.Consume()

	go removeExpireDataServer()
	for msg := range c {
		// fmt.Println("---------------got message--------------")
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

	ds := make([]string, 0)
	for k := range dataServers {
		ds = append(ds, k)
	}

	return ds
}

func ChooseRandomDataServers(n int, exclude map[int]string) []string {
	candidates := make([]string, 0)
	reverseExcludeMap := make(map[string]int)
	for id, addr := range exclude {
		reverseExcludeMap[addr] = id
	}
	servers := GetDataServers()
	for i := range servers {
		s := servers[i]
		_, excluded := reverseExcludeMap[s]
		if !excluded {
			candidates = append(candidates, s)
		}
	}

	length := len(candidates)
	if length < n {
		return candidates
	}

	ds := make([]string, n)
	p := rand.Perm(length)
	for i := 0; i < n; i++ {
		ds[i] = candidates[p[i]]
	}

	return ds
}
