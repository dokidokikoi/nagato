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
		fileR.GET("/locate/:hash", fileController.Locate)
		fileR.PUT("/:name", fileController.UploadMatter)
		fileR.GET("/:uuid", fileController.DownloadMatter)
		fileR.DELETE("/:uuid", fileController.DelMatter)
		fileR.HEAD("/:uuid", fileController.HeadMatter)

		fileR.POST("", fileController.GenUploadToken)
		fileR.HEAD("/temp/:token", fileController.Head)
		fileR.PUT("/temp/:token", fileController.UploadBigMatter)
	}

	matterR := apiR.Group("/matter")
	{
		matterController := matter.NewMatterController()
		matterR.POST("/dir", matterController.CreateDir)
		matterR.GET("/:uuid", matterController.Get)
		matterR.GET("", matterController.List)
		matterR.PATCH("/:uuid", matterController.Update)
	}

	blankR := apiR.Group("/blank")
	{
		blankController := blank.NewBlankController()
		blankR.POST("", blankController.Create)
		blankR.GET("", blankController.List)
		blankR.PATCH("/:id", blankController.Update)
		blankR.DELETE("/:id", blankController.Delete)
		blankR.GET("/:id", blankController.Get)
		blankR.POST("/search", blankController.Search)
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

	resourceR := apiR.Group("/resource")
	{
		resourceCtrl := matter.NewMatterController()
		resourceR.POST("/search", resourceCtrl.Search)
	}

	// userR := r.Group("/user")
	// {

	// }
}
