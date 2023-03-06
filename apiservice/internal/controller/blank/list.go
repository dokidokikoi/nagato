package blank

import (
	"nagato/apiservice/internal/model"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/dokidokikoi/go-common/query"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (b BlankController) List(ctx *gin.Context) {
	var pageQuery query.PageQuery
	if ctx.ShouldBindQuery(&pageQuery) != nil {
		zaplog.L().Error("参数校验失败")
		core.WriteResponse(ctx, myErrors.ApiErrValidation, "")
		return
	}

	currentUser := b.GetCurrentUser(ctx)
	listOption := pageQuery.GetListOption()
	listOption.Preload = []string{"Matters"}
	res, total, err := b.service.Blank().List(ctx, &model.Blank{UserID: currentUser.ID}, listOption)
	if err != nil {
		zaplog.L().Error("获取blank列表失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, "")
		return
	}

	core.WriteListResponse(ctx, nil, total, res)
}
