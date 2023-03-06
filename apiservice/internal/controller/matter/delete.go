package matter

import (
	"nagato/apiservice/internal/model"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c MatterController) DelMatter(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	// 判断当前用户是否拥有此文件
	currentUser := c.GetCurrentUser(ctx)
	err := c.service.Matter().Del(ctx, &model.Matter{UUID: uuid, UserID: currentUser.ID}, nil)
	if err != nil {
		zaplog.L().Error("删除文件失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, nil)
		return
	}

	// TODO: 更新es索引库
	// err = c.service.Matter().CreateLastestResource(ctx, name, "", 0)
	// if err != nil {
	// 	zap.L().Sugar().Errorf("删除文件失败, name: %s, err: %s", name, err.Error())
	// 	ctx.JSON(http.StatusInternalServerError, "")
	// 	return
	// }

	core.WriteResponse(ctx, myErrors.Success("删除文件成功"), nil)
}
