package inittask

import (
	"nagato/apiservice/internal/db/data"
	"nagato/apiservice/internal/heartbeat"
	"nagato/apiservice/internal/router"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	go heartbeat.ListenHeartbeat()
	data.SetStoreDBFactory()

	router.InitRoutes(r)
}
