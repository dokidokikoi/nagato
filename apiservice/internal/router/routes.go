package router

import (
	"nagato/apiservice/internal/controller/blank"
	"nagato/apiservice/internal/controller/matter"
	"nagato/apiservice/internal/controller/share"
	"nagato/apiservice/internal/controller/tag"
	"nagato/apiservice/internal/controller/user"
	"nagato/apiservice/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	apiR := r.Group("/api")

	userController := user.NewUserController()
	apiR.POST("/login", userController.Login)
	apiR.POST("/register", userController.Register)

	apiR.Use(middleware.Auth())
	fileR := apiR.Group("/file")
	{
		fileController := matter.NewMatterController()
		// 判断文件是否存在
		fileR.GET("/locate/:hash", fileController.Locate)
		// 一次性完成上传文件
		fileR.PUT("/:name", fileController.UploadMatter)
		// 下载文件
		fileR.GET("/:uuid", fileController.DownloadMatter)
		// 移除文件
		fileR.DELETE("/:uuid", fileController.DelMatter)
		// 获取已上传文件大小和文件名
		fileR.HEAD("/:uuid", fileController.HeadMatter)

		// 获取分片上传的token
		fileR.POST("", fileController.GenUploadToken)
		// 获取已上传的大小和分片的大小
		fileR.HEAD("/temp/:token", fileController.Head)
		// 分片上传
		fileR.PUT("/temp/:token", fileController.UploadBigMatter)
	}

	matterR := apiR.Group("/matter")
	{
		matterController := matter.NewMatterController()
		// 创建文件夹
		matterR.POST("/dir", matterController.CreateDir)
		// 获取目录树
		matterR.GET("/:uuid", matterController.Get)
		// 获取全部文件信息
		matterR.GET("", matterController.List)
		// 修改文件元信息
		matterR.PATCH("/:uuid", matterController.Update)
	}

	blankR := apiR.Group("/blank")
	{
		blankController := blank.NewBlankController()
		// 创建blank
		blankR.POST("", blankController.Create)
		blankR.GET("", blankController.List)
		blankR.PATCH("/:id", blankController.Update)
		blankR.DELETE("/:id", blankController.Delete)
		blankR.GET("/:id", blankController.Get)
	}

	tagR := apiR.Group("/tag")
	{
		tagController := tag.NewTagController()
		tagR.GET("", tagController.List)
		tagR.POST("", tagController.Create)
		tagR.DELETE("/:id", tagController.Delete)
	}

	shareR := apiR.Group("/share")
	{
		shareController := share.NewShareController()
		shareR.GET("/:uuid", shareController.Get)
		shareR.POST("", shareController.Create)
		shareR.DELETE("/:uuid", shareController.Delete)
		shareR.PUT("/:uuid", shareController.Save)

	}

	// userR := r.Group("/user")
	// {

	// }
}
