package inittask

import (
	"nagato/apiservice/internal/heartbeat"
	"nagato/apiservice/internal/router"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	go heartbeat.ListenHeartbeat()

	router.InitRoutes(r)
}
