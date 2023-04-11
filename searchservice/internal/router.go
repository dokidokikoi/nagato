package search

import (
	"nagato/searchservice/internal/controller/blank"

	"github.com/gin-gonic/gin"
)

func InitRoute(r gin.IRouter) {
	blankR := r.Group("/blank")
	{
		balnkCtrl := blank.NewController()
		blankR.GET("", balnkCtrl.Test)
	}

	resourceR := r.Group("/resource")
	{
		resourceR.GET("")
	}
}
