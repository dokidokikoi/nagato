package stream

import (
	"encoding/base64"
	"fmt"
	"io"
	"nagato/common/tools"
)

type RSGetStream struct {
	*decoder
}

func (s *RSGetStream) Close() {
	for i := range s.writers {
		if s.writers[i] != nil {
			s.writers[i].(*TempPutStream).Commit(true, base64.StdEncoding.EncodeToString(s.hashs[i].Sum(nil)))
		}
	}
}

func (s *RSGetStream) Seek(offset int64, whence int) (int64, error) {
	if whence != io.SeekCurrent {
		panic("only support SeekCurrent")
	}

	// 跳过多余的数据
	for offset != 0 {
		length := int64(BLOCK_SIZE)
		if offset < length {
			length = offset
		}
		buf := make([]byte, length)
		io.ReadFull(s, buf)
		offset -= length
	}

	return offset, nil
}

func NewRSGetStream(locateInfo map[int]string, dataServers []string, hash string, size uint) (*RSGetStream, error) {
	// 如果数据服务数量少于ALL_SHARDS，无法满足rs纠删码
	if len(locateInfo)+len(dataServers) != ALL_SHARDS {
		return nil, fmt.Errorf("dataServers number mismatch")
	}

	readers := make([]io.Reader, ALL_SHARDS)
	for i := 0; i < ALL_SHARDS; i++ {
		server := locateInfo[i]
		// 补全 locateInfo ，塞入所有数据服务的信息
		if server == "" {
			locateInfo[i] = dataServers[0]
			dataServers = dataServers[1:]
			continue
		}
		// 获取正常的数据输入流
		reader, err := NewGetStream(server, fmt.Sprintf("%s.%d", hash, i))
		if err == nil {
			readers[i] = reader
		}
	}

	writers := make([]io.Writer, ALL_SHARDS)
	perShard := tools.CalculatePerShard(size, DATA_SHARDS)
	var err error
	for i := range readers {
		// 如果 reader 为空则创建临时对象写入流用于恢复分片
		if readers[i] == nil {
			writers[i], err = NewTempPutStream(locateInfo[i], fmt.Sprintf("%s.%d", hash, i), perShard)
			if err != nil {
				return nil, err
			}
		}
	}

	dec := NewDecoder(readers, writers, size)
	return &RSGetStream{dec}, nil
}
