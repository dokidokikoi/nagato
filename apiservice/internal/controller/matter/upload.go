package matter

import (
	"nagato/apiservice/internal/locate"
	"nagato/apiservice/internal/model"
	commonErrors "nagato/common/errors"
	"nagato/common/tools"
	"net/http"
	"os/exec"
	"path"
	"strings"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	"github.com/dokidokikoi/go-common/log/zap"
	meta "github.com/dokidokikoi/go-common/meta/option"
	"github.com/gin-gonic/gin"
)

func (c MatterController) Locate(ctx *gin.Context) {
	hash := ctx.Param("hash")
	info := locate.Locate(hash)
	if len(info) == 0 {
		zap.L().Sugar().Errorf("%s: 文件不存在", hash)
		core.WriteResponse(ctx, commonErrors.ApiErrFileNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, info)
}

func (c MatterController) UploadMatter(ctx *gin.Context) {
	name := ctx.Param("name")
	hash := tools.GetHashFromHeader(ctx.Request.Header)
	if hash == "" {
		zap.L().Sugar().Errorf("name: %s hash: %s 不能为空", name, hash)
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}
	size := tools.GetSizeFromHeader(ctx.Request.Header)

	var err error
	// file, err := ctx.FormFile("file")
	// if err != nil {
	// 	zap.L().Sugar().Errorf("从请求获取文件失败, name: %s, hash: %s, err: %s", name, hash, err.Error())
	// 	ctx.JSON(http.StatusInternalServerError, "")
	// 	return
	// }
	// reader, err := file.Open()
	// if err != nil {
	// 	zap.L().Sugar().Errorf("获取文件失败, name: %s, hash: %s, err: %s", name, hash, err.Error())
	// 	ctx.JSON(http.StatusInternalServerError, "")
	// 	return
	// }
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		zap.L().Sugar().Errorf("生成uuid出错, err: %s", err.Error())
		core.WriteResponse(ctx, myErrors.ApiErrSystemErr, nil)
	}
	ext := strings.TrimLeft(path.Ext(name), ".")
	createMatter := &model.Matter{
		UUID:   strings.Trim(string(newUUID), "\n"),
		UserID: c.GetCurrentUser(ctx).ID,
		Name:   strings.ReplaceAll(name, "."+ext, ""),
		Sha256: hash,
		Size:   uint(size),
		Ext:    ext,
		Path:   "/" + name,
	}

	err = c.service.Matter().Upload(ctx, createMatter, hash, size, ctx.Request.Body)
	if err != nil {
		zap.L().Sugar().Errorf("上传文件失败, name: %s, hash: %s, err: %s", name, hash, err.Error())
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	matter, err := c.service.Matter().Get(ctx, &model.Matter{Path: createMatter.Path}, &meta.GetOption{Include: []string{"path"}})
	if err != nil {
		core.WriteResponse(ctx, myErrors.ApiErrRecordNotFound, nil)
		return
	}
	// 将上传文件元信息加入es
	if err := c.service.Matter().CreateResource(ctx, matter); err != nil {
		zap.L().Sugar().Errorf("上传文件元信息失败, name: %s, hash: %s, err: %s", name, hash, err.Error())
		core.WriteResponse(ctx, myErrors.ApiErrSystemErr, nil)
		return
	}
}
