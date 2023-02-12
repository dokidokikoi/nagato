package app

import (
	inittask "nagato/apiservice/init"
	"nagato/apiservice/internal/config"
	"nagato/common/middleware"

	"github.com/gin-gonic/gin"
)

type App struct {
}

func (a App) Run() {
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	inittask.Init(r)
	r.Run(config.Config().ServerConfig.Address())
}
