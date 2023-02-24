package stream

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"nagato/apiservice/internal/model"
	"nagato/common/tools"
	"net/http"
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
	resp, err := http.Head(fmt.Sprintf("http://%s/data/file/temp/%s", s.Servers[0], s.Uuids[0]))
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("发送数据服务head请求返回状态码不为200, code: %d", resp.StatusCode)
	}

	size := tools.GetSizeFromHeader(resp.Header) * DATA_SHARDS
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
