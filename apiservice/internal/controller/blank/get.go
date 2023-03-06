package blank

import (
	"nagato/apiservice/internal/model"
	"strconv"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	meta "github.com/dokidokikoi/go-common/meta/option"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (b BlankController) Get(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		zaplog.L().Error("参数校验失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrValidation, "")
		return
	}

	currentUser := b.GetCurrentUser(ctx)
	blank, err := b.service.Blank().Get(ctx, &model.Blank{ID: uint(id), UserID: currentUser.ID}, &meta.GetOption{Preload: []string{"Matters"}})
	if err != nil {
		zaplog.L().Error("获取blank失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, "")
		return
	}

	core.WriteResponse(ctx, nil, blank)
}
