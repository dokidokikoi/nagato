package service

import (
	"nagato/searchservice/internal/db"
	"nagato/searchservice/internal/db/data"
)

type IService interface {
	Blank() IBlankService
	Resource() IResourceService
}

type srv struct {
	store db.Store
}

func (s srv) Blank() IBlankService {
	return newBlankSrv(s.store)
}

func (s srv) Resource() IResourceService {
	return newResourceSrv(s.store)
}

func NewSrv() IService {
	return &srv{store: data.GetDataCenter()}
}
