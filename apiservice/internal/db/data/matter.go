package data

import (
	"context"
	"nagato/apiservice/internal/db/data/postgres"
	"nagato/apiservice/internal/db/data/redis"
	"nagato/apiservice/internal/model"
	"time"

	meta "github.com/dokidokikoi/go-common/meta/option"
	"gorm.io/gorm"
)

type matters struct {
	pg    *postgres.Store
	redis *redis.Store
}

func (m matters) Create(ctx context.Context, example *model.Matter, option *meta.CreateOption) error {
	return m.pg.Matters().Create(ctx, example, option)
}

func (m matters) CreateMany2Many(ctx context.Context, example *model.Matter, ids interface{}, option *meta.CreateOption) error {
	return m.pg.Matters().CreateMany2Many(ctx, example, ids, option)
}

func (m matters) CreateCollection(ctx context.Context, examples []*model.Matter, option *meta.CreateCollectionOption) []error {
	return m.pg.Matters().CreateCollection(ctx, examples, option)
}

func (m matters) Update(ctx context.Context, example *model.Matter, option *meta.UpdateOption) error {
	return m.pg.Matters().Update(ctx, example, option)
}

func (m matters) UpdateByWhere(ctx context.Context, node *meta.WhereNode, example *model.Matter, option *meta.UpdateOption) error {
	return m.pg.Matters().UpdateByWhere(ctx, node, example, option)
}

func (m matters) UpdateCollection(ctx context.Context, examples []*model.Matter, option *meta.UpdateCollectionOption) []error {
	return m.pg.Matters().UpdateCollection(ctx, examples, option)
}

func (m matters) Save(ctx context.Context, example *model.Matter, option *meta.UpdateOption) error {
	return m.pg.Matters().Save(ctx, example, option)
}

func (m matters) Get(ctx context.Context, example *model.Matter, option *meta.GetOption) (*model.Matter, error) {
	return m.pg.Matters().Get(ctx, example, option)
}

func (m matters) Count(ctx context.Context, example *model.Matter, option *meta.GetOption) (int64, error) {
	return m.pg.Matters().Count(ctx, example, option)
}

func (m matters) CountComplex(ctx context.Context, example *model.Matter, condition *meta.WhereNode, option *meta.GetOption) (int64, error) {
	return m.pg.Matters().CountComplex(ctx, example, condition, option)
}

func (m matters) List(ctx context.Context, example *model.Matter, option *meta.ListOption) ([]*model.Matter, error) {
	return m.pg.Matters().List(ctx, example, option)
}

func (m matters) ListComplex(ctx context.Context, example *model.Matter, condition *meta.WhereNode, option *meta.ListOption) ([]*model.Matter, error) {
	return m.pg.Matters().ListComplex(ctx, example, condition, option)
}

func (m matters) Delete(ctx context.Context, example *model.Matter, option *meta.DeleteOption) error {
	return m.pg.Matters().Delete(ctx, example, option)
}

func (m matters) DeleteCollection(ctx context.Context, examples []*model.Matter, option *meta.DeleteCollectionOption) []error {
	return m.pg.Matters().DeleteCollection(ctx, examples, option)
}

func (m matters) DeleteByIds(ctx context.Context, ids []uint) error {
	return m.pg.Matters().DeleteByIds(ctx, ids)
}

func (m matters) CountDB(ctx context.Context, example *model.Matter, option *meta.GetOption) *gorm.DB {
	return m.pg.Matters().CountDB(ctx, example, option)
}

func (m matters) CountComplexDB(ctx context.Context, example *model.Matter, condition *meta.WhereNode, option *meta.GetOption) *gorm.DB {
	return m.pg.Matters().CountComplexDB(ctx, example, condition, option)
}

func (m matters) ListDB(ctx context.Context, example *model.Matter, option *meta.ListOption) *gorm.DB {
	return m.pg.Matters().ListDB(ctx, example, option)
}

func (m matters) ListComplexDB(ctx context.Context, example *model.Matter, condition *meta.WhereNode, option *meta.ListOption) *gorm.DB {
	return m.pg.Matters().ListComplexDB(ctx, example, condition, option)
}

func (m matters) Insert(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return m.redis.Matters().Insert(ctx, key, value, expiration)
}

func (m matters) DelCache(ctx context.Context, key string) error {
	return m.redis.Matters().Del(ctx, key)
}

func (m matters) GetCache(ctx context.Context, key string) (string, error) {
	return m.redis.Matters().Get(ctx, key)
}

func newMatters(center *dataCenter) *matters {
	return &matters{pg: center.pg, redis: center.redis}
}
