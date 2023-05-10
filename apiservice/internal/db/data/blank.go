package data

import (
	"context"
	"fmt"
	"io"
	"nagato/apiservice/internal/db"
	"nagato/apiservice/internal/db/data/es"
	"nagato/apiservice/internal/db/data/postgres"
	"nagato/apiservice/internal/db/data/redis"
	"nagato/apiservice/internal/model"
	commonEs "nagato/common/es"
	commonEsModel "nagato/common/es/model"
	"strings"

	meta "github.com/dokidokikoi/go-common/meta/option"
	"gorm.io/gorm"
)

type blanks struct {
	pg    *postgres.Store
	redis *redis.Store
	esCli *es.Store
	cli   *redis.Store
}

func (b blanks) Create(ctx context.Context, example *model.Blank, option *meta.CreateOption) error {
	return b.pg.Blanks().Create(ctx, example, option)
}

func (b blanks) CreateMany2Many(ctx context.Context, example *model.Blank, ids interface{}, option *meta.CreateOption) error {
	return b.pg.Blanks().CreateMany2Many(ctx, example, ids, option)
}

func (b blanks) CreateCollection(ctx context.Context, examples []*model.Blank, option *meta.CreateCollectionOption) []error {
	return b.pg.Blanks().CreateCollection(ctx, examples, option)
}

func (b blanks) Update(ctx context.Context, example *model.Blank, option *meta.UpdateOption) error {
	return b.pg.Blanks().Update(ctx, example, option)
}

func (b blanks) UpdateByWhere(ctx context.Context, node *meta.WhereNode, example *model.Blank, option *meta.UpdateOption) error {
	return b.pg.Blanks().UpdateByWhere(ctx, node, example, option)
}

func (b blanks) UpdateCollection(ctx context.Context, examples []*model.Blank, option *meta.UpdateCollectionOption) []error {
	return b.pg.Blanks().UpdateCollection(ctx, examples, option)
}

func (b blanks) Save(ctx context.Context, example *model.Blank, option *meta.UpdateOption) error {
	return b.pg.Blanks().Save(ctx, example, option)
}

func (b blanks) Get(ctx context.Context, example *model.Blank, option *meta.GetOption) (*model.Blank, error) {
	return b.pg.Blanks().Get(ctx, example, option)
}

func (b blanks) Count(ctx context.Context, example *model.Blank, option *meta.GetOption) (int64, error) {
	return b.pg.Blanks().Count(ctx, example, option)
}

func (b blanks) CountComplex(ctx context.Context, example *model.Blank, condition *meta.WhereNode, option *meta.GetOption) (int64, error) {
	return b.pg.Blanks().CountComplex(ctx, example, condition, option)
}

func (b blanks) List(ctx context.Context, example *model.Blank, option *meta.ListOption) ([]*model.Blank, error) {
	return b.pg.Blanks().List(ctx, example, option)
}

func (b blanks) ListComplex(ctx context.Context, example *model.Blank, condition *meta.WhereNode, option *meta.ListOption) ([]*model.Blank, error) {
	return b.pg.Blanks().ListComplex(ctx, example, condition, option)
}

func (b blanks) Delete(ctx context.Context, example *model.Blank, option *meta.DeleteOption) error {
	return b.pg.Blanks().Delete(ctx, example, option)
}

func (b blanks) DeleteCollection(ctx context.Context, examples []*model.Blank, option *meta.DeleteCollectionOption) []error {
	return b.pg.Blanks().DeleteCollection(ctx, examples, option)
}

func (b blanks) DeleteByIds(ctx context.Context, ids []uint) error {
	return b.pg.Blanks().DeleteByIds(ctx, ids)
}

func (b blanks) CountDB(ctx context.Context, example *model.Blank, option *meta.GetOption) *gorm.DB {
	return b.pg.Blanks().CountDB(ctx, example, option)
}

func (b blanks) CountComplexDB(ctx context.Context, example *model.Blank, condition *meta.WhereNode, option *meta.GetOption) *gorm.DB {
	return b.pg.Blanks().CountComplexDB(ctx, example, condition, option)
}

func (b blanks) ListDB(ctx context.Context, example *model.Blank, option *meta.ListOption) *gorm.DB {
	return b.pg.Blanks().ListDB(ctx, example, option)
}

func (b blanks) ListComplexDB(ctx context.Context, example *model.Blank, condition *meta.WhereNode, option *meta.ListOption) *gorm.DB {
	return b.pg.Blanks().ListComplexDB(ctx, example, condition, option)
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

func (b blanks) UpdateDoc(userID uint, id string, body io.Reader) error {
	return b.esCli.Blanks().UpdateDoc(fmt.Sprintf("blank_%d", userID), id, body)
}

func (b blanks) DelDoc(userID uint, id string) error {
	return b.esCli.Blanks().DelDoc(fmt.Sprintf("blank_%d", userID), id)
}

func newBlanks(d dataCenter) db.IBlankStore {
	return &blanks{esCli: d.esCli, cli: d.redis, pg: d.pg}
}
