package main

import (
	"nagato/common/middleware"
	"nagato/dataservice/internal/config"
	"nagato/dataservice/internal/heartbeat"
	"nagato/dataservice/internal/locate"
	"nagato/dataservice/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	locate.CollectMatters()
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	router.InitRoutes(r)
	r.Run(config.LISTEN_ADDRESS)
}
