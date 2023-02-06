package main

import (
	"nagato/apiservice/internal/heartbeat"
	"nagato/apiservice/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	go heartbeat.ListenHeartbeat()
	r := gin.Default()
	router.InitRoutes(r)
	r.Run(":8100")
}
