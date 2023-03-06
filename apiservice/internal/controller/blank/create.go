package blank

import (
	"nagato/apiservice/internal/model"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (b BlankController) Create(ctx *gin.Context) {
	var input CreateBlank
	if ctx.ShouldBindJSON(&input) != nil {
		zaplog.L().Error("参数校验失败")
		core.WriteResponse(ctx, myErrors.ApiErrValidation, "")
		return
	}

	currentUser := b.GetCurrentUser(ctx)

	createBlank := &model.Blank{
		Type:    input.Type,
		Title:   input.Title,
		Content: input.Content,
		Tags:    input.Tags,
		UserID:  currentUser.ID,
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
	}

	if err := b.service.Blank().Create(ctx, createBlank); err != nil {
		zaplog.L().Error("新增blank数据错误", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, "")
		return
	}

	tags := make([]*model.Tag, 0)
	for _, t := range input.Tags {
		tags = append(tags, &model.Tag{TagName: t, UserID: currentUser.ID})
	}
	b.service.Tag().CreateCollection(ctx, tags)

	core.WriteResponse(ctx, myErrors.Success("新增blank成功"), "")
}
