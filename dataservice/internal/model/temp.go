package model

import (
	"encoding/json"
	"io"
	"nagato/dataservice/internal/config"
	"os"
	"strconv"
	"strings"
)

type TempInfo struct {
	Uuid string
	Name string
	Size int64
}

func (t *TempInfo) WriteToFile() error {
	f, e := os.Create(config.Config().FileSystemConfig.TempDir + t.Uuid)
	if e != nil {
		return e
	}
	defer f.Close()

	b, _ := json.Marshal(t)
	f.Write(b)
	return nil
}

func (t *TempInfo) Hash() string {
	s := strings.Split(t.Name, ".")
	return s[0]
}

func (t *TempInfo) ID() int {
	s := strings.Split(t.Name, ".")
	id, _ := strconv.Atoi(s[1])
	return id
}

// 读取临时文件的信息文件
func ReadFromTempFile(uuid string) (*TempInfo, error) {
	f, err := os.Open(config.Config().FileSystemConfig.TempDir + uuid)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, _ := io.ReadAll(f)
	var info TempInfo
	json.Unmarshal(b, &info)

	return &info, nil
}
