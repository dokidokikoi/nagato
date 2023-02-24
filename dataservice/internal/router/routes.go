package router

import (
	"nagato/dataservice/internal/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	dataR := r.Group("/data")

	control := controller.NewController()
	fileR := dataR.Group("/file")
	{
		fileController := control.Matter()
		fileR.POST("/temp/:name", fileController.CreateMatterTemp)
		fileR.PATCH("/temp/:uuid", fileController.SaveMatterTemp)
		fileR.PUT("/temp/:uuid", fileController.CommitMatter)
		fileR.DELETE("/temp/:uuid", fileController.DelMatterTemp)

		fileR.GET("/:hash", fileController.GetMatter)
	}
}
