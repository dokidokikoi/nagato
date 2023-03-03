package matter

import (
	"fmt"
	"nagato/apiservice/internal/locate"
	"nagato/apiservice/internal/model"
	"nagato/apiservice/stream"
	commonErrors "nagato/common/errors"
	"nagato/common/tools"
	"net/http"
	"net/url"
	"os/exec"
	"path"
	"strings"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	meta "github.com/dokidokikoi/go-common/meta/option"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c MatterController) Locate(ctx *gin.Context) {
	hash, _ := url.PathUnescape(ctx.Param("hash"))
	info := locate.Locate(hash)
	if len(info) == 0 {
		zaplog.L().Sugar().Errorf("%s: 文件不存在", hash)
		core.WriteResponse(ctx, commonErrors.ApiErrFileNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, info)
}

func (c MatterController) UploadMatter(ctx *gin.Context) {
	name := ctx.Param("name")
	hash := tools.GetHashFromHeader(ctx.Request.Header)
	if hash == "" {
		zaplog.L().Sugar().Errorf("name: %s hash: %s 不能为空", name, hash)
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
		zaplog.L().Sugar().Errorf("生成uuid出错, err: %s", err.Error())
		core.WriteResponse(ctx, myErrors.ApiErrSystemErr, nil)
		return
	}
	ext := strings.TrimLeft(path.Ext(name), ".")
	createMatter := &model.Matter{
		UUID:   strings.Trim(string(newUUID), "\n"),
		UserID: c.GetCurrentUser(ctx).ID,
		Name:   strings.ReplaceAll(name, "."+ext, ""),
		Sha256: hash,
		Size:   size,
		Ext:    ext,
		Path:   "/" + name,
	}

	err = c.service.Matter().Upload(ctx, createMatter, ctx.Request.Body)
	if err != nil {
		zaplog.L().Sugar().Errorf("上传文件失败, name: %s, hash: %s, err: %s", name, hash, err.Error())
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	matter, err := c.service.Matter().Get(ctx, &model.Matter{Path: createMatter.Path}, &meta.GetOption{Include: []string{"path"}})
	if err != nil {
		zaplog.L().Sugar().Errorf("获取文件数据库信息失败, name: %s, path: %s, err: %s", matter.Name, matter.Sha256, err.Error())
		core.WriteResponse(ctx, myErrors.ApiErrRecordNotFound, nil)
		return
	}
	// 将上传文件元信息加入es
	if err := c.service.Matter().CreateResource(ctx, matter); err != nil {
		zaplog.L().Sugar().Errorf("上传文件元信息失败, name: %s, hash: %s, err: %s", name, hash, err.Error())
		core.WriteResponse(ctx, myErrors.ApiErrSystemErr, nil)
		return
	}

	core.WriteResponse(ctx, myErrors.Success("上传成功"), nil)
}

func (c MatterController) GenUploadToken(ctx *gin.Context) {
	var input UploadMatter
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, "")
		return
	}

	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		zaplog.L().Sugar().Errorf("生成uuid出错, err: %s", err.Error())
		core.WriteResponse(ctx, myErrors.ApiErrSystemErr, nil)
		return
	}
	ext := strings.TrimLeft(path.Ext(input.Name), ".")
	if input.Path != "" {
		// TODO: 对路径中的文件夹校验是否存在
		// 不存在的情况需要新建文件夹
		if input.Path[len(input.Path)-1:] != "/" {
			input.Path = input.Path + "/"
		}
		if input.Path[0:1] == "/" {
			input.Path = input.Path[1:]
		}
	}
	// TODO: 同一用户下不能存在同一路径的文件
	// TODO: 获取 PUUID 并写入数据库
	createMatter := &model.Matter{
		UUID:    strings.Trim(string(newUUID), "\n"),
		UserID:  c.GetCurrentUser(ctx).ID,
		Name:    strings.ReplaceAll(input.Name, "."+ext, ""),
		Sha256:  input.Sha256,
		Size:    input.Size,
		Ext:     ext,
		Path:    "/" + input.Path + input.Name,
		Privacy: input.Privacy,
	}

	token, err := c.service.Matter().GenUploadToken(ctx, createMatter)
	if err != nil {
		zaplog.L().Sugar().Errorf("生成uploadToken出错, err: %s", err.Error())
		core.WriteResponse(ctx, myErrors.ApiErrSystemErr, nil)
		return
	}
	if token == "" {
		core.WriteResponse(ctx, myErrors.Success("上传成功"), nil)
		return
	}

	core.WriteResponse(ctx, nil, token)
	ctx.AbortWithStatus(http.StatusCreated)
}

func (c MatterController) UploadBigMatter(ctx *gin.Context) {
	token := ctx.Param("token")

	offset := tools.GetOffsetFromHeader(ctx.Request.Header)
	err := c.service.Matter().UploadBigMatter(ctx, token, offset, ctx.Request.Body)
	if err != nil {
		zaplog.L().Error("文件上传失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrSystemErr, "")
		return
	}

	core.WriteResponse(ctx, myErrors.Success("上传成功"), "")
}

func (c MatterController) Head(ctx *gin.Context) {
	token := ctx.Param("token")
	r, err := stream.NewRSResumablePutStreamFromToken(token)
	if err != nil {
		zaplog.L().Error("从token获取Stream失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrSystemErr, "")
		return
	}

	current, err := r.CurrentSize()
	if err != nil {
		zaplog.L().Error("获取已上传文件大小失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrSystemErr, "")
		return
	}
	ctx.Writer.Header().Set("content-length", fmt.Sprintf("%d", current))
	ctx.Writer.Header().Set("per-size", fmt.Sprintf("%d", (1<<21)-((1<<21)%stream.BLOCK_SIZE)))
}
