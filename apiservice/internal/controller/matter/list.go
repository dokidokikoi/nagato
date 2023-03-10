package matter

import (
	"nagato/apiservice/internal/model"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/dokidokikoi/go-common/query"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c MatterController) List(ctx *gin.Context) {
	var pageQuery query.PageQuery
	if ctx.ShouldBindQuery(&pageQuery) != nil {
		zaplog.L().Error("参数校验失败")
		core.WriteResponse(ctx, myErrors.ApiErrValidation, "")
		return
	}

	currentUser := c.GetCurrentUser(ctx)
	res, total, err := c.service.Matter().List(ctx, &model.Matter{UserID: currentUser.ID}, pageQuery.GetListOption())
	if err != nil {
		zaplog.L().Error("获取matter列表失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, "")
		return
	}

	core.WriteListResponse(ctx, nil, total, res)
}
