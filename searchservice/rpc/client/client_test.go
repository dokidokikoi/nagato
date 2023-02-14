package client

import (
	"context"
	"encoding/json"
	"log"
	apiservice "nagato/apiservice/pkg/plugin"
	pb "nagato/common/proto/search"
	"nagato/common/service"
	"testing"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// 测试代码

func TestClist(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}

	conn, err := service.GetRpcConn(cli, "nagato.search")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewSearchClient(conn)
	reply, err := client.GetDoc(context.Background(), &pb.DocReqest{Index: "metadata", Id: "test4_3_1"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] Search %+v\n", reply)

	res := new(apiservice.Resource)
	err = json.Unmarshal(reply.Doc, res)
	if err != nil {
		panic(err)
	}
	log.Printf("[grpc] Search %+v\n", res)
}
