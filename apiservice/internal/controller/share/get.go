package share

import (
	"nagato/apiservice/internal/model"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	meta "github.com/dokidokikoi/go-common/meta/option"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func (s ShareController) Get(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	share, err := s.service.Share().Get(ctx, &model.Share{UUID: uuid, Code: ""}, &meta.GetOption{Preload: []string{"Matters", "User"}})
	if err != nil {
		zaplog.L().Error("获取share失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, "")
		return
	}

	core.WriteResponse(ctx, nil, share)
}
