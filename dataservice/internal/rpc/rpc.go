package rpc

import (
	"fmt"
	"log"
	pb "nagato/common/proto/data"
	"nagato/dataservice/internal/config"
	"nagato/dataservice/internal/rpc/server"
	"net"

	"google.golang.org/grpc"
)

func Init() {

}

func Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Config().RpcConfig.Host, config.Config().RpcConfig.Port))
	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterDataServer(grpcServer, server.NewDataService())

	log.Default().Printf("Listening and serving rpc on %s:%d", config.Config().RpcConfig.Host, config.Config().RpcConfig.Port)
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
