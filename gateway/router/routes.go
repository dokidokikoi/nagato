package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InstallRoutes(r *gin.Engine) {
	r.LoadHTMLFiles("./admin-front/index.tmpl")
	r.Static("/pages", "./admin-front/pages")
	r.Static("/public", "./admin-front/public")

	adminR := r.Group("/admin")
	{
		adminR.Use(func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.tmpl", gin.H{})
		})
		adminR.Any("", func(ctx *gin.Context) {})
	}
}
