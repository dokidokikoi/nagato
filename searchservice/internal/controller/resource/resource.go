package resource

import "nagato/searchservice/internal/service"

type Controller struct {
	service service.IService
}

func NewController() Controller {
	return Controller{service: service.NewSrv()}
}
