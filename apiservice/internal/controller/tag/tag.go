package tag

import (
	"nagato/apiservice/internal/controller"
	"nagato/apiservice/internal/service"
)

type TagController struct {
	controller.Controller
	service service.IService
}

func NewTagController() *TagController {
	return &TagController{service: service.NewService()}
}

type CreateTag struct {
	TagName string `json:"tag_name" binding:"required"`
}
