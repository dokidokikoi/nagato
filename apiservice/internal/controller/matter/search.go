package matter

import (
	commonEsModel "nagato/common/es/model"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c MatterController) Search(ctx *gin.Context) {
	var inputR commonEsModel.ResourceReq
	if ctx.ShouldBindJSON(&inputR) != nil {
		zaplog.L().Error("参数校验失败")
		core.WriteResponse(ctx, myErrors.ApiErrValidation, "")
		return
	}

	res, total, err := c.service.Matter().Search(ctx, inputR)
	if err != nil {
		zaplog.L().Error("搜索resource出错", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrSystemErr, "")
		return
	}

	core.WriteListResponse(ctx, nil, total, res)
}
