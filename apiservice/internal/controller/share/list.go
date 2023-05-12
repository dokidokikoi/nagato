package share

import (
	"nagato/apiservice/internal/model"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/dokidokikoi/go-common/query"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s ShareController) List(ctx *gin.Context) {
	var pageQuery query.PageQuery
	if ctx.ShouldBindQuery(&pageQuery) != nil {
		zaplog.L().Error("参数校验失败")
		core.WriteResponse(ctx, myErrors.ApiErrValidation, "")
		return
	}

	option := pageQuery.GetListOption()
	option.GetOption.Preload = []string{"Matters"}
	currentUser := s.GetCurrentUser(ctx)
	res, total, err := s.service.Share().List(ctx, &model.Share{UserID: currentUser.ID}, option)
	if err != nil {
		zaplog.L().Error("获取matter列表失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, "")
		return
	}

	core.WriteListResponse(ctx, nil, total, res)
}
