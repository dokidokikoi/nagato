package inittask

import (
	"nagato/dataservice/internal/config"
	"nagato/dataservice/internal/heartbeat"
	"nagato/dataservice/internal/locate"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	config.Init("./internal/conf/application.yml")

	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	locate.CollectMatters()

	// r.UseRawPath = true
	// router.InitRoutes(r)
}
