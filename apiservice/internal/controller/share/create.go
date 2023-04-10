package share

import (
	"nagato/apiservice/internal/model"
	"nagato/common/tools"
	"os/exec"
	"strings"
	"time"

	commonErrors "nagato/common/errors"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s ShareController) Create(ctx *gin.Context) {
	var input CreateShare
	if err := ctx.ShouldBindJSON(&input); err != nil {
		zaplog.L().Error("参数校验失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrValidation, "")
		return
	}

	if input.Code == "" {
		input.Code = tools.GetRandStr(4)
	}
	if !input.ExpireInfinity || input.ExpireTime.IsZero() {
		input.ExpireTime = time.Now().AddDate(0, 0, 7)
	}

	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		zaplog.L().Sugar().Errorf("生成uuid出错, err: %s", err.Error())
		core.WriteResponse(ctx, myErrors.ApiErrSystemErr, nil)
		return
	}
	currentUser := s.GetCurrentUser(ctx)
	createShare := &model.Share{
		UUID:           strings.Trim(string(newUUID), "\n"),
		UserID:         currentUser.ID,
		Code:           input.Code,
		ExpireInfinity: input.ExpireInfinity,
		ExpireTime:     input.ExpireTime,
		Matters: func() []*model.Matter {
			matters := make([]*model.Matter, len(input.Matters))
			for i := range matters {
				matters[i] = &model.Matter{
					ID: input.Matters[i],
				}
			}
			return matters
		}(),
	}

	if len(input.Matters) <= 0 {
		zaplog.L().Info("share未包含任何资源")
		core.WriteResponse(ctx, commonErrors.ApiErrShareNoMatters, "")
		return
	}

	err = s.service.Share().Create(ctx, createShare)
	if err != nil {
		zaplog.L().Error("新增share数据错误", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, "")
		return
	}

	core.WriteResponse(ctx, myErrors.Success("添加share成功"), "")
}
