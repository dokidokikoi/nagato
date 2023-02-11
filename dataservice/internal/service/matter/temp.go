package matter

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"nagato/dataservice/internal/locate"
	"os"
)

type tempInfo struct {
	Uuid string
	Name string
	Size int64
}

const STORE_ROOT = "/tmp"

func (t *tempInfo) writeToFile() error {
	f, e := os.Create(STORE_ROOT + "/temp/" + t.Uuid)
	if e != nil {
		return e
	}
	defer f.Close()

	b, _ := json.Marshal(t)
	f.Write(b)
	return nil
}

// 将临时文件的信息存储到临时目录并创建出临时文件的文件
func (s matterSrv) CreateTempFile(ctx context.Context, name string, uuid string, size int64) error {
	info := tempInfo{
		Uuid: uuid,
		Name: name,
		Size: size,
	}

	e := info.writeToFile()
	if e != nil {
		return e
	}

	// 创建出临时文件的文件
	os.Create(STORE_ROOT + "/temp/" + info.Uuid + ".dat")
	return nil
}

func (s matterSrv) WriteTempFile(ctx context.Context, uuid string, data io.Reader) error {
	tempInfo, err := readFromTempFile(uuid)
	if err != nil {
		return err
	}

	infoFile := STORE_ROOT + "/temp/" + uuid
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

func (s matterSrv) CommitMatter(ctx context.Context, uuid string) error {
	tempInfo, err := readFromTempFile(uuid)
	if err != nil {
		return err
	}

	infoFile := STORE_ROOT + "/temp/" + uuid
	datFile := infoFile + ".dat"
	f, err := os.Open(datFile)
	if err != nil {
		return err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}

	actual := info.Size()
	// 将临时文件的信息文件删除
	os.Remove(infoFile)

	if actual != tempInfo.Size {
		os.Remove(datFile)
		return errors.New("文件大小不匹配")
	}

	os.Rename(datFile, STORE_ROOT+"/objects/"+tempInfo.Name)
	locate.Add(tempInfo.Name)

	return nil
}

// 删除临时文件
func (c matterSrv) DelMatterTemp(ctx context.Context, uuid string) {
	infoFile := STORE_ROOT + "/temp/" + uuid
	datFile := infoFile + ".dat"
	os.Remove(infoFile)
	os.Remove(datFile)
}

// 读取临时文件的信息文件
func readFromTempFile(uuid string) (*tempInfo, error) {
	f, err := os.Open(STORE_ROOT + "/temp/" + uuid)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, _ := io.ReadAll(f)
	var info tempInfo
	json.Unmarshal(b, &info)

	return &info, nil
}
