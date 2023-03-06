package matter

import (
	"nagato/apiservice/internal/model"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c MatterController) Get(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	currentUser := c.GetCurrentUser(ctx)
	matter, err := c.service.Matter().Get(ctx, &model.Matter{UUID: uuid, UserID: currentUser.ID}, nil)
	if err != nil {
		zaplog.L().Error("获取matter失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, "")
		return
	}

	children, err := c.service.Matter().ListMatter(ctx, &model.Matter{PUUID: matter.UUID}, nil)
	if err != nil {
		zaplog.L().Error("获取matter失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, "")
		return
	}

	matter.Children = children
	core.WriteResponse(ctx, nil, matter)
}
