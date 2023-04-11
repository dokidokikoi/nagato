package app

import (
	"fmt"
	pb "nagato/common/proto/search"
	inittask "nagato/searchservice/init"
	search "nagato/searchservice/internal"
	"nagato/searchservice/internal/config"
	"nagato/searchservice/internal/db/data/es"
	"nagato/searchservice/rpc/server"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
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

	router := gin.Default()
	search.InitRoute(router)
	httpSrv := http.NewServer(http.Address(fmt.Sprintf("%s:%d", config.Config().ServerConfig.Host, config.Config().ServerConfig.Port)))
	httpSrv.HandlePrefix("", router)

	r := etcd.New(client)
	app := kratos.New(
		kratos.Name(config.Config().RpcConfig.Name),
		kratos.Server(
			grpcSrv,
			httpSrv,
		),
		kratos.Registrar(r),
	)

	app.Run()
}
