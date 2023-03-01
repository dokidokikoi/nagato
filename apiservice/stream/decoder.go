package stream

import (
	"crypto/sha256"
	"hash"
	"io"

	"github.com/klauspost/reedsolomon"
)

type decoder struct {
	readers []io.Reader
	writers []io.Writer
	hashs   []hash.Hash
	enc     reedsolomon.Encoder
	size    uint
	// 用于缓存读取的数据
	cache     []byte
	cacheSize int
	// 当前已经读取了多少字节
	total uint
}

func (d *decoder) Read(p []byte) (n int, err error) {
	// cache无数据，从输入流中读入cache
	if d.cacheSize == 0 {
		err := d.getData()
		if err != nil {
			return 0, err
		}
	}

	// 从cache中读取数据
	length := len(p)
	if d.cacheSize < length {
		length = d.cacheSize
	}
	d.cacheSize -= length
	copy(p, d.cache[:length])
	d.cache = d.cache[length:]

	return length, nil
}

// 从输入流读数据到cache中
func (d *decoder) getData() error {
	// 文件读取完毕
	if d.total == d.size {
		return io.EOF
	}
	shards := make([][]byte, ALL_SHARDS)
	repairIds := make([]int, 0)
	for i := range shards {
		if d.readers[i] == nil {
			// 记录需要修复的分片id
			repairIds = append(repairIds, i)
		} else {
			// 从输入流中读取数据
			shards[i] = make([]byte, BLOCK_PER_SHARD)
			n, err := io.ReadFull(d.readers[i], shards[i])
			if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
				shards[i] = nil
			} else if n != BLOCK_PER_SHARD {
				// 未将缓冲读满，缩减切片大小。防止将零值读入cache
				shards[i] = shards[i][:n]
			}
		}
	}

	// 重新生成rs分片
	err := d.enc.Reconstruct(shards)
	if err != nil {
		return err
	}

	// 将待修复的分片写入输出流
	for i := range repairIds {
		id := repairIds[i]
		d.writers[id].Write(shards[id])
		d.hashs[id].Write(shards[id])
	}

	// 遍历 DATA_SHARDS 个数据分片,将每个分片中的数据添加到缓存cache中,
	// 修改缓存当前的大小cacheSize以及当前已经读取的全部数据的大小total。
	for i := 0; i < DATA_SHARDS; i++ {
		shardSize := uint(len(shards[i]))

		// 防止填充数据被当成原始对象数据返回
		if d.total+shardSize > d.size {
			shardSize -= d.total + shardSize - d.size
		}
		d.cache = append(d.cache, shards[i][:shardSize]...)
		d.cacheSize += int(shardSize)
		d.total += shardSize
	}

	return nil
}

func NewDecoder(readers []io.Reader, writers []io.Writer, size uint) *decoder {
	enc, _ := reedsolomon.New(DATA_SHARDS, PARITY_SHARDS)
	hashs := make([]hash.Hash, len(writers))
	// 准备计算每个分片的散列值，防止传输过程中导致数据缺失
	for i, _ := range hashs {
		hashs[i] = sha256.New()
	}
	return &decoder{readers, writers, hashs, enc, size, nil, 0, 0}
}
