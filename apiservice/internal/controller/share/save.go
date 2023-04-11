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

	share, err := s.service.Share().
		Get(ctx, &model.Share{UUID: uuid}, &meta.GetOption{Preload: []string{"Matters"}})
	if err != nil {
		zaplog.L().Error("获取share失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, "")
		return
	}
	if share.Code != input.Code {
		zaplog.L().Error("提取码错误")
		core.WriteResponse(ctx, commonErrors.ApiErrShareCode, "")
		return
	}

	if s.service.Share().IsExpired(ctx, share) {
		core.WriteResponse(ctx, commonErrors.ApiErrShareExpired, nil)
		return
	}

	// 保存到的文件夹，没有就保存在根目录
	if input.PUUID != "" {
		_, err := s.service.Matter().
			Get(ctx, &model.Matter{UUID: input.PUUID, Dir: true}, &meta.GetOption{Include: []string{"uuid", "dir"}})
		if err != nil {
			zaplog.L().Error("获取文件夹Matter失败", zap.Error(err))
			core.WriteResponse(ctx, commonErrors.ApiErrFolderNotFound, "")
			return
		}
	}

	m, err := s.service.Share().Receive(share.Matters, input.Matters)
	if err != nil {
		zaplog.L().Error("获取matters树形结构失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, nil)
	}
	currentUser := s.GetCurrentUser(ctx)
	saveMatter := make([]*model.Matter, 0)

	// 递归获取需要存储的文件信息
	var getAllSaveMatters func(matters []*model.Matter) error
	getAllSaveMatters = func(matters []*model.Matter) error {
		if len(matters) == 0 {
			return nil
		}
		for _, m := range matters {
			newUUID, err := exec.Command("uuidgen").Output()
			if err != nil {
				zaplog.L().Sugar().Errorf("生成uuid出错, err: %s", err.Error())
				return err
			}

			v := &model.Matter{
				PUUID:  input.PUUID,
				UserID: currentUser.ID,
				UUID:   strings.Trim(string(newUUID), "\n"),
				Sha256: m.Sha256,
				Name:   m.Name,
				Dir:    m.Dir,
				Size:   m.Size,
			}
			err = s.service.Matter().SetMatterPath(v)
			if err != nil {
				return err
			}
			saveMatter = append(saveMatter, v)

			subMatters, _, err := s.service.Matter().List(ctx, &model.Matter{PUUID: m.UUID}, nil)
			if err != nil {
				zaplog.L().Sugar().Errorf("生成matter列表出错, err: %s", err.Error())
				return err
			}

			if err := getAllSaveMatters(subMatters); err != nil {
				return err
			}
		}
		return nil
	}

	if err := getAllSaveMatters(m); err != nil {
		zaplog.L().Error("获取matters失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, nil)
		return
	}

	errs := s.service.Matter().CreateCollection(ctx, saveMatter)
	if errs != nil {
		zaplog.L().Error("保存matters列表失败", zap.Errors("errs", errs))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, nil)
		return
	}

	core.WriteResponse(ctx, myErrors.Success("保存成功"), "")
}
