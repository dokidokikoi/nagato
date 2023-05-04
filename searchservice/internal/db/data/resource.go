package data

import (
	"fmt"
	commonEs "nagato/common/es"
	commonEsModel "nagato/common/es/model"
	"nagato/searchservice/internal/db"
	"nagato/searchservice/internal/db/data/es"
	"nagato/searchservice/internal/db/data/redis"
	"strings"
)

type resources struct {
	esCli *es.Store
	cli   *redis.Store
}

func (r resources) CreateIndices(userID uint, indexReq string) error {
	return r.esCli.Resources().CreateIndex(fmt.Sprintf("resource_%d", userID), strings.NewReader(indexReq))
}

func (r resources) SearchResource(userID uint, req commonEsModel.ResourceReq) ([]commonEs.Result[commonEsModel.Resource], error) {
	return r.esCli.Resources().Search(fmt.Sprintf("resource_%d", userID), req)
}

func (r resources) CreateDocWithID(userID uint, id string, req commonEsModel.Resource) error {
	return r.esCli.Resources().CreateDocWithID(fmt.Sprintf("resource_%d", userID), id, req)
}

func newResources(d dataCenter) db.IResourceStore {
	return &resources{esCli: d.esCli, cli: d.cli}
}
