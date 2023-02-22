package matter

import (
	"nagato/apiservice/internal/controller"
	"nagato/apiservice/internal/service"
)

type MatterController struct {
	controller.Controller
	service service.IService
}

func NewMatterController() *MatterController {
	return &MatterController{service: service.NewService()}
}
