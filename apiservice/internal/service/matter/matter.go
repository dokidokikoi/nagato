package matter

import (
	"context"
	"io"
	"nagato/apiservice/internal/db"
	"nagato/apiservice/internal/model"
	"nagato/apiservice/stream"
	commonEsModel "nagato/common/es/model"
	"path/filepath"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type IMatterService interface {
	Upload(ctx context.Context, example *model.Matter, data io.Reader) error
	Download(ctx context.Context, hash string, size uint) (*stream.RSGetStream, error)
	GenUploadToken(ctx context.Context, example *model.Matter) (string, error)
	UploadBigMatter(ctx context.Context, token string, offset uint, data io.Reader) error

	Create(ctx context.Context, example *model.Matter) error
	CreateCollection(ctx context.Context, examples []*model.Matter) []error
	Update(ctx context.Context, example *model.Matter) error
	Del(ctx context.Context, example *model.Matter, option *meta.DeleteOption) error
	UpdateByWhereNode(ctx context.Context, node *meta.WhereNode, example *model.Matter) error
	Get(ctx context.Context, example *model.Matter, option *meta.GetOption) (*model.Matter, error)
	List(ctx context.Context, example *model.Matter, option *meta.ListOption) ([]*model.Matter, int64, error)
	ListMatter(ctx context.Context, example *model.Matter, option *meta.ListOption) ([]*model.Matter, error)
	ListRoot(ctx context.Context, example *model.Matter, option *meta.ListOption) ([]*model.Matter, error)

	// CreateResource(ctx context.Context, example *model.Matter) error
	// GetResourceMate(ctx context.Context, name string, version int) (*model.Matter, error)
	// SearchResourceAllVersion(ctx context.Context, name string, from, size int) ([]*model.Matter, error)
	Search(ctx context.Context, userID uint, resourceReq commonEsModel.ResourceReq) ([]commonEsModel.Resource, int64, error)
	CreateIndices(userID uint, indexReq string) error
	CreateDocWithID(userID uint, id string, req commonEsModel.Resource) error
	UpdateDoc(userID uint, id string, req commonEsModel.Resource) error
	DelDoc(userID uint, id string) error

	GetAllMatter(examples []*model.Matter) (map[uint]*model.Matter, []error)
	GetSubMatter(example *model.Matter) ([]*model.Matter, []error)
	SetMatterPath(example *model.Matter) error
	GetMatterPath(PUUID, name string, userID uint) (string, error)
}

type matterSrv struct {
	store db.Store
}

func (s matterSrv) Create(ctx context.Context, example *model.Matter) error {
	return s.store.Matters().Create(ctx, example, nil)
}

func (s matterSrv) CreateCollection(ctx context.Context, examples []*model.Matter) []error {
	return s.store.Matters().CreateCollection(ctx, examples, nil)
}

func (s matterSrv) Update(ctx context.Context, example *model.Matter) error {
	return s.store.Matters().Update(ctx, example, nil)
}

func (s matterSrv) Del(ctx context.Context, example *model.Matter, option *meta.DeleteOption) error {
	return s.store.Matters().Delete(ctx, example, option)
}

func (s matterSrv) UpdateByWhereNode(ctx context.Context, node *meta.WhereNode, example *model.Matter) error {
	return s.store.Matters().UpdateByWhere(ctx, node, example, nil)
}

func (s matterSrv) Get(ctx context.Context, example *model.Matter, option *meta.GetOption) (*model.Matter, error) {
	return s.store.Matters().Get(ctx, example, option)
}

func (s matterSrv) List(ctx context.Context, example *model.Matter, option *meta.ListOption) ([]*model.Matter, int64, error) {
	res, err := s.store.Matters().List(ctx, example, option)
	if err != nil {
		return nil, 0, err
	}
	if option == nil {
		option = &meta.ListOption{}
	}
	total, _ := s.store.Matters().Count(ctx, example, &option.GetOption)
	return res, total, nil
}

func (s matterSrv) ListMatter(ctx context.Context, example *model.Matter, option *meta.ListOption) ([]*model.Matter, error) {
	return s.store.Matters().List(ctx, example, option)
}

func (s matterSrv) ListRoot(ctx context.Context, example *model.Matter, option *meta.ListOption) ([]*model.Matter, error) {
	var res []*model.Matter
	err := s.store.Matters().ListDB(ctx, example, option).Where("p_uuid is null").Find(&res).Error
	return res, err
}

func (s matterSrv) GetAllMatter(examples []*model.Matter) (map[uint]*model.Matter, []error) {
	var errs []error
	res := make(map[uint]*model.Matter, 0)
	for _, e := range examples {
		subs, e := s.GetSubMatter(e)
		if e != nil {
			errs = append(errs, e...)
		}
		for _, m := range subs {
			res[m.ID] = m
		}
	}

	return res, errs
}

func (s matterSrv) GetSubMatter(example *model.Matter) ([]*model.Matter, []error) {
	var errs []error
	var result []*model.Matter
	res, err := s.ListMatter(context.Background(), &model.Matter{PUUID: example.UUID}, nil)
	if err != nil {
		errs = append(errs, err)
	}

	result = append(result, res...)
	for _, r := range res {
		re, e := s.GetSubMatter(r)
		if e != nil {
			errs = append(errs, e...)
		}
		result = append(result, re...)
	}

	return result, errs
}

func (s matterSrv) SetMatterPath(example *model.Matter) error {
	if example.PUUID == "" {
		example.Path = filepath.Join("/", example.Name)
		return nil
	}

	pMatter, err := s.Get(context.Background(), &model.Matter{UUID: example.PUUID, UserID: example.UserID, Dir: true}, nil)
	if err != nil {
		return err
	}

	example.Path = filepath.Join(pMatter.Path, example.Name)
	return nil
}

func (s matterSrv) GetMatterPath(PUUID, name string, userID uint) (string, error) {
	if PUUID == "" {
		return filepath.Join("/", name), nil
	}

	pMatter, err := s.Get(context.Background(), &model.Matter{UUID: PUUID, UserID: userID, Dir: true}, nil)
	if err != nil {
		return "", err
	}

	return filepath.Join(pMatter.Path, name), nil
}

func NewMatterService(store db.Store) IMatterService {
	return &matterSrv{store: store}
}
