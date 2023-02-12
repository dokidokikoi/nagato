package matter

import (
	"context"
	"nagato/common/tools"
	"nagato/dataservice/internal/config"
	"nagato/dataservice/internal/locate"
	"os"
)

func (s matterSrv) GetFilePath(ctx context.Context, hash string) string {
	file := config.Config().FileSystemConfig.StoreDir + hash
	f, _ := os.Open(file)
	h, err := tools.CalculateHash(f)
	if err != nil {
		return ""
	}
	f.Close()

	// 数据校验，因为可能由于硬件上的问题导致数据出错，例如数据降解
	if h != hash {
		locate.Del(hash)
		os.Remove(file)
		return ""
	}

	return file
}
