package server

import (
	"context"
	"nagato/apiservice/internal/config"
	"nagato/common/service"

	pb "nagato/common/proto/search"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	SearchSrv *SearchService
)

type SearchService struct {
	client pb.SearchClient
}

func (s SearchService) CreateDocByID(ctx context.Context, index, id, body string) error {
	req := pb.DocReqest{
		Index: index,
		Id:    id,
		Body:  body,
	}

	_, err := s.client.CreateDocByID(ctx, &req)
	return err
}

func (s SearchService) GetDoc(ctx context.Context, index, id string) ([]byte, error) {
	req := pb.DocReqest{
		Index: index,
		Id:    id,
	}

	doc, err := s.client.GetDoc(ctx, &req)
	if err != nil {
		return nil, err
	}
	return doc.Doc, nil
}

func (s *SearchService) UpdateDoc(ctx context.Context, index, id, body string) error {
	req := &pb.DocReqest{
		Index: index,
		Id:    id,
		Body:  body,
	}
	_, err := s.client.UpdateDoc(ctx, req)
	return err
}

func (s *SearchService) DelDoc(ctx context.Context, index, id string) error {
	req := &pb.DocReqest{
		Index: index,
		Id:    id,
	}
	_, err := s.client.DelDoc(ctx, req)
	return err
}

func (s *SearchService) BulkDoc(ctx context.Context, index, body string) error {
	req := &pb.DocReqest{
		Index: index,
		Body:  body,
	}
	_, err := s.client.BulkDoc(ctx, req)
	return err
}

func (s *SearchService) CreateIndex(ctx context.Context, index, body string) error {
	req := &pb.IndexReqest{
		Index: index,
		Body:  body,
	}
	_, err := s.client.CreateIndex(ctx, req)
	return err
}

func (s *SearchService) DelIndices(ctx context.Context, indices ...string) error {
	req := &pb.DelIndexReqest{
		Indices: indices,
	}
	_, err := s.client.DelIndices(ctx, req)
	return err
}

func (s *SearchService) SearchDoc(ctx context.Context, index, body string) ([]byte, error) {
	req := &pb.SearchReqest{
		Index: index,
		Body:  body,
	}
	resp, err := s.client.SearchDoc(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Doc, nil
}

func init() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{config.Config().EtcdConfig.Address()},
	})
	if err != nil {
		panic(err)
	}

	conn, err := service.GetRpcConn(cli, config.Config().ServiceConfig[0])
	if err != nil {
		panic(err)
	}

	client := pb.NewSearchClient(conn)

	SearchSrv = &SearchService{client: client}
}
