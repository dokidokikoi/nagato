package inittask

import (
	"nagato/dataservice/internal/heartbeat"
	"nagato/dataservice/internal/locate"
	"nagato/dataservice/internal/router"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	locate.CollectMatters()
	router.InitRoutes(r)
}
