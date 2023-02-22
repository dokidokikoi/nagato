package matter

import (
	"context"
	"errors"
	"fmt"
	"io"
	"nagato/apiservice/internal/heartbeat"
	"nagato/apiservice/internal/locate"
	"nagato/apiservice/internal/model"
	"nagato/apiservice/stream"
	"nagato/common/tools"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

func (s matterSrv) Upload(ctx context.Context, example *model.Matter, hash string, size int64, data io.Reader) error {
	if locate.Exist(hash) {
		_, err := s.Get(ctx, &model.Matter{Sha256: hash}, &meta.GetOption{Include: []string{"sha256"}})
		if err != nil {
			err = s.Create(ctx, example)
			return err
		}
		return nil
	}

	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return fmt.Errorf("cannot find any dataServer")
	}
	tempPutStream, err := stream.NewTempPutStream(server, hash, size)
	if err != nil {
		return err
	}

	// io.TeeReader() 返回 io.Reader，读取返回的 reader 时，也向 Writer（stream） 中写
	reader := io.TeeReader(data, tempPutStream)
	d, err := tools.CalculateHash(reader)
	if err != nil {
		return err
	}

	// 数据校验，不一致则从临时目录删除
	if d != hash {
		tempPutStream.Commit(false)
		return errors.New("文件散列值不匹配")
	}

	tempPutStream.Commit(true)

	return s.Create(ctx, example)
}
