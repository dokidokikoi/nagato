package blank

import (
	"nagato/apiservice/internal/model"
	"strconv"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (b BlankController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		zaplog.L().Error("参数校验失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrValidation, "")
		return
	}
	var input UpdateBlank
	if ctx.ShouldBindJSON(&input) != nil {
		zaplog.L().Error("参数校验失败")
		core.WriteResponse(ctx, myErrors.ApiErrValidation, "")
		return
	}

	currentUser := b.GetCurrentUser(ctx)
	if _, err := b.service.Blank().Get(ctx, &model.Blank{ID: uint(id), UserID: currentUser.ID}, nil); err != nil {
		zaplog.L().Error("获取blank失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, "")
		return
	}

	updateBlank := &model.Blank{
		ID:      uint(id),
		Type:    input.Type,
		Title:   input.Title,
		Content: input.Content,
		Matters: func() []model.Matter {
			if len(input.Matters) <= 0 {
				return nil
			}

			matters := make([]model.Matter, len(input.Matters))
			for i, _ := range input.Matters {
				matters[i] = model.Matter{
					ID: input.Matters[i],
				}
			}
			return matters
		}(),
		Tags: input.Tags,
	}

	err = b.service.Blank().UpdateBlank(ctx, updateBlank)
	if err != nil {
		zaplog.L().Error("更新blank失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, "")
		return
	}

	tags := make([]*model.Tag, 0)
	for _, t := range input.Tags {
		tags = append(tags, &model.Tag{TagName: t, UserID: currentUser.ID})
	}
	b.service.Tag().CreateCollection(ctx, tags)

	core.WriteResponse(ctx, myErrors.Success("更新blank成功"), nil)
}
