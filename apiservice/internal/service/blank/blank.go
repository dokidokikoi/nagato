package blank

import (
	"bytes"
	"context"
	"encoding/json"
	"nagato/apiservice/internal/db"
	"nagato/apiservice/internal/model"
	commonEsModel "nagato/common/es/model"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type IBlankService interface {
	Create(ctx context.Context, example *model.Blank) error
	List(ctx context.Context, example *model.Blank, option *meta.ListOption) ([]*model.Blank, int64, error)
	Update(ctx context.Context, example *model.Blank) error
	UpdateBlank(ctx context.Context, example *model.Blank) error
	Save(ctx context.Context, example *model.Blank) error
	Del(ctx context.Context, example *model.Blank) error
	Get(ctx context.Context, example *model.Blank, option *meta.GetOption) (*model.Blank, error)

	Search(ctx context.Context, UserID uint, blankReq commonEsModel.BlankReq, resourceReq commonEsModel.ResourceReq) ([]commonEsModel.Blank, int64, error)
	CreateIndices(userID uint, indexReq string) error
	CreateDocWithID(userID uint, id string, req commonEsModel.Blank) error
	UpdateDoc(userID uint, id string, req commonEsModel.Blank) error
	DelDoc(userID uint, id string) error
}

type blankSrv struct {
	store db.Store
}

func (b blankSrv) Create(ctx context.Context, example *model.Blank) error {
	return b.store.Blanks().Create(ctx, example, nil)
}

func (b blankSrv) List(ctx context.Context, example *model.Blank, option *meta.ListOption) ([]*model.Blank, int64, error) {
	res, err := b.store.Blanks().List(ctx, example, option)
	if err != nil {
		return nil, 0, err
	}
	total, _ := b.store.Blanks().Count(ctx, example, &option.GetOption)
	return res, total, nil
}

func (b blankSrv) Update(ctx context.Context, example *model.Blank) error {
	return b.store.Blanks().Update(ctx, example, nil)
}

func (b blankSrv) UpdateBlank(ctx context.Context, example *model.Blank) error {
	b.store.BlankMatters().Delete(ctx, &model.BlankMatter{BlankID: example.ID}, nil)
	return b.store.Blanks().Update(ctx, example, nil)
}

func (b blankSrv) Save(ctx context.Context, example *model.Blank) error {
	return b.store.Blanks().Save(ctx, example, nil)
}

func (b blankSrv) Del(ctx context.Context, example *model.Blank) error {
	return b.store.Blanks().Delete(ctx, example, nil)
}

func (b blankSrv) Get(ctx context.Context, example *model.Blank, option *meta.GetOption) (*model.Blank, error) {
	return b.store.Blanks().Get(ctx, example, option)
}

func (b blankSrv) Search(ctx context.Context, UserID uint, blankReq commonEsModel.BlankReq, resourceReq commonEsModel.ResourceReq) ([]commonEsModel.Blank, int64, error) {
	resourceReq.Select = []string{"id"}
	resourceReq.Text = blankReq.Text
	resourceReq.Page = 0
	resourceReq.PageSize = 1000
	res, err := b.store.Matters().SearchResource(UserID, resourceReq)
	if err != nil {
		return nil, 0, err
	}

	ids := make([]uint, len(res))
	for i := range res {
		ids[i] = res[i].Source.ID
	}

	blankReq.MatterIDs = ids
	result, err := b.store.Blanks().SearchBlank(0, blankReq)
	if err != nil {
		return nil, 0, err
	}

	blanks := make([]commonEsModel.Blank, len(result))
	for i := range result {
		blanks[i] = result[i].Source
	}

	return blanks, 0, nil
}

func (b blankSrv) CreateIndices(userID uint, indexReq string) error {
	return b.store.Blanks().CreateIndices(userID, indexReq)
}

func (b blankSrv) CreateDocWithID(userID uint, id string, req commonEsModel.Blank) error {
	return b.store.Blanks().CreateDocWithID(userID, id, req)
}

func (b blankSrv) UpdateDoc(userID uint, id string, req commonEsModel.Blank) error {
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return b.store.Blanks().UpdateDoc(userID, id, bytes.NewBuffer(data))
}

func (b blankSrv) DelDoc(userID uint, id string) error {
	return b.store.Blanks().DelDoc(userID, id)
}

func NewBlankSrv(store db.Store) IBlankService {
	return &blankSrv{store: store}
}
