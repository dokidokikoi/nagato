package controller

import (
	"io"
	"nagato/apiservice/internal/es"
	"nagato/apiservice/internal/locate"
	"nagato/apiservice/internal/model"
	"nagato/apiservice/internal/service"
	"nagato/common/tools"
	"net/http"
	"strconv"

	"github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
)

type MatterController struct {
	service service.IService
}

func (c MatterController) Locate(ctx *gin.Context) {
	hash := ctx.Param("hash")
	info := locate.Locate(hash)
	if len(info) == 0 {
		zap.L().Sugar().Errorf("%s: 文件不存在", hash)
		ctx.JSON(http.StatusNotFound, "文件不存在")
		return
	}
	ctx.JSON(http.StatusOK, info)
}

func (c MatterController) UploadMatter(ctx *gin.Context) {
	name := ctx.Param("name")
	hash := tools.GetHashFromHeader(ctx.Request.Header)
	if hash == "" {
		zap.L().Sugar().Errorf("name: %s hash: %s 不能为空", name, hash)
		ctx.JSON(http.StatusBadRequest, "")
		return
	}
	size := tools.GetSizeFromHeader(ctx.Request.Header)
	file, err := ctx.FormFile("file")
	if err != nil {
		zap.L().Sugar().Errorf("从请求获取文件失败, name: %s, hash: %s, err: %s", name, hash, err.Error())
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	reader, err := file.Open()
	if err != nil {
		zap.L().Sugar().Errorf("获取文件失败, name: %s, hash: %s, err: %s", name, hash, err.Error())
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	createMatter := &model.Matter{
		Name:   file.Filename,
		Sha256: hash,
		Size:   uint(file.Size),
	}
	if err = c.service.Matter().Create(ctx, createMatter); err != nil {
		zap.L().Sugar().Errorf("保存文件信息失败, name: %s, hash: %s, err: %s", name, hash, err.Error())
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	err = c.service.Matter().Upload(ctx, hash, file.Size, reader)
	if err != nil {
		zap.L().Sugar().Errorf("上传文件失败, name: %s, hash: %s, err: %s", name, hash, err.Error())
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	// 将上传文件元信息加入es
	if err := c.service.Matter().CreateLastestResource(ctx, name, hash, size); err != nil {
		zap.L().Sugar().Errorf("上传文件元信息失败, name: %s, hash: %s, err: %s", name, hash, err.Error())
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
}

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

	if meta.Hash == "" {
		zap.L().Sugar().Errorf("文件不存在, name: %s, version: %s", name, version)
		ctx.JSON(http.StatusNotFound, "文件不存在")
		return
	}

	r, err := c.service.Matter().Download(ctx, meta.Hash)
	if err != nil {
		zap.L().Sugar().Errorf("下载文件失败, name: %s, hash: %s, version: %s, err: %s", name, meta.Hash, version, err.Error())
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	io.Copy(ctx.Writer, r)
}

func (c MatterController) DelMatter(ctx *gin.Context) {
	name := ctx.Param("name")
	err := c.service.Matter().CreateLastestResource(ctx, name, "", 0)
	if err != nil {
		zap.L().Sugar().Errorf("删除文件失败, name: %s, err: %s", name, err.Error())
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
}

func (c MatterController) VersionList(ctx *gin.Context) {
	name := ctx.Param("name")
	from := 0
	size := 100
	var res []*es.Resource
	for {
		metas, err := c.service.Matter().SearchResourceAllVersion(name, from, size)
		if err != nil {
			zap.L().Sugar().Errorf("获取文件元信息失败, name: %s, err: %s", name, err.Error())
			ctx.JSON(http.StatusInternalServerError, "")
			return
		}

		res = append(res, metas...)
		if len(metas) < size {
			break
		}
		from += size
	}
	ctx.JSON(http.StatusOK, res)
}

func newMatterController(srv service.IService) *MatterController {
	return &MatterController{service: srv}
}
