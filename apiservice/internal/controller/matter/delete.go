package matter

import (
	"nagato/apiservice/internal/model"
	"strconv"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	meta "github.com/dokidokikoi/go-common/meta/option"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c MatterController) DelMatter(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 32)
	if err != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}
	currentUser := c.GetCurrentUser(ctx)
	matter, err := c.service.Matter().Get(ctx, &model.Matter{ID: uint(id), UserID: currentUser.ID}, nil)
	if err != nil || matter == nil {
		zaplog.L().Error("删除文件失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrAccessDenied, nil)
		return
	}
	err = c.service.Matter().Del(ctx, &model.Matter{UserID: currentUser.ID}, &meta.DeleteOption{Select: []string{"user_id"}})
	if err != nil {
		zaplog.L().Error("删除文件失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, nil)
		return
	}
	// err = c.service.Matter().CreateLastestResource(ctx, name, "", 0)
	// if err != nil {
	// 	zap.L().Sugar().Errorf("删除文件失败, name: %s, err: %s", name, err.Error())
	// 	ctx.JSON(http.StatusInternalServerError, "")
	// 	return
	// }

	core.WriteResponse(ctx, myErrors.Success("删除文件成功"), nil)
}
