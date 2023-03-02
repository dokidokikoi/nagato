package data

import (
	"context"
	"io"
	"log"
	pb "nagato/common/proto/data"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DataService struct {
	client pb.DataClient
}

func (d DataService) CreateTempInfo(ctx context.Context, name string, size int64) (uuid string, err error) {
	resp, err := d.client.CreateTempInfo(ctx, &pb.CreateTempInfoReq{
		Name: name,
		Size: size,
	})

	if err != nil {
		return
	}

	uuid = resp.Uuid
	return
}

func (d DataService) UploadTempFile(ctx context.Context, uuid string, reader io.Reader) error {
	stream, err := d.client.UploadTempFile(ctx)
	if err != nil {
		return err
	}

	// 先将uuid传过去
	err = stream.Send(&pb.UploadTempFileReq{Data: nil, Uuid: uuid})
	if err != nil {
		return err
	}

	// 传输文件
	buf := make([]byte, 1024*1024)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Failed to read from file: %v", err)
		}
		if err := stream.Send(&pb.UploadTempFileReq{Data: buf[:n]}); err != nil {
			log.Fatalf("Error sending file chunk: %v", err)
		}
	}

	// 关闭文件数据流并接收响应
	_, err = stream.CloseAndRecv()
	if err != nil {
		return err
	}

	return nil
}

func (d DataService) CommitTempFile(ctx context.Context, uuid, hash string) error {
	_, err := d.client.CommitTempFile(ctx, &pb.CommitTempFileReq{
		Uuid: uuid,
		Hash: hash,
	})
	if err != nil {
		return err
	}

	return nil
}

func (d DataService) DeleteTempFile(ctx context.Context, uuid string) error {
	_, err := d.client.DeleteTempFile(ctx, &pb.CommonReq{Uuid: uuid})
	return err
}

func (d DataService) GetTempFile(ctx context.Context, uuid string) (io.Reader, error) {
	stream, err := d.client.GetTempFile(ctx, &pb.CommonReq{Uuid: uuid})
	if err != nil {
		return nil, err
	}

	return NewRpcGetStream(&TempFileDataRecv{stream}), nil
}

func (d DataService) HeadTempFile(ctx context.Context, uuid string) (int64, error) {
	resp, err := d.client.HeadTempFile(ctx, &pb.CommonReq{Uuid: uuid})
	if err != nil {
		return 0, err
	}

	return resp.ContentLength, nil
}

func (d DataService) GetMatter(ctx context.Context, namePrefix string) (io.Reader, error) {
	stream, err := d.client.GetMatter(ctx, &pb.GetMatterReq{
		NamePrefix: namePrefix,
	})
	if err != nil {
		return nil, err
	}

	return NewRpcGetStream(FileDataRecv{stream}), nil
}

func (d DataService) CheckTempFileHash(ctx context.Context, uuid, hash string, offset int64) error {
	_, err := d.client.CheckTempFileHash(ctx, &pb.CheckTempFileHashReq{
		Uuid:   uuid,
		Hash:   hash,
		Offset: offset,
	})

	return err
}

func GetDataClient(addr string) (*DataService, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}

	return &DataService{pb.NewDataClient(conn)}, nil
}
