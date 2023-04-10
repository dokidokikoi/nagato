package share

import (
	"context"
	"nagato/apiservice/internal/db"
	"nagato/apiservice/internal/model"
	"time"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type IShareService interface {
	Create(ctx context.Context, example *model.Share) error
	Update(ctx context.Context, example *model.Share) error
	Save(ctx context.Context, example *model.Share, option *meta.UpdateOption) error
	Del(ctx context.Context, example *model.Share) error
	Get(ctx context.Context, example *model.Share, option *meta.GetOption) (*model.Share, error)

	Receive(src []*model.Matter, target []uint) ([]*model.Matter, error)
	receiveMap(src []*model.Matter, m map[uint]struct{}) ([]*model.Matter, error)
	IsExpired(ctx context.Context, example *model.Share) bool
}

type shareSrv struct {
	store db.Store
}

func (s shareSrv) Create(ctx context.Context, example *model.Share) error {
	// tx := s.store.Transaction().TransactionBegin()
	// err := tx.Shares().Create(ctx, example, &meta.CreateOption{Omit: "Matters"})
	// if err != nil {
	// 	tx.Transaction().TransactionRollback()
	// 	errs = append(errs, err)
	// }

	// shareMatters := make([]*model.ShareMatter, len(matterIDs))
	// for i := range matterIDs {
	// 	shareMatters[i] = &model.ShareMatter{
	// 		MatterID: matterIDs[i],
	// 		ShareID:  example.ID,
	// 	}
	// }
	// e := tx.ShareMatters().CreateCollection(ctx, shareMatters, nil)
	// if e != nil {
	// 	tx.Transaction().TransactionRollback()
	// 	errs = append(errs, e...)
	// }

	// tx.Transaction().TransactionCommit()

	return s.store.Shares().Create(ctx, example, &meta.CreateOption{Omit: "Matters.*"})
}

func (s shareSrv) Update(ctx context.Context, example *model.Share) error {
	return s.store.Shares().Update(ctx, example, nil)
}

func (s shareSrv) Save(ctx context.Context, example *model.Share, option *meta.UpdateOption) error {
	return s.store.Shares().Save(ctx, example, option)
}

func (s shareSrv) Del(ctx context.Context, example *model.Share) error {
	return s.store.Shares().Delete(ctx, example, nil)
}

func (s shareSrv) Get(ctx context.Context, example *model.Share, option *meta.GetOption) (*model.Share, error) {
	return s.store.Shares().Get(ctx, example, option)
}

func (s shareSrv) Receive(src []*model.Matter, target []uint) ([]*model.Matter, error) {
	m := make(map[uint]struct{}, len(target))
	for _, id := range target {
		m[id] = struct{}{}
	}

	return s.receiveMap(src, m)
}

func (s shareSrv) receiveMap(src []*model.Matter, m map[uint]struct{}) ([]*model.Matter, error) {
	res := make([]*model.Matter, 0)
	next := make([]*model.Matter, 0)
	for _, matter := range src {
		_, ok := m[matter.ID]
		if ok {
			res = append(res, matter)
			delete(m, matter.ID)
		} else {
			list, _ := s.store.Matters().List(context.Background(), &model.Matter{PUUID: matter.UUID}, nil)
			next = append(next, list...)
		}
	}
	if len(res) > 0 {
		return res, nil
	}

	var err error
	r := make([]*model.Matter, 0)
	if len(m) > 0 && len(next) > 0 {
		r, err = s.receiveMap(next, m)
		if err != nil {
			return nil, err
		}
	}

	res = append(res, r...)
	return res, nil
}

func (s shareSrv) IsExpired(ctx context.Context, example *model.Share) bool {
	return !example.ExpireInfinity && example.ExpireTime.Before(time.Now())
}

func NewShareSrv(store db.Store) IShareService {
	return &shareSrv{store: store}
}
