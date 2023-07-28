package stream

import (
	"crypto/sha256"
	"hash"
	"io"

	"github.com/klauspost/reedsolomon"
)

const (
	BLOCK_PER_SHARD = 8000                          // 分片缓冲大小
	BLOCK_SIZE      = BLOCK_PER_SHARD * DATA_SHARDS // 缓冲大小
)

// 这一层用于将文件数据经过 rs 处理
type encoder struct {
	writers []io.Writer         // 对应每个数据服务的 writer
	hashs   []hash.Hash         // 每片数据的散列值
	enc     reedsolomon.Encoder // rs
	cache   []byte              // 缓冲切片
}

// 使用 rs 纠删码处理文件数据，并将分片写入到不同的数据服务
// 实现 io.Writer 接口
func (e *encoder) Write(p []byte) (n int, err error) {
	length := len(p)
	current := 0
	for length != 0 {
		next := BLOCK_SIZE - len(e.cache)
		if next > length {
			next = length
		}
		e.cache = append(e.cache, p[current:current+next]...)

		// 为提高效率，将缓冲写满后才处理数据
		// 提交文件时，会再单独调用一次 Flush 方法
		if len(e.cache) == BLOCK_SIZE {
			e.Flush()
		}
		current += next
		length -= next
	}
	return len(p), nil
}

// 使用 rs 纠删码处理文件数据，并将分片写入到不同的数据服务
func (e *encoder) Flush() {
	if len(e.cache) == 0 {
		return
	}

	// 分片
	shards, _ := e.enc.Split(e.cache)
	e.enc.Encode(shards)

	for i := range shards {
		// 写入到数据服务磁盘
		e.writers[i].Write(shards[i])
		// 计算分片散列值
		e.hashs[i].Write(shards[i])
	}
	e.cache = []byte{}
}

func NewEncoder(writers []io.Writer) *encoder {
	// 初始化 rs，设置数据片和校验片大小
	enc, _ := reedsolomon.New(DATA_SHARDS, PARITY_SHARDS)
	hashs := make([]hash.Hash, len(writers))
	for i := range hashs {
		hashs[i] = sha256.New()
	}
	return &encoder{writers, hashs, enc, nil}
}
