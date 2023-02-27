package server

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	pb "nagato/common/proto/data"
	"nagato/dataservice/internal/config"
	"nagato/dataservice/internal/model"
	"nagato/dataservice/internal/service"
	"net/url"
	"os"
	"os/exec"
	"strings"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"go.uber.org/zap"
)

type DataService struct {
	pb.UnimplementedDataServer
	service service.IService
}

func (d *DataService) CreateTempInfo(ctx context.Context, req *pb.CreateTempInfoReq) (resp *pb.CreateTempInfoResp, err error) {
	resp = &pb.CreateTempInfoResp{}
	hashEncode := url.PathEscape(req.Name)

	output, err := exec.Command("uuidgen").Output()
	if err != nil {
		return
	}
	uuid := strings.TrimSuffix(string(output), "\n")
	if err = d.service.Matter().CreateTempFile(ctx, hashEncode, uuid, req.Size); err != nil {
		zaplog.L().Error("创建临时文件信息出错", zap.Error(err))
		return
	}

	fmt.Println(uuid)
	resp.Uuid = uuid
	return
}

func (d *DataService) UploadTempFile(stream pb.Data_UploadTempFileServer) error {
	uuid := ""
	var tempInfo *model.TempInfo
	var tempFile *os.File
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.UploadTempFileResp{
				Status: 0,
			})
		}
		if err != nil {
			return err
		}

		if uuid == "" {
			uuid = req.GetUuid()
			tempInfo, err = model.ReadFromTempFile(uuid)
			if err != nil {
				return err
			}
			tempFile, err = openTempInfoFile(uuid)
			if err != nil {
				return err
			}
			defer tempFile.Close()
		} else {
			data := req.GetData()
			_, err = io.Copy(tempFile, bytes.NewBuffer(data))
			if err != nil {
				return err
			}

			info, err := tempFile.Stat()
			if err != nil {
				return err
			}
			actual := info.Size()
			if actual > tempInfo.Size {
				removeTempFile(uuid)
				return errors.New("文件大小不匹配")
			}
		}
	}
}

func (d *DataService) CommitTempFile(ctx context.Context, req *pb.CommitTempFileReq) (resp *pb.CommitTempFileResp, err error) {
	resp = &pb.CommitTempFileResp{}
	err = d.service.Matter().CommitMatter(ctx, req.Uuid, req.Hash)
	if err != nil {
		resp.Status = -1
		zaplog.L().Error("转正临时文件出错", zap.Error(err))
	}
	return
}

func (d *DataService) DeleteTempFile(ctx context.Context, req *pb.CommonReq) (resp *pb.DeleteTempFileResp, err error) {
	d.service.Matter().DelMatterTemp(ctx, req.Uuid)
	return
}

func (d *DataService) GetTempFile(req *pb.CommonReq, stream pb.Data_GetTempFileServer) error {
	f, err := os.Open(config.Config().FileSystemConfig.TempDir + req.Uuid + ".dat")
	if err != nil {
		zaplog.L().Error("打开文件错误", zap.Error(err))
		return err
	}
	defer f.Close()

	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if err := stream.Send(&pb.GetTempFileResp{Data: buf[:n]}); err != nil {
			return err
		}
	}

	return nil
}

func (d *DataService) HeadTempFile(ctx context.Context, req *pb.CommonReq) (resp *pb.HeadTempFileResp, err error) {
	resp = &pb.HeadTempFileResp{}
	f, err := os.Open(config.Config().FileSystemConfig.TempDir + req.Uuid + ".dat")
	if err != nil {
		zaplog.L().Error("打开文件错误", zap.Error(err))
		resp.Status = -1
		return
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		zaplog.L().Error("获取文件信息错误", zap.Error(err))
		resp.Status = -1
		return
	}

	resp.ContentLength = info.Size()
	return
}

func (d *DataService) GetMatter(req *pb.GetMatterReq, stream pb.Data_GetMatterServer) error {
	if req.NamePrefix == "" {
		zaplog.L().Error("hash不能为空")
		return errors.New("hash不能为空")
	}

	path, err := d.service.Matter().GetFilePath(context.Background(), req.NamePrefix)
	if err != nil {
		zaplog.L().Sugar().Errorf("获取文件失败, err: %+v", err)
		return err
	}
	f, err := os.Open(path)
	if err != nil {
		zaplog.L().Sugar().Errorf("打开文件失败, err: %+v", err)
		return err
	}
	defer f.Close()

	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if err := stream.Send(&pb.GetMatterResp{Data: buf[:n]}); err != nil {
			return err
		}
	}

	return nil
}

func NewDataService() *DataService {
	return &DataService{
		service: service.NewService(),
	}
}

func openTempInfoFile(uuid string) (*os.File, error) {
	infoFile := config.Config().FileSystemConfig.TempDir + uuid
	datFile := infoFile + ".dat"
	f, err := os.OpenFile(datFile, os.O_WRONLY|os.O_APPEND, 0)
	if err != nil {
		return nil, err
	}

	return f, err
}

func removeTempFile(uuid string) {
	infoFile := config.Config().FileSystemConfig.TempDir + uuid
	datFile := infoFile + ".dat"
	os.Remove(datFile)
	os.Remove(infoFile)
}
