package service

import (
	"nagato/apiservice/internal/db"
	"nagato/apiservice/internal/db/data"
	"nagato/apiservice/internal/service/blank"
	"nagato/apiservice/internal/service/matter"
	"nagato/apiservice/internal/service/share"
	"nagato/apiservice/internal/service/tag"
	"nagato/apiservice/internal/service/user"
)

type IService interface {
	Matter() matter.IMatterService
	User() user.IUserService
	Blank() blank.IBlankService
	Tag() tag.ITagService
	Share() share.IShareService
}

type service struct {
	store db.Store
}

func (s service) Matter() matter.IMatterService {
	return matter.NewMatterService(s.store)
}

func (s service) User() user.IUserService {
	return user.NewUserService(s.store)
}

func (s service) Blank() blank.IBlankService {
	return blank.NewBlankSrv(s.store)
}

func (s service) Tag() tag.ITagService {
	return tag.NewTagSrv(s.store)
}

func (s service) Share() share.IShareService {
	return share.NewShareSrv(s.store)
}

func NewService() IService {
	store, err := data.GetStoreDBFactory()
	if err != nil {
		panic(err)
	}
	return &service{store: store}
}
