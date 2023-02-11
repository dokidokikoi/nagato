package service

import (
	"nagato/apiservice/internal/es"
	"nagato/apiservice/internal/service/matter"
)

type IService interface {
	Matter() matter.IMatterService
}

type service struct {
	esCli es.IEsStore
}

func (s service) Matter() matter.IMatterService {
	return matter.NewMatterService(s.esCli)
}

func NewService() IService {
	return &service{esCli: es.NewEsSore()}
}
