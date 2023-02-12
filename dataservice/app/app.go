package app

import (
	"nagato/common/middleware"
	inittask "nagato/dataservice/init"
	"nagato/dataservice/internal/config"

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
