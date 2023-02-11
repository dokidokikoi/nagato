package controller

import (
	"io"
	"nagato/dataservice/internal/service"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type MatterController struct {
	service service.IService
}

// 将临时文件的信息存储到临时目录并创建出临时文件的文件
func (c MatterController) CreateMatterTemp(ctx *gin.Context) {
	name := ctx.Param("name")

	output, _ := exec.Command("uuidgen").Output()
	uuid := strings.TrimSuffix(string(output), "\n")
	size, e := strconv.ParseInt(ctx.Request.Header.Get("size"), 0, 64)
	if e != nil {
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	if c.service.Matter().CreateTempFile(ctx, name, uuid, size) != nil {
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
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
}

// 临时文件转正，转正后将临时文件删除
func (c MatterController) CommitMatter(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	err := c.service.Matter().CommitMatter(ctx, uuid)
	if err != nil {
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
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	path := c.service.Matter().GetFilePath(ctx, hash)
	f, _ := os.Open(path)
	defer f.Close()
	io.Copy(ctx.Writer, f)
}

func newMatterController(srv service.IService) *MatterController {
	return &MatterController{service: srv}
}
