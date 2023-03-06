package blank

import (
	"nagato/apiservice/internal/controller"
	"nagato/apiservice/internal/service"
)

type BlankController struct {
	controller.Controller
	service service.IService
}

func NewBlankController() *BlankController {
	return &BlankController{service: service.NewService()}
}

type CreateBlank struct {
	Type    string   `json:"type"`
	Title   string   `json:"title" binding:"required"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
	Matters []uint   `json:"matters"`
}

type UpdateBlank struct {
	Type    string   `json:"type"`
	Title   string   `json:"title" binding:"required"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
	Matters []uint   `json:"matters"`
}
