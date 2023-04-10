package matter

import (
	"nagato/apiservice/internal/model"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c MatterController) Update(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	var input UpdateMatter
	if ctx.ShouldBindJSON(&input) != nil {
		zaplog.L().Error("参数校验失败")
		core.WriteResponse(ctx, myErrors.ApiErrValidation, "")
		return
	}

	matter, err := c.service.Matter().Get(ctx, &model.Matter{UUID: uuid}, nil)
	if err != nil {
		zaplog.L().Error("获取matter失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, "")
		return
	}

	matter.Name = input.Name
	matter.Privacy = input.Privacy
	matter.Ext = input.Ext
	err = c.service.Matter().Update(ctx, matter)
	if err != nil {
		zaplog.L().Error("更新matter失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, "")
		return
	}

	core.WriteResponse(ctx, myErrors.Success("更新matter成功"), "")
}
