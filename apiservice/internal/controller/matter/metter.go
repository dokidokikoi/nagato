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
	PUUID   string `json:"puuid"`
}

type UpdateMatter struct {
	Name    string `json:"name" binding:"required"`
	Privacy bool   `json:"privacy"`
	Ext     string `json:"ext"  binding:"required"`
	PUUID   string `json:"puuid" binding:"required"`
}

type CreateDir struct {
	Name    string `json:"name" binding:"required"`
	Privacy bool   `json:"privacy"`
	PUUID   string `json:"puuid"`
}
