package matter

import (
	"errors"
	"nagato/apiservice/internal/model"
	"os/exec"
	"strings"

	commonErrors "nagato/common/errors"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
)

func (c MatterController) CreateDir(ctx *gin.Context) {
	var input CreateDir
	if ctx.ShouldBindJSON(&input) != nil {
		zaplog.L().Error("参数校验失败")
		core.WriteResponse(ctx, myErrors.ApiErrValidation, "")
		return
	}
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		zaplog.L().Sugar().Errorf("生成uuid出错, err: %s", err.Error())
		core.WriteResponse(ctx, myErrors.ApiErrSystemErr, nil)
		return
	}
	currentUser := c.GetCurrentUser(ctx)
	createMatter := &model.Matter{
		UUID:   strings.Trim(string(newUUID), "\n"),
		UserID: currentUser.ID,
		Name:   input.Name,
		PUUID:  input.PUUID,
		Dir:    true,
	}

	c.service.Matter().SetMatterPath(createMatter)
	if err := c.service.Matter().Create(ctx, createMatter); err != nil {
		zaplog.L().Sugar().Errorf("保存matter出错, err: %s", err.Error())
		if errors.Is(err, myErrors.ErrNameDuplicate) {
			core.WriteResponse(ctx, commonErrors.ApiErrFolderRepeatFile, "")
			return
		}
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, "")
		return
	}

	core.WriteResponse(ctx, myErrors.Success("创建成功"), "")
}
