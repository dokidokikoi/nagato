package router

import (
	"nagato/apiservice/internal/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	control := controller.NewController()
	fileR := r.Group("/file")
	{
		fileController := control.Matter()
		fileR.GET("/locate/:hash", fileController.Locate)
		fileR.PUT("/:name", fileController.UploadMatter)
		fileR.GET("/:name", fileController.DownloadMatter)
		fileR.DELETE("/:name", fileController.DelMatter)
		fileR.GET("/versions/:name", fileController.VersionList)
	}
}
