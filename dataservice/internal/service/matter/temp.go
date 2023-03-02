package matter

import (
	"context"
	"errors"
	"fmt"
	"io"
	"nagato/common/tools"
	"nagato/dataservice/internal/config"
	"nagato/dataservice/internal/locate"
	"nagato/dataservice/internal/model"
	"net/url"
	"os"
	"time"
)

// 将临时文件的信息存储到临时目录并创建出临时文件的文件
func (s matterSrv) CreateTempFile(ctx context.Context, hashEncode string, uuid string, size int64) error {
	info := model.TempInfo{
		Uuid: uuid,
		Name: hashEncode,
		Size: size,
	}

	err := info.WriteToFile()
	if err != nil {
		return err
	}

	// 创建出临时文件的文件
	_, err = os.Create(config.Config().FileSystemConfig.TempDir + info.Uuid + ".dat")
	if err != nil {
		return err
	}
	go func() {
		time.Sleep(5 * 24 * time.Hour)
		os.Remove(config.Config().FileSystemConfig.TempDir + info.Uuid + ".dat")
		os.Remove(config.Config().FileSystemConfig.TempDir + info.Uuid)
	}()
	return nil
}

func (s matterSrv) WriteTempFile(ctx context.Context, uuid string, data io.Reader) error {
	tempInfo, err := model.ReadFromTempFile(uuid)
	if err != nil {
		return err
	}

	infoFile := config.Config().FileSystemConfig.TempDir + uuid
	datFile := infoFile + ".dat"
	f, err := os.OpenFile(datFile, os.O_WRONLY|os.O_APPEND, 0)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, data)
	if err != nil {
		return err
	}

	info, err := f.Stat()
	if err != nil {
		return nil
	}
	actual := info.Size()
	if actual > tempInfo.Size {
		os.Remove(datFile)
		os.Remove(infoFile)
		return errors.New("文件大小不匹配")
	}
	return nil
}

func (s matterSrv) CheckTempFileHash(ctx context.Context, uuid, hash string, offset int64) error {
	infoFile := config.Config().FileSystemConfig.TempDir + uuid
	datFile := infoFile + ".dat"
	f, err := os.Open(datFile)
	if err != nil {
		return err
	}

	cacheSize := 1 << 20
	cnt := offset / int64(cacheSize)
	buf := make([]byte, cacheSize)
	for ; cnt > 0; cnt-- {
		_, err := io.ReadFull(f, buf)
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			return err
		}
	}
	buf = buf[:int(offset%int64(cacheSize))]
	_, err = io.ReadFull(f, buf)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return err
	}

	d, err := tools.CalculateHash(f)
	if err != nil {
		os.Remove(infoFile)
		os.Remove(datFile)
		return err
	}

	if d != hash {
		os.Remove(infoFile)
		os.Remove(datFile)
		return fmt.Errorf("hash 不匹配,希望的hash: %s, 文件实际hash: %s", hash, d)
	}

	return nil
}

func (s matterSrv) CommitMatter(ctx context.Context, uuid, hash string) error {
	tempInfo, err := model.ReadFromTempFile(uuid)
	if err != nil {
		return err
	}

	infoFile := config.Config().FileSystemConfig.TempDir + uuid
	datFile := infoFile + ".dat"
	f, err := os.Open(datFile)
	if err != nil {
		os.Remove(infoFile)
		os.Remove(datFile)
		return err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		os.Remove(infoFile)
		os.Remove(datFile)
		return err
	}

	actual := info.Size()
	// 将临时文件的信息文件删除
	os.Remove(infoFile)

	if actual != tempInfo.Size {
		os.Remove(datFile)
		return errors.New("文件大小不匹配")
	}

	d, err := tools.CalculateHash(f)
	if err != nil {
		os.Remove(datFile)
		return err
	}

	// if d != hash {
	// 	os.Remove(datFile)
	// 	return fmt.Errorf("hash 不匹配,希望的hash: %s, 文件实际hash: %s", hash, d)
	// }

	if err := os.Rename(datFile, config.Config().FileSystemConfig.StoreDir+tempInfo.Name+"."+url.PathEscape(d)); err != nil {
		os.Remove(datFile)
		return err
	}
	locate.Add(tempInfo.Hash(), tempInfo.ID())

	fmt.Printf("fileName: %s, storeName: %s", datFile, config.Config().FileSystemConfig.StoreDir+tempInfo.Name+"."+url.PathEscape(d))

	return nil
}

// 删除临时文件
func (c matterSrv) DelMatterTemp(ctx context.Context, uuid string) {
	infoFile := config.Config().FileSystemConfig.TempDir + uuid
	datFile := infoFile + ".dat"
	os.Remove(infoFile)
	os.Remove(datFile)
}
