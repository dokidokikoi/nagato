package controller

import "nagato/apiservice/internal/service"

type IController interface {
	Matter() *MatterController
}

type controller struct {
	service service.IService
}

func (c controller) Matter() *MatterController {
	return newMatterController(c.service)
}

func NewController() IController {
	return &controller{service: service.NewService()}
}
