package controller

import (
	"fmt"
	"io"
	"nagato/common/tools"
	"nagato/dataservice/internal/config"
	"nagato/dataservice/internal/service"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MatterController struct {
	service service.IService
}

// 将临时文件的信息存储到临时目录并创建出临时文件的文件
func (c MatterController) CreateMatterTemp(ctx *gin.Context) {
	hashEncode := url.PathEscape(ctx.Param("name"))

	output, _ := exec.Command("uuidgen").Output()
	uuid := strings.TrimSuffix(string(output), "\n")
	size, err := strconv.ParseInt(ctx.Request.Header.Get("size"), 0, 64)
	if err != nil {
		zaplog.L().Error("创建临时文件信息出错", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	if c.service.Matter().CreateTempFile(ctx, hashEncode, uuid, size) != nil {
		zaplog.L().Error("创建临时文件信息出错", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	io.Copy(ctx.Writer, strings.NewReader(uuid))
}

// 将临时文件存入临时目录
func (c MatterController) SaveMatterTemp(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	err := c.service.Matter().WriteTempFile(ctx, uuid, ctx.Request.Body)
	if err != nil {
		zaplog.L().Error("创建临时文件出错", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
}

// 临时文件转正，转正后将临时文件删除
func (c MatterController) CommitMatter(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	hash := tools.GetHashFromHeader(ctx.Request.Header)
	err := c.service.Matter().CommitMatter(ctx, uuid, hash)
	if err != nil {
		zaplog.L().Error("转正临时文件出错", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
}

// 删除临时文件
func (c MatterController) DelMatterTemp(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	c.service.Matter().DelMatterTemp(ctx, uuid)
	ctx.JSON(http.StatusOK, "")
}

// 获取文件
func (c MatterController) GetMatter(ctx *gin.Context) {
	hash := ctx.Param("hash")

	if hash == "" {
		zaplog.L().Error("hash不能为空")
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	path, err := c.service.Matter().GetFilePath(ctx, hash)
	if err != nil {
		zaplog.L().Sugar().Errorf("获取文件失败, err: %+v", err)
		ctx.JSON(http.StatusNotFound, "")
		return
	}
	f, err := os.Open(path)
	if err != nil {
		zaplog.L().Sugar().Errorf("打开文件失败, err: %+v", err)
		ctx.JSON(http.StatusNotFound, "")
		return
	}
	defer f.Close()
	io.Copy(ctx.Writer, f)
}

func (c MatterController) GetMatterTemp(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	f, err := os.Open(config.Config().FileSystemConfig.TempDir + uuid + ".dat")
	if err != nil {
		zaplog.L().Error("打开文件错误", zap.Error(err))
		ctx.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	io.Copy(ctx.Writer, f)
}

func (c MatterController) HeadMatterTemp(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	f, err := os.Open(config.Config().FileSystemConfig.TempDir + uuid + ".dat")
	if err != nil {
		zaplog.L().Error("打开文件错误", zap.Error(err))
		ctx.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		zaplog.L().Error("获取文件信息错误", zap.Error(err))
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx.Writer.Header().Set("content-length", fmt.Sprintf("%d", info.Size()))
}

func newMatterController(srv service.IService) *MatterController {
	return &MatterController{service: srv}
}
