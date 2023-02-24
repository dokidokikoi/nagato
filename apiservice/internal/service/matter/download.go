package matter

import (
	"context"
	"fmt"
	"nagato/apiservice/internal/heartbeat"
	"nagato/apiservice/internal/locate"
	"nagato/apiservice/stream"
)

func (s matterSrv) Download(ctx context.Context, hash string, size uint) (*stream.RSGetStream, error) {
	locateInfo := locate.Locate(hash)
	if len(locateInfo) < stream.DATA_SHARDS {
		return nil, fmt.Errorf("object %s locate failed, result %+v", hash, locateInfo)
	}
	dataServers := make([]string, 0)

	// locateinfo 的数量小于 ALL_SHARDS， 需要对缺失的分片修正
	if len(locateInfo) != stream.ALL_SHARDS {
		dataServers = heartbeat.ChooseRandomDataServers(stream.ALL_SHARDS-len(locateInfo), locateInfo)
	}
	// if server == "" {
	// 	return nil, fmt.Errorf("matter %s locate failed", hash)
	// }

	return stream.NewRSGetStream(locateInfo, dataServers, hash, size)
}
