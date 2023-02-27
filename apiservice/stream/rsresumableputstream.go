package stream

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"nagato/apiservice/internal/model"

	dataRpc "nagato/apiservice/rpc/client/data"
)

type RSResumablePutStream struct {
	*RSPutStream
	*model.ResumableToken
}

func (s *RSResumablePutStream) ToToken() string {
	b, _ := json.Marshal(s)
	return base64.StdEncoding.EncodeToString(b)
}

// 获取已上传文件大小
func (s *RSResumablePutStream) CurrentSize() (uint, error) {
	dataRpc, err := dataRpc.GetDataClient(s.Servers[0])
	if err != nil {
		return 0, err
	}
	perShard, err := dataRpc.HeadTempFile(context.Background(), s.Uuids[0])
	if err != nil {
		return 0, err
	}

	size := uint(perShard) * DATA_SHARDS
	// 由于rs纠删码计算出的size可能大于文件真实size
	if size > s.Size {
		size = s.Size
	}

	return size, nil
}

func NewRSResumablePutStream(dataServers []string, name, hash string, size uint) (*RSResumablePutStream, error) {
	rsPutStream, err := NewRSPutStream(dataServers, hash, size)
	if err != nil {
		return nil, err
	}

	uuids := make([]string, ALL_SHARDS)
	for i := range uuids {
		uuids[i] = rsPutStream.writers[i].(*TempPutStream).Uuid
	}

	token := &model.ResumableToken{Name: name, Size: size, Hash: hash, Servers: dataServers, Uuids: uuids}
	return &RSResumablePutStream{
		RSPutStream:    rsPutStream,
		ResumableToken: token,
	}, nil
}

func NewRSResumablePutStreamFromToken(token string) (*RSResumablePutStream, error) {
	resumableStr, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	var t model.ResumableToken
	err = json.Unmarshal(resumableStr, &t)
	if err != nil {
		return nil, err
	}

	writers := make([]io.Writer, ALL_SHARDS)
	for i := range writers {
		writers[i] = &TempPutStream{
			Server: t.Servers[i],
			Uuid:   t.Uuids[i],
		}
	}

	enc := NewEncoder(writers)

	return &RSResumablePutStream{&RSPutStream{enc}, &t}, nil
}
