package locate

import (
	"nagato/common/rabbitmq"
	"nagato/common/types"
	"nagato/dataservice/internal/config"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

// 存储目录下的对象文件
// 避免每次查询文件是否存在时访问磁盘，将所有文件的散列值读入内存
var matterHash = make(map[string]int)
var mutex sync.Mutex

func Locate(hash string) int {
	mutex.Lock()
	id, ok := matterHash[hash]
	mutex.Unlock()
	if !ok {
		return -1
	}

	return id
}

func StartLocate() {
	q := rabbitmq.New(config.Config().RabbitMqConfig.Dns())
	defer q.Close()

	q.Bind("dataServers")
	c := q.Consume()

	for msg := range c {
		hash, err := strconv.Unquote(string(msg.Body))
		if err != nil {
			panic(err)
		}
		id := Locate(hash)
		if id != -1 {
			q.Send(msg.ReplyTo, types.LocateMessage{Addr: config.Config().ServerConfig.Address(), Id: id})
		}
	}
}

func CollectMatters() {
	// 读取存储目录里的所有文件信息
	files, _ := filepath.Glob(config.Config().FileSystemConfig.StoreDir + "*")

	// 获取每个文件的文件名（散列值），加入缓存
	for i := range files {
		file := strings.Split(filepath.Base(files[i]), ".")
		if len(file) != 3 {
			panic(files[i])
		}
		hash := file[0]
		id, err := strconv.Atoi(file[1])
		if err != nil {
			panic(err)
		}
		matterHash[hash] = id
	}
}

// 将用户上传的文件加入内存
func Add(hash string, id int) {
	mutex.Lock()
	defer mutex.Unlock()
	matterHash[hash] = id
}

// 将文件移出内存
func Del(hash string) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(matterHash, hash)
}
