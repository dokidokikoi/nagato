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

type blanks struct {
	esCli *es.Store
	cli   *redis.Store
}

func (b blanks) CreateIndices(userID uint, indexReq string) error {
	return b.esCli.Blanks().CreateIndex(fmt.Sprintf("blank_%d", userID), strings.NewReader(indexReq))
}

func (b blanks) SearchBlank(userID uint, req commonEsModel.BlankReq) ([]commonEs.Result[commonEsModel.Blank], error) {
	return b.esCli.Blanks().SearchBlank(fmt.Sprintf("blank_%d", userID), req)
}

func (b blanks) CreateDocWithID(userID uint, id string, req commonEsModel.Blank) error {
	return b.esCli.Blanks().CreateDocWithID(fmt.Sprintf("blank_%d", userID), id, req)
}

func newBlanks(d dataCenter) db.IBlankStore {
	return &blanks{esCli: d.esCli, cli: d.cli}
}
