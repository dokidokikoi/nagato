package tag

import (
	"nagato/apiservice/internal/model"
	"strconv"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (t TagController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		zaplog.L().Error("参数校验失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrValidation, "")
		return
	}

	currentUser := t.GetCurrentUser(ctx)
	err = t.service.Tag().Del(ctx, &model.Tag{ID: uint(id), UserID: currentUser.ID})
	if err != nil {
		zaplog.L().Error("删除tag失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, "")
		return
	}

	core.WriteResponse(ctx, myErrors.Success("删除blank成功"), "")
}
