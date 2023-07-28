package stream

import (
	"encoding/base64"
	"fmt"
	"io"

	"nagato/common/tools"
)

const (
	// 数据片个数
	DATA_SHARDS = 3
	// 校验片个数
	PARITY_SHARDS = 1
	// 总个数
	ALL_SHARDS = DATA_SHARDS + PARITY_SHARDS
)

type RSPutStream struct {
	*encoder
}

// 将完成上传的临时文件转正
func (s *RSPutStream) Commit(success bool) {
	// 将缓存里的数据完全写入， 传入hash校验分片
	s.Flush()
	for i := range s.writers {
		s.writers[i].(*TempPutStream).Commit(success, base64.StdEncoding.EncodeToString(s.encoder.hashs[i].Sum(nil)))
	}
}

func NewRSPutStream(dataServers []string, hash string, size uint) (*RSPutStream, error) {
	if len(dataServers) != ALL_SHARDS {
		return nil, fmt.Errorf("dataServers number mismatch")
	}

	// 每个分片的大小,size/DATA_SHARDS 再向上取整
	perShard := tools.CalculatePerShard(size, DATA_SHARDS)
	writers := make([]io.Writer, ALL_SHARDS)
	var err error
	for i := range writers {
		writers[i], err = NewTempPutStream(dataServers[i], fmt.Sprintf("%s.%d", hash, i), perShard)
		if err != nil {
			return nil, err
		}
	}
	enc := NewEncoder(writers)

	return &RSPutStream{enc}, nil
}
