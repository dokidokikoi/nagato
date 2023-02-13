package main

import (
	"nagato/common/es"
	pb "nagato/common/proto/search"
	"nagato/searchservice/rpc/server"
	"os"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	Name = "nagato.search"
)

func main() {
	cert, _ := os.ReadFile("http_ca.crt")
	cli, err := es.NewEsClient("elastic", "ADR*piFezssmbUhhN8*S", cert, "https://127.0.0.1:9200")
	if err != nil {
		panic(err)
	}

	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}

	s := server.NewSearchService(cli)
	grpcSrv := grpc.NewServer(
		grpc.Address(":10300"),
		grpc.Middleware(
			recovery.Recovery(),
		),
	)
	pb.RegisterSearchServer(grpcSrv, s)

	r := etcd.New(client)
	app := kratos.New(
		kratos.Name(Name),
		kratos.Server(
			grpcSrv,
		),
		kratos.Registrar(r),
	)

	app.Run()
}
