package blank

import "github.com/gin-gonic/gin"

func (c Controller) Test(ctx *gin.Context) {
	ctx.JSON(200, "hi, nice to meet you")
}
