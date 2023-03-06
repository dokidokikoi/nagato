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

type UploadMatter struct {
	Name    string `json:"name" binding:"required"`
	Sha256  string `json:"sha256" binding:"required"`
	Size    uint   `json:"size" binding:"required"`
	Privacy bool   `json:"privacy"`
	Path    string `json:"path"`
}

type UpdateMatter struct {
	Name    string `json:"name" binding:"required"`
	Privacy bool   `json:"privacy"  binding:"required"`
	Path    string `json:"path"  binding:"required"`
	Ext     string `json:"ext"  binding:"required"`
}
