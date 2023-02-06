package controller

import (
	"encoding/json"
	"io"
	"nagato/apiservice/internal/es"
	"nagato/apiservice/internal/locate"
	"nagato/apiservice/internal/service"
	commonEs "nagato/common/es"
	"nagato/common/tools"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MatterController struct {
	service service.IService
}

func (c MatterController) Locate(ctx *gin.Context) {
	name := ctx.Param("name")
	info := locate.Locate(name)
	if len(info) == 0 {
		ctx.JSON(http.StatusNotFound, "")
		return
	}
	b, _ := json.Marshal(info)
	ctx.JSON(http.StatusOK, b)
}

func (c MatterController) UploadMatter(ctx *gin.Context) {
	name := ctx.Param("name")
	hash := tools.GetHashFromHeader(ctx.Request.Header)
	if hash == "" {
		ctx.JSON(http.StatusBadRequest, "")
		return
	}
	size := tools.GetSizeFromHeader(ctx.Request.Header)

	err := c.service.Matter().Upload(ctx, hash, size, ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	// 将上传文件元信息加入es
	if es.CreateLastestResource(name, hash, size) != nil {
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
}

func (c MatterController) DownloadMatter(ctx *gin.Context) {
	name := ctx.Param("name")
	version, err := strconv.Atoi(ctx.Query("version"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "version字段为数字")
		return
	}

	meta, err := es.GetResourceMate(name, version)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	if meta.Hash == "" {
		ctx.JSON(http.StatusNotFound, "文件不存在")
		return
	}

	r, err := c.service.Matter().Download(ctx, meta.Hash)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	io.Copy(ctx.Writer, r)
}

func (c MatterController) DelMatter(ctx *gin.Context) {
	name := ctx.Param("name")
	err := es.CreateLastestResource(name, "", 0)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
}

func (c MatterController) VersionList(ctx *gin.Context) {
	name := ctx.Param("name")
	from := 0
	size := 100
	var res []commonEs.Metadata
	for {
		metas, err := es.SearchResourceAllVersion(name, from, size)
		if err != nil {
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
