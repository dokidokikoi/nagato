package blank

import (
	commonEsModel "nagato/common/es/model"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c BlankController) Search(ctx *gin.Context) {

	var input commonEsModel.BlankReq
	if err := ctx.ShouldBindJSON(&input); err != nil {
		zaplog.L().Error("参数校验失败")
		core.WriteResponse(ctx, myErrors.ApiErrValidation, err.Error())
		return
	}

	res, total, err := c.service.Blank().Search(ctx, input, commonEsModel.ResourceReq{})
	if err != nil {
		zaplog.L().Error("搜索blank出错", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrSystemErr, "")
		return
	}

	core.WriteListResponse(ctx, nil, total, res)
}
