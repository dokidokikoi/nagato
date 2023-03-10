package share

import (
	"nagato/apiservice/internal/model"
	"time"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s ShareController) Delete(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	currentUser := s.GetCurrentUser(ctx)
	share, err := s.service.Share().Get(ctx, &model.Share{UUID: uuid, UserID: currentUser.ID}, nil)
	if err != nil {
		zaplog.L().Error("获取share失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, "")
		return
	}
	share.ExpireInfinity = false
	share.ExpireTime = time.Now()

	if err := s.service.Share().Save(ctx, share); err != nil {
		zaplog.L().Error("更新share失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, "")
		return
	}

	core.WriteResponse(ctx, myErrors.Success("删除share成功"), "")
}
