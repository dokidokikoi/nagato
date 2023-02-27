package app

import (
	inittask "nagato/dataservice/init"
	"nagato/dataservice/internal/rpc"

	"github.com/dokidokikoi/go-common/middleware"

	"github.com/gin-gonic/gin"
)

type App struct {
}

func (a App) Run() {
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	inittask.Init(r)
	// go func() {
	rpc.Init()
	rpc.Run()
	// }()

	// r.Run(config.Config().ServerConfig.Address())
}
