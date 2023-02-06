package main

import (
	"nagato/dataservice/internal/config"
	"nagato/dataservice/internal/heartbeat"
	"nagato/dataservice/internal/locate"
	"nagato/dataservice/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	r := gin.Default()
	router.InitRoutes(r)
	r.Run(config.LISTEN_ADDRESS)
}
