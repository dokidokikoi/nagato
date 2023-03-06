package tag

import (
	"nagato/apiservice/internal/model"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (t TagController) Create(ctx *gin.Context) {
	var input CreateTag
	if ctx.ShouldBindJSON(&input) != nil {
		zaplog.L().Error("参数校验失败")
		core.WriteResponse(ctx, myErrors.ApiErrValidation, "")
		return
	}

	currentUser := t.GetCurrentUser(ctx)
	createTag := &model.Tag{
		TagName: input.TagName,
		UserID:  currentUser.ID,
	}
	err := t.service.Tag().Create(ctx, createTag)
	if err != nil {
		zaplog.L().Error("新增tag数据错误", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, "")
		return
	}

	core.WriteResponse(ctx, myErrors.Success("新增tag成功"), "")
}
