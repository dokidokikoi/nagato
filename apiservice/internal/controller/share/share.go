package share

import (
	"nagato/apiservice/internal/controller"
	"nagato/apiservice/internal/service"
	"time"
)

type ShareController struct {
	controller.Controller
	service service.IService
}

func NewShareController() *ShareController {
	return &ShareController{service: service.NewService()}
}

type CreateShare struct {
	Code           string    `json:"code"`            // 提取码
	ExpireInfinity bool      `json:"expire_infinity"` // 是否永不过期
	ExpireTime     time.Time `json:"expireTime"`
	Matters        []uint    `json:"matters" binding:"min=0"`
}

type SaveShare struct {
	Code    string `json:"code" binding:"required"`
	PUUID   string `json:"puuid"`
	Matters []uint `json:"matters" binding:"min=0"`
}
