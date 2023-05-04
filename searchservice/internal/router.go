package search

import (
	"nagato/searchservice/internal/controller/blank"
	"nagato/searchservice/internal/controller/resource"

	"github.com/gin-gonic/gin"
)

func InitRoute(r gin.IRouter) {
	// r.Use(middleware.Auth())
	blankR := r.Group("/blank")
	{
		balnkCtrl := blank.NewController()
		blankR.POST("", balnkCtrl.Search)
	}

	resourceR := r.Group("/resource")
	{
		resourceCtrl := resource.NewController()
		resourceR.POST("", resourceCtrl.Search)
	}
}
