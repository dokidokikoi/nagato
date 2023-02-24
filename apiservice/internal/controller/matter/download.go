package matter

import (
	"io"
	"nagato/apiservice/internal/model"
	commonErrors "nagato/common/errors"
	"time"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	"github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
)

func (c MatterController) DownloadMatter(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	matter, err := c.service.Matter().Get(ctx, &model.Matter{UUID: uuid}, nil)
	if err != nil {
		zap.L().Sugar().Errorf("文件不存在: uuid: %s, err: %s", uuid, err.Error())
		core.WriteResponse(ctx, commonErrors.ApiErrFileNotFound, nil)
		return
	}
	r, err := c.service.Matter().Download(ctx, matter.Sha256, matter.Size)
	if err != nil {
		zap.L().Sugar().Errorf("下载文件失败: path: %s, hash: %s, err: %s", matter.Path, matter.Sha256, err.Error())
		core.WriteResponse(ctx, myErrors.ApiErrSystemErr, nil)
		return
	}
	// GET对象时会对缺失的分片进行即时修复,
	// 修复的过程也使用数据服务的 temp 接口,
	// RSGetStream 的 Close 方法用于在流关闭时将临时对象转正
	defer r.Close()

	c.service.Matter().Update(ctx, &model.Matter{ID: matter.ID, Times: matter.Times + 1, VisitTime: time.Now()})

	io.Copy(ctx.Writer, r)
}
