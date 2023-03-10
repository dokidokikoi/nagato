package share

import (
	"nagato/apiservice/internal/model"
	"os/exec"
	"strings"

	commonErrors "nagato/common/errors"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	meta "github.com/dokidokikoi/go-common/meta/option"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s ShareController) Save(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	var input SaveShare
	if ctx.ShouldBindJSON(&input) != nil {
		zaplog.L().Error("参数校验失败")
		core.WriteResponse(ctx, myErrors.ApiErrValidation, "")
		return
	}

	share, err := s.service.Share().Get(ctx, &model.Share{UUID: uuid, Code: input.Code}, &meta.GetOption{Preload: []string{"Matters"}})
	if err != nil {
		zaplog.L().Error("获取share失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, "")
		return
	}

	if input.PUUID != "" {
		_, err := s.service.Matter().Get(ctx, &model.Matter{UUID: input.PUUID, Dir: true}, &meta.GetOption{Include: []string{"uuid", "dir"}})
		if err != nil {
			zaplog.L().Error("获取文件夹Matter失败", zap.Error(err))
			core.WriteResponse(ctx, commonErrors.ApiErrFolderNotFound, "")
			return
		}
	}

	m, errs := s.service.Matter().GetAllMatter(share.Matters)
	if errs != nil {
		zaplog.L().Error("获取matters树形结构失败", zap.Errors("errs", errs))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, nil)
	}
	currentUser := s.GetCurrentUser(ctx)
	saveMatter := make([]*model.Matter, 0)
	for _, i := range input.Matters {
		m, ok := m[i]
		if !ok {
			continue
		}
		newUUID, err := exec.Command("uuidgen").Output()
		if err != nil {
			zaplog.L().Sugar().Errorf("生成uuid出错, err: %s", err.Error())
			core.WriteResponse(ctx, myErrors.ApiErrSystemErr, nil)
			continue
		}

		saveMatter = append(saveMatter, &model.Matter{
			PUUID:  input.PUUID,
			UserID: currentUser.ID,
			UUID:   strings.Trim(string(newUUID), "\n"),
			Sha256: m.Sha256,
			Name:   m.Name,
			Dir:    m.Dir,
			Size:   m.Size,
		})
	}

	errs = s.service.Matter().CreateCollection(ctx, saveMatter)
	if errs != nil {
		zaplog.L().Error("获取matters树形结构失败", zap.Errors("errs", errs))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, nil)
		return
	}

	core.WriteResponse(ctx, myErrors.Success("保存成功"), "")
}
