package locate

import (
	"nagato/common/rabbitmq"
	"nagato/dataservice/internal/config"
	"path/filepath"
	"strconv"
	"sync"
)

// 存储目录下的对象文件
// 避免每次查询文件是否存在时访问磁盘，将所有文件的散列值读入内存
var matterHash = make(map[string]int)
var mutex sync.Mutex

func Locate(hash string) bool {
	mutex.Lock()
	_, ok := matterHash[hash]
	mutex.Unlock()
	return ok
}

func StartLocate() {
	q := rabbitmq.New(config.Config().RabbitMqConfig.Dns())
	defer q.Close()

	q.Bind("dataServers")
	c := q.Consume()

	for msg := range c {
		hash, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}

		exist := Locate(hash)
		if exist {
			q.Send(msg.ReplyTo, config.Config().ServerConfig.Address())
		}
	}
}

func CollectMatters() {
	// 读取存储目录里的所有文件信息
	files, _ := filepath.Glob(config.Config().FileSystemConfig.StoreDir + "*")

	// 获取每个文件的文件名（散列值），加入缓存
	for i := range files {
		hash := filepath.Base(files[i])
		matterHash[hash] = 1
	}
}

// 将用户上传的文件加入内存
func Add(hash string) {
	mutex.Lock()
	defer mutex.Unlock()
	matterHash[hash] = 1
}

// 将文件移出内存
func Del(hash string) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(matterHash, hash)
}
