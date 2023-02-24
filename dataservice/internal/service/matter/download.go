package matter

import (
	"context"
	"fmt"
	"nagato/common/tools"
	"nagato/dataservice/internal/config"
	"nagato/dataservice/internal/locate"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
)

func (s matterSrv) GetFilePath(ctx context.Context, name string) (string, error) {
	files, _ := filepath.Glob(config.Config().FileSystemConfig.StoreDir + name + ".*")
	if len(files) != 1 {
		return "", fmt.Errorf("数据丢失: %s", name)
	}
	file := files[0]
	fileHash := strings.Split(file, ".")[2]

	f, _ := os.Open(file)
	h, err := tools.CalculateHash(f)
	if err != nil {
		return "", fmt.Errorf("计算hash错误: %s", err.Error())
	}
	f.Close()

	h = url.PathEscape(h)
	// 数据校验，因为可能由于硬件上的问题导致数据出错，例如数据降解
	if h != fileHash {
		zaplog.L().Sugar().Error("数据丢失: %s", name)
		locate.Del(fileHash)
		os.Remove(file)
		return "", fmt.Errorf("数据丢失: %s", name)
	}

	return file, nil
}
