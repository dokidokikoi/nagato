package service

import "nagato/dataservice/internal/service/matter"

type IService interface {
	Matter() matter.IMatterService
}

type service struct {
}

func (s service) Matter() matter.IMatterService {
	return matter.NewMatterService()
}

func NewService() IService {
	return &service{}
}
