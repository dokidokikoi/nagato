package service

import (
	"nagato/apiservice/internal/db"
	"nagato/apiservice/internal/es"
	"nagato/apiservice/internal/service/matter"
)

type IService interface {
	Matter() matter.IMatterService
}

type service struct {
	esCli es.IEsStore
	store db.Store
}

func (s service) Matter() matter.IMatterService {
	return matter.NewMatterService(s.esCli, s.store)
}

func NewService() IService {
	store, err := db.GetStoreFactory()
	if err != nil {
		panic(err)
	}
	return &service{esCli: es.NewEsSore(), store: store}
}
