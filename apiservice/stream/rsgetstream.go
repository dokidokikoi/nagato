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

func NewRSGetStream(locateInfo map[int]string, dataServers []string, hash string, size uint) (*RSGetStream, error) {
	if len(locateInfo)+len(dataServers) != ALL_SHARDS {
		return nil, fmt.Errorf("dataServers number mismatch")
	}

	readers := make([]io.Reader, ALL_SHARDS)
	for i := 0; i < ALL_SHARDS; i++ {
		server := locateInfo[i]
		if server == "" {
			locateInfo[i] = dataServers[0]
			dataServers = dataServers[1:]
			continue
		}
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
