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

// 上传文件
func (s matterSrv) Upload(ctx context.Context, example *model.Matter, data io.Reader) error {
	// 检查是否已经上传
	if locate.Exist(example.Sha256) {
		_, err := s.Get(ctx, &model.Matter{Path: example.Path}, &meta.GetOption{Include: []string{"sha256"}})
		if err != nil {
			err = s.Create(ctx, example)
			return err
		}
		return nil
	}

	// 获取切片数量的服务地址
	servers := heartbeat.ChooseRandomDataServers(stream.ALL_SHARDS, nil)
	if len(servers) < stream.ALL_SHARDS {
		return fmt.Errorf("cannot find enough dataServer")
	}

	tempPutStream, err := stream.NewRSPutStream(servers, example.Sha256, example.Size)
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
	if d != example.Sha256 {
		tempPutStream.Commit(false)
		return errors.New("文件散列值不匹配")
	}

	// 临时文件转正
	tempPutStream.Commit(true)

	// 将文件元信息写入数据库
	return s.Create(ctx, example)
}

// 生成大文件上传 token
func (s matterSrv) GenUploadToken(ctx context.Context, example *model.Matter) (string, error) {
	if locate.Exist(example.Sha256) {
		_, err := s.Get(ctx, &model.Matter{Path: example.Path, UserID: example.UserID}, nil)
		if err != nil {
			err = s.Create(ctx, example)
			return "", err
		}
		return "", nil
	}

	servers := heartbeat.ChooseRandomDataServers(stream.ALL_SHARDS, nil)
	if len(servers) < stream.ALL_SHARDS {
		return "", fmt.Errorf("cannot find enough dataServer")
	}

	rsPutStream, err := stream.NewRSResumablePutStream(servers, example.Name, example.Sha256, example.Size)
	if err != nil {
		return "", err
	}

	err = s.store.Matters().Create(ctx, example, nil)
	if err != nil {
		return "", err
	}

	return rsPutStream.ToToken(), nil
}

// 大文件上传(需要token)
// offset 是已经上传文件的大小
func (s matterSrv) UploadBigMatter(ctx context.Context, token string, offset uint, data io.Reader) error {
	r, err := stream.NewRSResumablePutStreamFromToken(token)
	if err != nil {
		return err
	}

	current, err := r.CurrentSize()
	if err != nil {
		return err
	}

	if current != offset {
		return fmt.Errorf("续传的文件偏移量与已上传文件大小不匹配, current: %d, offset: %d", current, offset)
	}

	bytes := make([]byte, stream.BLOCK_SIZE)
	for {
		n, err := io.ReadFull(data, bytes)
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			return fmt.Errorf("读取缓存错误, err: %s", err.Error())
		}
		current += uint(n)
		if current > r.Size {
			r.Commit(false)
			return fmt.Errorf("上传文件大于文件实际大小, current: %d, size: %d", current, r.Size)
		}

		if n != stream.BLOCK_SIZE && current != r.Size {
			err := r.Check(int64(offset) / 3)
			if err != nil {
				r.Commit(false)
			}
			return err
		}

		r.Write(bytes[:n])

		// 文件全部写入完成
		if current == r.Size {
			// 将最后的缓冲内的数据写入
			r.Flush()
			getStream, err := stream.NewRSResumableGetStream(r.Servers, r.Uuids, r.Size)
			if err != nil {
				return err
			}

			hash, err := tools.CalculateHash(getStream)
			if err != nil {
				return err
			}
			// if hash != r.Hash {
			// 	r.Commit(false)
			// 	return fmt.Errorf("resumable put done but hash mismatch")
			// }

			if locate.Exist(hash) {
				r.Commit(false)
			} else {
				r.Commit(true)
			}

			return nil
		}
	}
}
