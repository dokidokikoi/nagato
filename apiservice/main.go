package main

import (
	"nagato/apiservice/internal/heartbeat"
	"nagato/apiservice/internal/router"
	"nagato/common/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	go heartbeat.ListenHeartbeat()
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	router.InitRoutes(r)
	r.Run(":8100")
}
