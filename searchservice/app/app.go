package app

import (
	pb "nagato/common/proto/search"
	inittask "nagato/searchservice/init"
	"nagato/searchservice/internal/config"
	"nagato/searchservice/internal/db/data/es"
	"nagato/searchservice/rpc/server"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type App struct {
}

func (a App) Run() {
	inittask.Init()
	cli, err := es.GetEsClient()
	if err != nil {
		panic(err)
	}
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{config.Config().EtcdConfig.Address()},
	})
	if err != nil {
		panic(err)
	}

	s := server.NewSearchService(cli)
	grpcSrv := grpc.NewServer(
		grpc.Address(config.Config().RpcConfig.Address()),
		grpc.Middleware(
			recovery.Recovery(),
		),
	)
	pb.RegisterSearchServer(grpcSrv, s)

	r := etcd.New(client)
	app := kratos.New(
		kratos.Name(config.Config().RpcConfig.Name),
		kratos.Server(
			grpcSrv,
		),
		kratos.Registrar(r),
	)

	app.Run()
}
