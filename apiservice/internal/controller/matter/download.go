package matter

import (
	"io"
	"nagato/apiservice/internal/model"
	"net/http"
	"strconv"
	"time"

	"github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
)

func (c MatterController) DownloadMatter(ctx *gin.Context) {
	name := ctx.Param("name")
	version, err := strconv.Atoi(ctx.Query("version"))
	if err != nil {
		zap.L().Sugar().Errorf("文件版本号必须为整型, name: %s, verion: %s", name, version)
		ctx.JSON(http.StatusBadRequest, "version字段为数字")
		return
	}

	meta, err := c.service.Matter().GetResourceMate(ctx, name, version)
	if err != nil {
		zap.L().Sugar().Errorf("获取文件元信息失败, name: %s, version: %s, err: %s", name, version, err.Error())
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	if meta.Sha256 == "" {
		zap.L().Sugar().Errorf("文件不存在, name: %s, version: %s", name, version)
		ctx.JSON(http.StatusNotFound, "文件不存在")
		return
	}

	r, err := c.service.Matter().Download(ctx, meta.Sha256)
	if err != nil {
		zap.L().Sugar().Errorf("下载文件失败, name: %s, hash: %s, version: %s, err: %s", name, meta.Sha256, version, err.Error())
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	c.service.Matter().Update(ctx, &model.Matter{ID: meta.ID, Times: meta.Times + 1, VisitTime: time.Now()})

	io.Copy(ctx.Writer, r)
}
