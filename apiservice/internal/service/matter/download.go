package matter

import (
	"context"
	"fmt"
	"nagato/apiservice/internal/heartbeat"
	"nagato/apiservice/internal/locate"
	"nagato/apiservice/stream"
)

func (s matterSrv) Download(ctx context.Context, hash string, size uint) (*stream.RSGetStream, error) {
	// 文件是否存在于数据服务
	locateInfo := locate.Locate(hash)
	if len(locateInfo) < stream.DATA_SHARDS {
		return nil, fmt.Errorf("文件 %s 定位失败, result %+v", hash, locateInfo)
	}
	dataServers := make([]string, 0)

	// locateinfo 的数量小于 ALL_SHARDS， 需要对缺失的分片修正
	if len(locateInfo) < stream.ALL_SHARDS {
		dataServers = heartbeat.ChooseRandomDataServers(stream.ALL_SHARDS-len(locateInfo), locateInfo)
	}

	return stream.NewRSGetStream(locateInfo, dataServers, hash, size)
}
