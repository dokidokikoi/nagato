package server

import (
	"context"
	"fmt"
	"io"
	pb "nagato/common/proto/search"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type SearchService struct {
	pb.UnimplementedSearchServer
	cli *elasticsearch.Client
}

func (s *SearchService) GetDoc(ctx context.Context, in *pb.DocReqest) (*pb.DocResponse, error) {
	resp, err := s.cli.GetSource(in.Index, in.Id)
	if err != nil {
		return &pb.DocResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return &pb.DocResponse{}, fmt.Errorf("es状态码不为200, code: %d", resp.StatusCode)
	}

	result, _ := io.ReadAll(resp.Body)

	res := &pb.DocResponse{
		Doc: result,
	}

	return res, nil
}

func (s *SearchService) CreateDocByID(ctx context.Context, in *pb.DocReqest) (*pb.Response, error) {
	_, err := s.cli.Create(in.Index, in.Id, strings.NewReader(in.Body))
	return &pb.Response{}, err
}

func (s *SearchService) UpdateDoc(ctx context.Context, in *pb.DocReqest) (*pb.Response, error) {
	_, err := s.cli.Index(in.Index, strings.NewReader(in.Body), s.cli.Index.WithDocumentID(in.Id))
	return &pb.Response{}, err
}

func (s *SearchService) DelDoc(ctx context.Context, in *pb.DocReqest) (*pb.Response, error) {
	_, err := s.cli.Get(in.Index, in.Id)
	return &pb.Response{}, err
}

func (s *SearchService) BulkDoc(ctx context.Context, in *pb.DocReqest) (*pb.Response, error) {
	_, err := s.cli.Bulk(strings.NewReader(in.Body), s.cli.Bulk.WithIndex(in.Index))
	return &pb.Response{}, err
}

func (s *SearchService) CreateIndex(ctx context.Context, in *pb.IndexReqest) (*pb.Response, error) {
	_, err := s.cli.Indices.Create(in.Index, s.cli.Indices.Create.WithBody(strings.NewReader(in.Body)))
	return &pb.Response{}, err
}

func (s *SearchService) DelIndices(ctx context.Context, in *pb.DelIndexReqest) (*pb.Response, error) {
	_, err := s.cli.Indices.Delete(in.Indices)
	return &pb.Response{}, err
}

func (s *SearchService) SearchDoc(ctx context.Context, in *pb.SearchReqest) (*pb.DocResponse, error) {
	resp, err := s.cli.Search(s.cli.Search.WithIndex(in.Index), s.cli.Search.WithBody(strings.NewReader(in.Body)))
	if err != nil {
		return &pb.DocResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return &pb.DocResponse{}, fmt.Errorf("es状态码不为200, code: %d", resp.StatusCode)
	}

	result, _ := io.ReadAll(resp.Body)

	res := &pb.DocResponse{
		Doc: result,
	}

	return res, nil
}

func NewSearchService(cli *elasticsearch.Client) *SearchService {
	return &SearchService{cli: cli}
}
