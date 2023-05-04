package blank

import (
	commonEsModel "nagato/common/es/model"
	"nagato/searchservice/internal/service"

	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"go.uber.org/zap"

	"github.com/dokidokikoi/go-common/core"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service service.IService
}

func (c Controller) Search(ctx *gin.Context) {
	var inputB commonEsModel.BlankReq
	var inputR commonEsModel.ResourceReq
	if ctx.ShouldBindJSON(&inputB) != nil || ctx.ShouldBindJSON(&inputR) != nil {
		zaplog.L().Error("参数校验失败")
		core.WriteResponse(ctx, myErrors.ApiErrValidation, "")
		return
	}

	res, total, err := c.service.Blank().Search(ctx, inputB, inputR)
	if err != nil {
		zaplog.L().Error("搜索blank出错", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrSystemErr, "")
		return
	}

	core.WriteListResponse(ctx, nil, total, res)
}

func NewController() Controller {
	return Controller{service: service.NewSrv()}
}
