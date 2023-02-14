package service

import (
	"nagato/apiservice/internal/db"
	"nagato/apiservice/internal/service/matter"
)

type IService interface {
	Matter() matter.IMatterService
}

type service struct {
	store db.Store
}

func (s service) Matter() matter.IMatterService {
	return matter.NewMatterService(s.store)
}

func NewService() IService {
	store, err := db.GetStoreFactory()
	if err != nil {
		panic(err)
	}
	return &service{store: store}
}
