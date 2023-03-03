package router

import (
	"nagato/apiservice/internal/controller/matter"
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
		// fileR.GET("/user/:id", fileController.UserMatterList)
	}

	// userR := r.Group("/user")
	// {

	// }
}
