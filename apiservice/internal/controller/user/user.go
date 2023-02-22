package user

import (
	"nagato/apiservice/internal/controller"
	"nagato/apiservice/internal/service"
)

type UserController struct {
	controller.Controller
	service service.IService
}

func NewUserController() *UserController {
	return &UserController{service: service.NewService()}
}
