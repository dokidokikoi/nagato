package service

import (
	"context"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

func GetRpcConn(client *clientv3.Client, serviceName string) (*grpc.ClientConn, error) {
	r := etcd.New(client)

	conn, err := transgrpc.DialInsecure(
		context.Background(),
		transgrpc.WithEndpoint("discovery:///"+serviceName),
		transgrpc.WithDiscovery(r),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
